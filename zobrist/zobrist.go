// Package zobrist implements Zobrist hashing for chess position fingerprinting.
//
// Zobrist hashing assigns a random 64-bit key to each unique aspect of a chess
// position (piece-color-square combinations, castling rights, en passant squares,
// and side to move). A position's hash is the XOR of all applicable keys, allowing
// incremental updates when making/unmaking moves.
//
// The random keys are generated using crypto/rand, ensuring high-quality randomness
// with negligible collision probability for practical use in transposition tables.
package zobrist

import (
	"crypto/rand"
	"encoding/binary"

	"saiko.cz/sachista/bitboard"
)

// Keys holds the random values for all board state components used in Zobrist hashing.
// The hash does not include move clocks (half-move clock and full move number).
type Keys struct {
	Pieces    [bitboard.NumberOfColors][bitboard.NumberOfPieces + 1][bitboard.NumberOfSquares]uint64
	Castling  [bitboard.NumberOfColors][bitboard.NumberOfCastlingOptions]uint64
	EnPassant [bitboard.NumberOfSquares]uint64
	Side      uint64
}

// NewKeys generates a new set of cryptographically random Zobrist keys.
func NewKeys() *Keys {
	z := &Keys{}

	for square := range bitboard.NumberOfSquares {
		for side := range bitboard.NumberOfColors {
			for piece := range bitboard.NumberOfPieces + 1 {
				z.Pieces[side][piece][square] = randUInt64()
			}
		}
		z.EnPassant[square] = randUInt64()
	}

	for i := range 4 {
		z.Castling[0][i] = randUInt64()
		z.Castling[1][i] = randUInt64()
	}

	z.Side = randUInt64()
	return z
}

// randUInt64 returns a cryptographically random 64-bit unsigned integer.
func randUInt64() uint64 {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(b[:])
}
