package zobrist

// RANDOM RKISS is our pseudo random number generator (PRNG) used to compute hash keys.
// George Marsaglia invented the RNG-Kiss-family in the early 90's. This is a
// specific version that Heinz van Saanen derived from some public domain code
// by Bob Jenkins. Following the feature list, as tested by Heinz.
//
// - Quite platform independent
// - Passes ALL dieharder tests! Here *nix sys-rand() e.g. fails miserably:-)
// - ~12 times faster than my *nix sys-rand()
// - ~4 times faster than SSE2-version of Mersenne twister
// - Average cycle length: ~2^126
// - 64 bit seed
// - Return doubles with a full 53 bit mantissa
// - Thread safe
//
// PRNG Inspired by Stockfish GPL source code

import (
	"crypto/rand"
	"encoding/binary"
	"saiko.cz/sachista/bitboard"
)

// Keys for Zobrist checksum of the board
// Hash does not include move clocks
type Keys struct {
	Pieces    [bitboard.NumberOfColors][bitboard.NumberOfPieces + 1][bitboard.NumberOfSquares]uint64
	Castling  [bitboard.NumberOfColors][bitboard.NumberOfCastlingOptions]uint64
	EnPassant [bitboard.NumberOfSquares]uint64
	Side      uint64
}

// NewKeys initializes new random number keys
func NewKeys() *Keys {
	z := &Keys{}

	// Generate random values for all unique states
	// We do not need to seed the generator, numbers may be the same each time

	for square := 0; square < bitboard.NumberOfSquares; square++ {
		for side := 0; side < bitboard.NumberOfColors; side++ {
			for piece := 0; piece < bitboard.NumberOfPieces+1; piece++ {
				z.Pieces[side][piece][square] = randUInt64()
			}
		}
		z.EnPassant[square] = randUInt64()
	}

	for i := 0; i < 4; i++ {
		z.Castling[0][i] = randUInt64()
		z.Castling[1][i] = randUInt64()
	}

	z.Side = randUInt64()
	return z
}

// randUInt64 generate random number
// https://stackoverflow.com/questions/44482738/random-64-bit-integer-from-crypto-rand
func randUInt64() uint64 {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(b[:])
}
