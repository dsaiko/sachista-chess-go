package zobrist

/// RANDOM RKISS is our pseudo random number generator (PRNG) used to compute hash keys.
/// George Marsaglia invented the RNG-Kiss-family in the early 90's. This is a
/// specific version that Heinz van Saanen derived from some public domain code
/// by Bob Jenkins. Following the feature list, as tested by Heinz.
///
/// - Quite platform independent
/// - Passes ALL dieharder tests! Here *nix sys-rand() e.g. fails miserably:-)
/// - ~12 times faster than my *nix sys-rand()
/// - ~4 times faster than SSE2-version of Mersenne twister
/// - Average cycle length: ~2^126
/// - 64 bit seed
/// - Return doubles with a full 53 bit mantissa
/// - Thread safe
///
/// PRNG Inspired by Stockfish GPL source code

import (
	"math/rand"

	"saiko.cz/sachista/constants"
)

type ZobristKeys struct {
	RndPieces    [constants.NumberOfColors][constants.NumberOfPieces + 1][constants.NumberOfSquares]uint64
	RndCastling  [constants.NumberOfColors][constants.NumberOfCastlingOptions]uint64
	RndEnPassant [constants.NumberOfSquares]uint64
	RndSide      uint64
}

func NewZobristKeys() *ZobristKeys {
	z := &ZobristKeys{}

	// Generate random values for all unique states
	// We do not need to seed the generator, numbers may be the same each time

	for square := 0; square < constants.NumberOfSquares; square++ {
		for side := 0; side < constants.NumberOfColors; side++ {
			for piece := 0; piece < constants.NumberOfPieces+1; piece++ {
				z.RndPieces[side][piece][square] = rand.Uint64()
			}
		}
		z.RndEnPassant[square] = rand.Uint64()
	}

	for i := 0; i < 4; i++ {
		z.RndCastling[0][i] = rand.Uint64()
		z.RndCastling[1][i] = rand.Uint64()
	}

	z.RndSide = rand.Uint64()
	return z
}
