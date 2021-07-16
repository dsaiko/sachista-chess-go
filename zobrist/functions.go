package zobrist

import (
	"math/rand"
	"saiko.cz/sachista/constants"
)

func NewZobrist() *Zobrist {
	z := &Zobrist{}

	// Generate random values for all unique states
	// We do not need to seed the generator, numbers may be the same each time

	for square := 0; square < constants.NumberOfSquares; square++ {
		for side := 0; side < 2; side++ {
			for piece := 0; piece < 7; piece++ {
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
