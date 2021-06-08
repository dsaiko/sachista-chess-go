package board

import (
	"bytes"
	"math/bits"
	"saiko.cz/sachista/index"
	"strconv"
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

// FromNotation returns BitBoard filled with pieces defined by notation parameters
func FromNotation(notations ...string) BitBoard {
	b := Empty

	for _, n := range notations {
		b |= FromIndex1(index.FromNotation(n))
	}
	return b
}

// FromIndex returns BitBoard filled with pieces defined by index parameters
func FromIndex(indices ...index.Index) BitBoard {
	b := Empty

	for _, i := range indices {
		b |= FromIndex1(i)
	}
	return b
}

// FromIndex1 returns BitBoard of one piece index
func FromIndex1(i index.Index) BitBoard {
	return 1 << i
}

// OneNorth shifts all existing BitBoard pieces by one
func (b BitBoard) OneNorth() BitBoard {
	return b << 8
}

// OneSouth  shifts all existing BitBoard pieces by one
func (b BitBoard) OneSouth() BitBoard {
	return b >> 8
}

// OneEast  shifts all existing BitBoard pieces by one
func (b BitBoard) OneEast() BitBoard {
	return (b << 1) & ^FileA
}

// OneNorthEast  shifts all existing BitBoard pieces by one
func (b BitBoard) OneNorthEast() BitBoard {
	return (b << 9) & ^FileA
}

// OneSouthEast  shifts all existing BitBoard pieces by one
func (b BitBoard) OneSouthEast() BitBoard {
	return (b >> 7) & ^FileA
}

// OneWest  shifts all existing BitBoard pieces by one
func (b BitBoard) OneWest() BitBoard {
	return (b >> 1) & ^FileH
}

// OneSouthWest  shifts all existing BitBoard pieces by one
func (b BitBoard) OneSouthWest() BitBoard {
	return (b >> 9) & ^FileH
}

// OneNorthWest  shifts all existing BitBoard pieces by one
func (b BitBoard) OneNorthWest() BitBoard {
	return (b << 7) & ^FileH
}

// Shift shifts all existing BitBoard pieces by multiple steps
//goland:noinspection GoAssignmentToReceiver
func (b BitBoard) Shift(dx int, dy int) BitBoard {

	//dy = up/down
	if dy > 0 {
		b <<= dy * 8
	}
	if dy < 0 {
		b >>= (-dy) * 8
	}

	//dx = left / right
	if dx > 0 {
		for i := 0; i < dx; i++ {
			b = b.OneEast()
		}
	}
	if dx < 0 {
		for i := 0; i < -dx; i++ {
			b = b.OneWest()
		}
	}

	return b
}

// MirrorVertical returns board with ranks (rows) in reverse order
func (b BitBoard) MirrorVertical() BitBoard {
	result := Empty

	result |= (b >> 56) & Rank1
	result |= ((b >> 48) & Rank1) << 8
	result |= ((b >> 40) & Rank1) << 16
	result |= ((b >> 32) & Rank1) << 24
	result |= ((b >> 24) & Rank1) << 32
	result |= ((b >> 16) & Rank1) << 40
	result |= ((b >> 8) & Rank1) << 48
	result |= (b & Rank1) << 56

	return result
}

// MirrorHorizontal returns board which mirrors the bitboard horizontally
//goland:noinspection GoAssignmentToReceiver
func (b BitBoard) MirrorHorizontal() BitBoard {
	const k1 = BitBoard(0x5555555555555555)
	const k2 = BitBoard(0x3333333333333333)
	const k4 = BitBoard(0x0f0f0f0f0f0f0f0f)

	b = ((b >> 1) & k1) | ((b & k1) << 1)
	b = ((b >> 2) & k2) | ((b & k2) << 2)
	b = ((b >> 4) & k4) | ((b & k4) << 4)

	return b
}

// FlipA1H8 returns board flipped around A1H8 diagonal
func (b BitBoard) FlipA1H8() BitBoard {
	const k1 = BitBoard(0x5500550055005500)
	const k2 = BitBoard(0x3333000033330000)
	const k4 = BitBoard(0x0f0f0f0f00000000)

	var t = k4 & (b ^ (b << 28))

	b ^= t ^ (t >> 28)
	t = k2 & (b ^ (b << 14))
	b ^= t ^ (t >> 14)
	t = k1 & (b ^ (b << 7))
	b ^= t ^ (t >> 7)

	return b
}

// ToIndices returns array of indices set in the BitBoard
func (b BitBoard) ToIndices() []index.Index {
	popCount := b.PopCount()

	result := make([]index.Index, popCount)
	i := 0

	for b > 0 {
		result[i] = index.Index(b.BitPop())
		i++
	}

	return result
}

func (b BitBoard) String() string {
	reversedRanks := b.MirrorVertical()
	var result bytes.Buffer

	result.WriteString(header)

	for i := 0; i < BitWidth; i++ {
		if (i % 8) == 0 {
			if i > 0 {
				//print right column digit
				result.WriteString(strconv.Itoa(9 - (i / 8)))
				result.WriteString("\n")
			}

			//print left column digit
			result.WriteString(strconv.Itoa(8 - (i / 8)))
			result.WriteString(" ")
		}

		if reversedRanks&(1<<i) != 0 {
			result.WriteString("x ")
		} else {
			result.WriteString("- ")
		}
	}

	result.WriteString("1\n") //last right column digit
	result.WriteString(header)

	return result.String()
}
