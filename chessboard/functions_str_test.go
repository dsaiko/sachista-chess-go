package chessboard

import (
	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/bitboard"
	"testing"
)

func TestBoard_String(t *testing.T) {
	b := StandardBoard()

	b.Pieces[White][Queen] = bitboard.D5

	expected := `  a b c d e f g h
8 r n b q k b n r 8
7 p p p p p p p p 7
6 - - - - - - - - 6
5 - - - Q - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 P P P P P P P P 2
1 R N B - K B N R 1
  a b c d e f g h
`
	assert.Equal(t, expected, b.String())
}

func TestStandardBoard(t *testing.T) {
	b := StandardBoard()

	assert.Equal(t, 16, b.PiecesOfColor(White).PopCount())
	assert.Equal(t, 16, b.PiecesOfColor(Black).PopCount())
	assert.Equal(t, 32, b.AllPieces().PopCount())

	expected := `  a b c d e f g h
8 r n b q k b n r 8
7 p p p p p p p p 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 P P P P P P P P 2
1 R N B Q K B N R 1
  a b c d e f g h
`
	assert.Equal(t, expected, b.String())

	hash := b.ZobristHash
	assert.NotZero(t, hash)

	b = Empty()
	hash = b.ZobristHash
	assert.Zero(t, hash)

	b.UpdateZobrist()
	hash = b.ZobristHash
	assert.Zero(t, hash)
}
