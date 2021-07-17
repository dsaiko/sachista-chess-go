package bitboard

import (
	"bytes"
	"math/bits"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
	"strconv"
)

// PopCount returns the number of bits set in the bitboard
func (b Board) PopCount() int {
	return bits.OnesCount64(uint64(b))
}

// BitScan returns the index of first 1 bit or 64 if no bits are set
func (b Board) BitScan() int {
	return bits.TrailingZeros64(uint64(b))
}

// BitPop returns index of first set bit and resets this bit in the bitboard
func (b *Board) BitPop() int {
	i := b.BitScan()
	*b &= *b - 1

	return i
}

// FromNotation returns Board filled with pieces defined by notation parameters
func FromNotation(notations ...string) Board {
	b := Empty

	for _, n := range notations {
		b |= FromIndex1(index.FromNotation(n))
	}
	return b
}

// FromIndex returns Board filled with pieces defined by index parameters
func FromIndex(indices ...index.Index) Board {
	b := Empty

	for _, i := range indices {
		b |= FromIndex1(i)
	}
	return b
}

// FromIndex1 returns Board of one piece index
func FromIndex1(i index.Index) Board {
	return 1 << i
}

// OneNorth shifts all existing Board pieces by one
func (b Board) OneNorth() Board {
	return b << 8
}

// OneSouth  shifts all existing Board pieces by one
func (b Board) OneSouth() Board {
	return b >> 8
}

// OneEast  shifts all existing Board pieces by one
func (b Board) OneEast() Board {
	return (b << 1) & ^FileA
}

// OneNorthEast  shifts all existing Board pieces by one
func (b Board) OneNorthEast() Board {
	return (b << 9) & ^FileA
}

// OneSouthEast  shifts all existing Board pieces by one
func (b Board) OneSouthEast() Board {
	return (b >> 7) & ^FileA
}

// OneWest  shifts all existing Board pieces by one
func (b Board) OneWest() Board {
	return (b >> 1) & ^FileH
}

// OneSouthWest  shifts all existing Board pieces by one
func (b Board) OneSouthWest() Board {
	return (b >> 9) & ^FileH
}

// OneNorthWest  shifts all existing Board pieces by one
func (b Board) OneNorthWest() Board {
	return (b << 7) & ^FileH
}

// Shift shifts all existing Board pieces by multiple steps
//goland:noinspection GoAssignmentToReceiver
func (b Board) Shift(dx int, dy int) Board {

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

// MirrorVertical returns bitboard with ranks (rows) in reverse order
func (b Board) MirrorVertical() Board {
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

// MirrorHorizontal returns bitboard which mirrors the bitboard horizontally
//goland:noinspection GoAssignmentToReceiver
func (b Board) MirrorHorizontal() Board {
	const k1 = Board(0x5555555555555555)
	const k2 = Board(0x3333333333333333)
	const k4 = Board(0x0f0f0f0f0f0f0f0f)

	b = ((b >> 1) & k1) | ((b & k1) << 1)
	b = ((b >> 2) & k2) | ((b & k2) << 2)
	b = ((b >> 4) & k4) | ((b & k4) << 4)

	return b
}

// FlipA1H8 returns bitboard flipped around A1H8 diagonal
func (b Board) FlipA1H8() Board {
	const k1 = Board(0x5500550055005500)
	const k2 = Board(0x3333000033330000)
	const k4 = Board(0x0f0f0f0f00000000)

	var t = k4 & (b ^ (b << 28))

	b ^= t ^ (t >> 28)
	t = k2 & (b ^ (b << 14))
	b ^= t ^ (t >> 14)
	t = k1 & (b ^ (b << 7))
	b ^= t ^ (t >> 7)

	return b
}

// ToIndices returns array of indices set in the Board
func (b Board) ToIndices() []index.Index {
	popCount := b.PopCount()

	result := make([]index.Index, popCount)
	i := 0

	for b > 0 {
		result[i] = index.Index(b.BitPop())
		i++
	}

	return result
}

func (b Board) String() string {
	reversedRanks := b.MirrorVertical()
	var buffer bytes.Buffer

	buffer.WriteString(BoardHeader)

	for i := 0; i < constants.NumberOfSquares; i++ {
		if (i % 8) == 0 {
			if i > 0 {
				//print right column digit
				buffer.WriteString(strconv.Itoa(9 - (i / 8)))
				buffer.WriteString("\n")
			}

			//print left column digit
			buffer.WriteString(strconv.Itoa(8 - (i / 8)))
			buffer.WriteString(" ")
		}

		if reversedRanks&(1<<i) != 0 {
			buffer.WriteString("x ")
		} else {
			buffer.WriteString("- ")
		}
	}

	buffer.WriteString("1\n") //last right column digit
	buffer.WriteString(BoardHeader)

	return buffer.String()
}
