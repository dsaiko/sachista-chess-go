package engine

import (
	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/board"
	"saiko.cz/sachista/index"
	"testing"
)

func TestBoardFromIndex(t *testing.T) {
	b := BoardFromIndex(
		index.A1, index.A2, index.A3, index.A4, index.A5, index.A6, index.A7, index.A8,
		index.H1, index.H2, index.H3, index.H4, index.H5, index.H6, index.H7, index.H8,
		index.B1, index.C1, index.D1, index.E1, index.F1, index.G1,
		index.B8, index.C8, index.D8, index.E8, index.F8, index.G8,
	)
	assert.Equal(t, board.Frame, b)

	b = BoardFromNotation(
		"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8",
		"H1", "H2", "H3", "H4", "H5", "H6", "H7", "H8",
		"B1", "C1", "D1", "E1", "F1", "G1",
		"B8", "C8", "D8", "E8", "F8", "G8",
	)
	assert.Equal(t, board.Frame, b)

}

func TestIndexToBitBoard(t *testing.T) {

	for i := 0; i < 64; i++ {
		assert.Equal(t, board.Fields[i], IndexToBitBoard(index.Index(i)))
	}
}
