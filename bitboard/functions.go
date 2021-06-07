package bitboard

import "math/bits"

// PopCount returns the number of bits set in the board
func (b bitboard) PopCount() int {
	return bits.OnesCount64(uint64(b))
}

// BitScan returns the index of first 1 bit or 64 if no bits are set
func (b bitboard) BitScan() int {
	return bits.TrailingZeros64(uint64(b))
}

// BitPop returns index of first set bit and resets this bit in the board
func (b *bitboard) BitPop() int {
	i := b.BitScan()
	*b &= *b - 1

	return i
}
