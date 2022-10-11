package bitboard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitBoard_BitPop(t *testing.T) {
	b := UniverseBoard

	count := 64
	var index Index
	for i := 0; i < 64; i++ {
		assert.Equal(t, count, b.PopCount())
		index, b = b.BitPop()
		assert.Equal(t, Index(i), index)
		count--
	}
}

func TestFromIndex(t *testing.T) {
	b := BoardFromIndices(
		IndexA1, IndexA2, IndexA3, IndexA4, IndexA5, IndexA6, IndexA7, IndexA8,
		IndexH1, IndexH2, IndexH3, IndexH4, IndexH5, IndexH6, IndexH7, IndexH8,
		IndexB1, IndexC1, IndexD1, IndexE1, IndexF1, IndexG1,
		IndexB8, IndexC8, IndexD8, IndexE8, IndexF8, IndexG8,
	)
	assert.Equal(t, BoardFrame, b)

	b = BoardFromNotation(
		"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8",
		"H1", "H2", "H3", "H4", "H5", "H6", "H7", "H8",
		"B1", "C1", "D1", "E1", "F1", "G1",
		"B8", "C8", "D8", "E8", "F8", "G8",
	)
	assert.Equal(t, BoardFrame, b)
}

func TestFromIndex1(t *testing.T) {
	for i := 0; i < 64; i++ {
		assert.Equal(t, BoardFields[i], BoardFromIndex(Index(i)))
	}
}

func TestBitBoard_ToIndices(t *testing.T) {
	board := BoardFromNotation("a1", "a8", "h1", "h8", "f7")
	indices := board.ToIndices()

	contains := func(i Index) bool {
		for _, v := range indices {
			if v == i {
				return true
			}
		}
		return false
	}

	assert.Equal(t, 5, len(indices))
	assert.True(t, contains(IndexA1))
	assert.True(t, contains(IndexA8))
	assert.True(t, contains(IndexH1))
	assert.True(t, contains(IndexH8))
	assert.True(t, contains(IndexF7))

	board = EmptyBoard
	indices = board.ToIndices()
	assert.Equal(t, 0, len(indices))
}

func TestBitBoard_OneWest(t *testing.T) {
	assert.Equal(t, EmptyBoard, BoardA1.ShiftedOneWest())
	assert.Equal(t, EmptyBoard, BoardA8.ShiftedOneWest())
	assert.Equal(t, BoardG1, BoardH1.ShiftedOneWest())
	assert.Equal(t, BoardG8, BoardH8.ShiftedOneWest())
}

func TestBitBoard_OneNorth(t *testing.T) {
	assert.Equal(t, BoardA2, BoardA1.ShiftedOneNorth())
	assert.Equal(t, EmptyBoard, BoardA8.ShiftedOneNorth())
	assert.Equal(t, BoardH2, BoardH1.ShiftedOneNorth())
	assert.Equal(t, EmptyBoard, BoardH8.ShiftedOneNorth())
}

func TestBitBoard_OneEast(t *testing.T) {
	assert.Equal(t, BoardB1, BoardA1.ShiftedOneEast())
	assert.Equal(t, BoardB8, BoardA8.ShiftedOneEast())
	assert.Equal(t, EmptyBoard, BoardH1.ShiftedOneEast())
	assert.Equal(t, EmptyBoard, BoardH8.ShiftedOneEast())
}

func TestBitBoard_OneSouth(t *testing.T) {
	assert.Equal(t, EmptyBoard, BoardA1.ShiftedOneSouth())
	assert.Equal(t, BoardA7, BoardA8.ShiftedOneSouth())
	assert.Equal(t, EmptyBoard, BoardH1.ShiftedOneSouth())
	assert.Equal(t, BoardH7, BoardH8.ShiftedOneSouth())
}

func TestBitBoard_OneSouthEast(t *testing.T) {
	assert.Equal(t, EmptyBoard, BoardA1.ShiftedOneSouthEast())
	assert.Equal(t, BoardB7, BoardA8.ShiftedOneSouthEast())
	assert.Equal(t, EmptyBoard, BoardH1.ShiftedOneSouthEast())
	assert.Equal(t, EmptyBoard, BoardH8.ShiftedOneSouthEast())
}

func TestBitBoard_OneSouthWest(t *testing.T) {
	assert.Equal(t, EmptyBoard, BoardA1.ShiftedOneSouthWest())
	assert.Equal(t, EmptyBoard, BoardA8.ShiftedOneSouthWest())
	assert.Equal(t, EmptyBoard, BoardH1.ShiftedOneSouthWest())
	assert.Equal(t, BoardG7, BoardH8.ShiftedOneSouthWest())
}

func TestBitBoard_OneNorthEast(t *testing.T) {
	assert.Equal(t, BoardB2, BoardA1.ShiftedOneNorthEast())
	assert.Equal(t, EmptyBoard, BoardA8.ShiftedOneNorthEast())
	assert.Equal(t, EmptyBoard, BoardH1.ShiftedOneNorthEast())
	assert.Equal(t, EmptyBoard, BoardH8.ShiftedOneNorthEast())
}

func TestBitBoard_OneNorthWest(t *testing.T) {
	assert.Equal(t, EmptyBoard, BoardA1.ShiftedOneNorthWest())
	assert.Equal(t, EmptyBoard, BoardA8.ShiftedOneNorthWest())
	assert.Equal(t, BoardG2, BoardH1.ShiftedOneNorthWest())
	assert.Equal(t, EmptyBoard, BoardH8.ShiftedOneNorthWest())
}

func TestBitBoard_FlipA1H8(t *testing.T) {
	diag1 := ^BoardA1H8[7]
	diag2 := ^BoardA1H8[5]

	assert.Equal(t, diag1, diag1.FlippedA1H8())
	assert.Equal(t, diag2, diag2.FlippedA1H8().FlippedA1H8())
	assert.NotEqual(t, diag2, diag2.FlippedA1H8())

	b1 := BoardA1 | BoardH1 | BoardH8
	b2 := BoardA1 | BoardA8 | BoardH8

	assert.Equal(t, b2, b1.FlippedA1H8())
}

func TestBitBoard_MirrorHorizontal(t *testing.T) {
	b2 := BoardA1H8[7]

	assert.Equal(t, BoardFrame, BoardFrame.MirroredHorizontal())
	assert.Equal(t, b2, b2.MirroredHorizontal().MirroredHorizontal())
	assert.NotEqual(t, b2, b2.MirroredHorizontal())
	assert.Equal(t, BoardFiles[1], BoardFiles[6].MirroredHorizontal())
}

func TestBitBoard_MirrorVertical(t *testing.T) {
	b2 := BoardA1H8[7]

	assert.Equal(t, BoardFrame, BoardFrame.MirroredVertical())
	assert.Equal(t, b2, b2.MirroredVertical().MirroredVertical())
	assert.NotEqual(t, b2, b2.MirroredVertical())
	assert.Equal(t, BoardRanks[1], BoardRanks[6].MirroredVertical())
}

func TestBitBoard_String(t *testing.T) {
	b := BoardFrame | BoardC3
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
