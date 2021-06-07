package bitboard

import "math/bits"

func (b bitboard) PopCount() int {
	return bits.OnesCount64(uint64(b))
}
