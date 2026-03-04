package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPiece_String(t *testing.T) {
	assert.Equal(t, "K", King.String(White))
	assert.Equal(t, "k", King.String(Black))
	assert.Equal(t, "Q", Queen.String(White))
	assert.Equal(t, "q", Queen.String(Black))
	assert.Equal(t, "R", Rook.String(White))
	assert.Equal(t, "r", Rook.String(Black))
	assert.Equal(t, "B", Bishop.String(White))
	assert.Equal(t, "b", Bishop.String(Black))
	assert.Equal(t, "N", Knight.String(White))
	assert.Equal(t, "n", Knight.String(Black))
	assert.Equal(t, "P", Pawn.String(White))
	assert.Equal(t, "p", Pawn.String(Black))
}

func TestPieceFromNotation(t *testing.T) {
	tests := []struct {
		input string
		piece Piece
		color Color
	}{
		{"K", King, White},
		{"k", King, Black},
		{"Q", Queen, White},
		{"q", Queen, Black},
		{"R", Rook, White},
		{"r", Rook, Black},
		{"B", Bishop, White},
		{"b", Bishop, Black},
		{"N", Knight, White},
		{"n", Knight, Black},
		{"P", Pawn, White},
		{"p", Pawn, Black},
		{"X", NoPiece, White}, // unknown character
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			piece, color := PieceFromNotation(tt.input)
			assert.Equal(t, tt.piece, piece)
			assert.Equal(t, tt.color, color)
		})
	}
}

func TestColor_String(t *testing.T) {
	assert.Equal(t, "w", White.String())
	assert.Equal(t, "b", Black.String())
}
