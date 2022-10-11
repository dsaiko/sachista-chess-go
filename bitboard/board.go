package bitboard

import (
	"bytes"
	"math/bits"
	"strconv"
)

// Board is representing 8x8 (64 bit) bitboard where each bit represent existing piece on the given position
type Board uint64

// PopCount returns the number of bits set in the bitboard
func (b Board) PopCount() int {
	return bits.OnesCount64(uint64(b))
}

// BitScan returns the index of first 1 bit or 64 if no bits are set
func (b Board) BitScan() Index {
	return Index(bits.TrailingZeros64(uint64(b)))
}

// BitPop returns index of first set bit and resets this bit in the bitboard
func (b Board) BitPop() (Index, Board) {
	return b.BitScan(), b & (b - 1)
}

// ShiftedOneNorth shifts all existing Board pieces by one
func (b Board) ShiftedOneNorth() Board {
	return b << 8
}

// ShiftedOneSouth  shifts all existing Board pieces by one
func (b Board) ShiftedOneSouth() Board {
	return b >> 8
}

// ShiftedOneEast  shifts all existing Board pieces by one
func (b Board) ShiftedOneEast() Board {
	return (b << 1) & ^BoardFileA
}

// ShiftedOneNorthEast  shifts all existing Board pieces by one
func (b Board) ShiftedOneNorthEast() Board {
	return (b << 9) & ^BoardFileA
}

// ShiftedOneSouthEast  shifts all existing Board pieces by one
func (b Board) ShiftedOneSouthEast() Board {
	return (b >> 7) & ^BoardFileA
}

// ShiftedOneWest  shifts all existing Board pieces by one
func (b Board) ShiftedOneWest() Board {
	return (b >> 1) & ^BoardFileH
}

// ShiftedOneSouthWest  shifts all existing Board pieces by one
func (b Board) ShiftedOneSouthWest() Board {
	return (b >> 9) & ^BoardFileH
}

// ShiftedOneNorthWest  shifts all existing Board pieces by one
func (b Board) ShiftedOneNorthWest() Board {
	return (b << 7) & ^BoardFileH
}

// Shifted shifts all existing Board pieces by multiple steps
func (b Board) Shifted(dx int, dy int) Board {
	// dy = up/down
	if dy > 0 {
		//goland:noinspection GoAssignmentToReceiver
		b <<= dy * 8
	}
	if dy < 0 {
		//goland:noinspection GoAssignmentToReceiver
		b >>= (-dy) * 8
	}

	// dx = left / right
	if dx > 0 {
		for i := 0; i < dx; i++ {
			//goland:noinspection GoAssignmentToReceiver
			b = b.ShiftedOneEast()
		}
	}
	if dx < 0 {
		for i := 0; i < -dx; i++ {
			//goland:noinspection GoAssignmentToReceiver
			b = b.ShiftedOneWest()
		}
	}

	return b
}

// MirroredVertical returns bitboard with ranks (rows) in reverse order
func (b Board) MirroredVertical() Board {
	result := EmptyBoard

	result |= (b >> 56) & BoardRank1
	result |= ((b >> 48) & BoardRank1) << 8
	result |= ((b >> 40) & BoardRank1) << 16
	result |= ((b >> 32) & BoardRank1) << 24
	result |= ((b >> 24) & BoardRank1) << 32
	result |= ((b >> 16) & BoardRank1) << 40
	result |= ((b >> 8) & BoardRank1) << 48
	result |= (b & BoardRank1) << 56

	return result
}

// MirroredHorizontal returns bitboard which mirrors the bitboard horizontally
//
//goland:noinspection GoAssignmentToReceiver
func (b Board) MirroredHorizontal() Board {
	const k1 = Board(0x5555555555555555)
	const k2 = Board(0x3333333333333333)
	const k4 = Board(0x0f0f0f0f0f0f0f0f)

	b = ((b >> 1) & k1) | ((b & k1) << 1)
	b = ((b >> 2) & k2) | ((b & k2) << 2)
	b = ((b >> 4) & k4) | ((b & k4) << 4)

	return b
}

// FlippedA1H8 returns bitboard flipped around A1H8 diagonal
func (b Board) FlippedA1H8() Board {
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
func (b Board) ToIndices() []Index {
	popCount := b.PopCount()

	result := make([]Index, popCount)
	i := 0
	for b != EmptyBoard {
		//goland:noinspection GoAssignmentToReceiver
		result[i], b = b.BitPop()
		i++
	}
	return result
}

func (b Board) String() string {
	reversedRanks := b.MirroredVertical()
	var buffer bytes.Buffer

	buffer.WriteString(BoardHeader)

	for i := 0; i < NumberOfSquares; i++ {
		if (i % 8) == 0 {
			if i > 0 {
				// print right column digit
				buffer.WriteString(strconv.Itoa(9 - (i / 8)))
				buffer.WriteString("\n")
			}

			// print left column digit
			buffer.WriteString(strconv.Itoa(8 - (i / 8)))
			buffer.WriteString(" ")
		}

		if reversedRanks&(1<<i) != 0 {
			buffer.WriteString("x ")
		} else {
			buffer.WriteString("- ")
		}
	}

	buffer.WriteString("1\n") // last right column digit
	buffer.WriteString(BoardHeader)

	return buffer.String()
}

// BoardFromIndices returns Board filled with pieces defined by index parameters
func BoardFromIndices(indices ...Index) Board {
	b := EmptyBoard
	for i := range indices {
		b |= BoardFromIndex(indices[i])
	}
	return b
}

// BoardFromIndex returns Board of one piece index
func BoardFromIndex(i Index) Board {
	return 1 << i
}

// BoardFromNotation returns Board filled with pieces defined by notation parameters
func BoardFromNotation(notations ...string) Board {
	b := EmptyBoard
	for i := range notations {
		b |= BoardFromIndex(IndexFromNotation(notations[i]))
	}
	return b
}
