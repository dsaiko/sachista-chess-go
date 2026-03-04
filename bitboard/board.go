package bitboard

import (
	"bytes"
	"math/bits"
	"strconv"
)

// Board represents an 8x8 chess board as a 64-bit unsigned integer,
// where each bit corresponds to a square (bit 0 = a1, bit 63 = h8).
type Board uint64

// PopCount returns the number of bits set in the bitboard
func (b Board) PopCount() int {
	return bits.OnesCount64(uint64(b))
}

// BitScan returns the index of the first set bit, or 64 if no bits are set.
func (b Board) BitScan() Index {
	return Index(bits.TrailingZeros64(uint64(b)))
}

// BitPop returns the index of the first set bit and a new board with that bit cleared.
func (b Board) BitPop() (Index, Board) {
	return b.BitScan(), b & (b - 1)
}

// ShiftedOneNorth returns a board with all bits shifted one rank up.
func (b Board) ShiftedOneNorth() Board {
	return b << 8
}

// ShiftedOneSouth returns a board with all bits shifted one rank down.
func (b Board) ShiftedOneSouth() Board {
	return b >> 8
}

// ShiftedOneEast returns a board with all bits shifted one file right, masking off wrap-around.
func (b Board) ShiftedOneEast() Board {
	return (b << 1) & ^BoardFileA
}

// ShiftedOneNorthEast returns a board with all bits shifted one square diagonally (north-east).
func (b Board) ShiftedOneNorthEast() Board {
	return (b << 9) & ^BoardFileA
}

// ShiftedOneSouthEast returns a board with all bits shifted one square diagonally (south-east).
func (b Board) ShiftedOneSouthEast() Board {
	return (b >> 7) & ^BoardFileA
}

// ShiftedOneWest returns a board with all bits shifted one file left, masking off wrap-around.
func (b Board) ShiftedOneWest() Board {
	return (b >> 1) & ^BoardFileH
}

// ShiftedOneSouthWest returns a board with all bits shifted one square diagonally (south-west).
func (b Board) ShiftedOneSouthWest() Board {
	return (b >> 9) & ^BoardFileH
}

// ShiftedOneNorthWest returns a board with all bits shifted one square diagonally (north-west).
func (b Board) ShiftedOneNorthWest() Board {
	return (b << 7) & ^BoardFileH
}

// Shifted returns a board with all bits shifted by dx files (positive=east) and dy ranks (positive=north).
func (b Board) Shifted(dx int, dy int) Board {
	if dy > 0 {
		b <<= dy * 8
	}
	if dy < 0 {
		b >>= (-dy) * 8
	}

	if dx > 0 {
		for range dx {
			b = b.ShiftedOneEast()
		}
	}
	if dx < 0 {
		for i := 0; i < -dx; i++ {
			b = b.ShiftedOneWest()
		}
	}

	return b
}

// MirroredVertical returns a board with ranks (rows) in reverse order.
func (b Board) MirroredVertical() Board {
	result := EmptyBoard

	result |= (b >> 56) & BoardRank1         //nolint:revive
	result |= ((b >> 48) & BoardRank1) << 8  //nolint:revive
	result |= ((b >> 40) & BoardRank1) << 16 //nolint:revive
	result |= ((b >> 32) & BoardRank1) << 24 //nolint:revive
	result |= ((b >> 24) & BoardRank1) << 32 //nolint:revive
	result |= ((b >> 16) & BoardRank1) << 40 //nolint:revive
	result |= ((b >> 8) & BoardRank1) << 48  //nolint:revive
	result |= (b & BoardRank1) << 56         //nolint:revive

	return result
}

// MirroredHorizontal returns a board with files (columns) in reverse order.
// Uses the delta swap technique with bit-manipulation masks for alternating bit groups.
func (b Board) MirroredHorizontal() Board {
	const k1 = Board(0x5555555555555555) // alternating single bits
	const k2 = Board(0x3333333333333333) // alternating bit pairs
	const k4 = Board(0x0f0f0f0f0f0f0f0f) // alternating nibbles

	b = ((b >> 1) & k1) | ((b & k1) << 1)
	b = ((b >> 2) & k2) | ((b & k2) << 2)
	b = ((b >> 4) & k4) | ((b & k4) << 4)

	return b
}

// FlippedA1H8 returns a board flipped along the a1-h8 diagonal (transpose).
func (b Board) FlippedA1H8() Board {
	const k1 = Board(0x5500550055005500)
	const k2 = Board(0x3333000033330000)
	const k4 = Board(0x0f0f0f0f00000000)

	t := k4 & (b ^ (b << 28)) //nolint:revive

	b ^= t ^ (t >> 28)       //nolint:revive
	t = k2 & (b ^ (b << 14)) //nolint:revive
	b ^= t ^ (t >> 14)       //nolint:revive
	t = k1 & (b ^ (b << 7))
	b ^= t ^ (t >> 7)

	return b
}

// ToIndices returns a slice of all set bit positions as Index values.
func (b Board) ToIndices() []Index {
	popCount := b.PopCount()

	result := make([]Index, popCount)
	i := 0
	for b != EmptyBoard {
		result[i], b = b.BitPop()
		i++
	}
	return result
}

// String returns an ASCII representation of the board with rank/file labels.
func (b Board) String() string {
	reversedRanks := b.MirroredVertical()
	var buffer bytes.Buffer

	buffer.WriteString(BoardHeader)

	for i := range NumberOfSquares {
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

// BoardFromIndices returns a board with bits set at all specified index positions.
func BoardFromIndices(indices ...Index) Board {
	b := EmptyBoard
	for i := range indices {
		b |= BoardFromIndex(indices[i])
	}
	return b
}

// BoardFromIndex returns a board with a single bit set at the given index.
func BoardFromIndex(i Index) Board {
	return 1 << i
}

// BoardFromNotation returns a board with bits set at positions specified by algebraic notation strings (e.g. "e4", "d7").
func BoardFromNotation(notations ...string) Board {
	b := EmptyBoard
	for i := range notations {
		b |= BoardFromIndex(IndexFromNotation(notations[i]))
	}
	return b
}
