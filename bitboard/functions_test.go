package bitboard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"saiko.cz/sachista/index"
)

func TestBitBoard_BitPop(t *testing.T) {
	b := Universe

	count := 64
	for i := 0; i < 64; i++ {
		assert.Equal(t, count, b.PopCount())
		assert.Equal(t, i, b.BitPop())
		count--
	}
}

func TestFromIndex(t *testing.T) {
	b := FromIndex(
		index.A1, index.A2, index.A3, index.A4, index.A5, index.A6, index.A7, index.A8,
		index.H1, index.H2, index.H3, index.H4, index.H5, index.H6, index.H7, index.H8,
		index.B1, index.C1, index.D1, index.E1, index.F1, index.G1,
		index.B8, index.C8, index.D8, index.E8, index.F8, index.G8,
	)
	assert.Equal(t, Frame, b)

	b = FromNotation(
		"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8",
		"H1", "H2", "H3", "H4", "H5", "H6", "H7", "H8",
		"B1", "C1", "D1", "E1", "F1", "G1",
		"B8", "C8", "D8", "E8", "F8", "G8",
	)
	assert.Equal(t, Frame, b)

}

func TestFromIndex1(t *testing.T) {
	for i := 0; i < 64; i++ {
		assert.Equal(t, Fields[i], FromIndex1(index.Index(i)))
	}
}

func TestBitBoard_ToIndices(t *testing.T) {
	board := FromNotation("a1", "a8", "h1", "h8", "f7")
	indices := board.ToIndices()

	contains := func(i index.Index) bool {
		for _, v := range indices {
			if v == i {
				return true
			}
		}
		return false
	}

	assert.Equal(t, 5, len(indices))
	assert.True(t, contains(index.A1))
	assert.True(t, contains(index.A8))
	assert.True(t, contains(index.H1))
	assert.True(t, contains(index.H8))
	assert.True(t, contains(index.F7))

	board = Empty
	indices = board.ToIndices()
	assert.Equal(t, 0, len(indices))
}

func TestBitBoard_OneWest(t *testing.T) {
	assert.Equal(t, Empty, A1.OneWest())
	assert.Equal(t, Empty, A8.OneWest())
	assert.Equal(t, G1, H1.OneWest())
	assert.Equal(t, G8, H8.OneWest())
}

func TestBitBoard_OneNorth(t *testing.T) {
	assert.Equal(t, A2, A1.OneNorth())
	assert.Equal(t, Empty, A8.OneNorth())
	assert.Equal(t, H2, H1.OneNorth())
	assert.Equal(t, Empty, H8.OneNorth())
}

func TestBitBoard_OneEast(t *testing.T) {
	assert.Equal(t, B1, A1.OneEast())
	assert.Equal(t, B8, A8.OneEast())
	assert.Equal(t, Empty, H1.OneEast())
	assert.Equal(t, Empty, H8.OneEast())
}

func TestBitBoard_OneSouth(t *testing.T) {
	assert.Equal(t, Empty, A1.OneSouth())
	assert.Equal(t, A7, A8.OneSouth())
	assert.Equal(t, Empty, H1.OneSouth())
	assert.Equal(t, H7, H8.OneSouth())
}

func TestBitBoard_OneSouthEast(t *testing.T) {
	assert.Equal(t, Empty, A1.OneSouthEast())
	assert.Equal(t, B7, A8.OneSouthEast())
	assert.Equal(t, Empty, H1.OneSouthEast())
	assert.Equal(t, Empty, H8.OneSouthEast())
}

func TestBitBoard_OneSouthWest(t *testing.T) {
	assert.Equal(t, Empty, A1.OneSouthWest())
	assert.Equal(t, Empty, A8.OneSouthWest())
	assert.Equal(t, Empty, H1.OneSouthWest())
	assert.Equal(t, G7, H8.OneSouthWest())
}

func TestBitBoard_OneNorthEast(t *testing.T) {
	assert.Equal(t, B2, A1.OneNorthEast())
	assert.Equal(t, Empty, A8.OneNorthEast())
	assert.Equal(t, Empty, H1.OneNorthEast())
	assert.Equal(t, Empty, H8.OneNorthEast())
}

func TestBitBoard_OneNorthWest(t *testing.T) {
	assert.Equal(t, Empty, A1.OneNorthWest())
	assert.Equal(t, Empty, A8.OneNorthWest())
	assert.Equal(t, G2, H1.OneNorthWest())
	assert.Equal(t, Empty, H8.OneNorthWest())
}

func TestBitBoard_FlipA1H8(t *testing.T) {
	diag1 := ^A1H8[7]
	diag2 := ^A1H8[5]

	assert.Equal(t, diag1, diag1.FlipA1H8())
	assert.Equal(t, diag2, diag2.FlipA1H8().FlipA1H8())
	assert.NotEqual(t, diag2, diag2.FlipA1H8())

	b1 := A1 | H1 | H8
	b2 := A1 | A8 | H8

	assert.Equal(t, b2, b1.FlipA1H8())
}

func TestBitBoard_MirrorHorizontal(t *testing.T) {
	b2 := A1H8[7]

	assert.Equal(t, Frame, Frame.MirrorHorizontal())
	assert.Equal(t, b2, b2.MirrorHorizontal().MirrorHorizontal())
	assert.NotEqual(t, b2, b2.MirrorHorizontal())
	assert.Equal(t, File[1], File[6].MirrorHorizontal())
}

func TestBitBoard_MirrorVertical(t *testing.T) {
	b2 := A1H8[7]

	assert.Equal(t, Frame, Frame.MirrorVertical())
	assert.Equal(t, b2, b2.MirrorVertical().MirrorVertical())
	assert.NotEqual(t, b2, b2.MirrorVertical())
	assert.Equal(t, Rank[1], Rank[6].MirrorVertical())
}

func TestBitBoard_String(t *testing.T) {
	b := Frame | C3
	expected := `  a b c d e f g h
8 x x x x x x x x 8
7 x - - - - - - x 7
6 x - - - - - - x 6
5 x - - - - - - x 5
4 x - - - - - - x 4
3 x - x - - - - x 3
2 x - - - - - - x 2
1 x x x x x x x x 1
  a b c d e f g h
`

	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
