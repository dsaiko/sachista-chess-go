package board

import (
	"math/bits"
	"saiko.cz/sachista/index"
)

// PopCount returns the number of bits set in the board
func (b BitBoard) PopCount() int {
	return bits.OnesCount64(uint64(b))
}

// BitScan returns the index of first 1 bit or 64 if no bits are set
func (b BitBoard) BitScan() int {
	return bits.TrailingZeros64(uint64(b))
}

// BitPop returns index of first set bit and resets this bit in the board
func (b *BitBoard) BitPop() int {
	i := b.BitScan()
	*b &= *b - 1

	return i
}

func FromNotation(notations ...string) BitBoard {
	b := Empty

	for _, n := range notations {
		b |= FromIndex1(index.FromNotation(n))
	}
	return b
}

func FromIndex(indices ...index.Index) BitBoard {
	b := Empty

	for _, i := range indices {
		b |= FromIndex1(i)
	}
	return b
}

func FromIndex1(i index.Index) BitBoard {
	return 1 << i
}

func (b BitBoard) OneNorth() BitBoard {
	return b << 8
}

func (b BitBoard) OneSouth() BitBoard {
	return b >> 8
}

func (b BitBoard) OneEast() BitBoard {
	return (b << 1) & ^FileA
}

func (b BitBoard) OneNorthEast() BitBoard {
	return (b << 9) & ^FileA
}

func (b BitBoard) OneSouthEast() BitBoard {
	return (b << 7) & ^FileA
}

func (b BitBoard) OneWest() BitBoard {
	return (b >> 1) & ^FileH
}

func (b BitBoard) OneSouthWest() BitBoard {
	return (b >> 9) & ^FileH
}

func (b BitBoard) OneNorthWest() BitBoard {
	return (b >> 7) & ^FileH
}

func (b BitBoard) String() string {
	return "?"
}
