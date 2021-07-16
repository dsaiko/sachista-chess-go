package zobrist

import "saiko.cz/sachista/constants"

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

type Zobrist struct {
	RndPieces    [constants.NumberOfColors][constants.NumberOfPieces + 1][constants.NumberOfSquares]uint64
	RndCastling  [constants.NumberOfColors][constants.NumberOfCastlingOptions]uint64
	RndEnPassant [constants.NumberOfSquares]uint64
	RndSide      uint64
}
