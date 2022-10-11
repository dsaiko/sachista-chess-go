package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/bitboard"
)

func TestBoard_String(t *testing.T) {
	b := StandardBoard()

	b.Pieces[White][Queen] = bitboard.BoardD5

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

	assert.Equal(t, 16, b.PiecesByColor(White).PopCount())
	assert.Equal(t, 16, b.PiecesByColor(Black).PopCount())
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

	assert.NotZero(t, b.ZobristHash)

	hash := b.ZobristHash
	assert.NotZero(t, hash)

	b = EmptyBoard()
	hash = b.ZobristHash
	assert.Zero(t, hash)

	b.ZobristHash = b.Hash()
	assert.Zero(t, b.ZobristHash)
}

func TestFromString(t *testing.T) {
	b := BoardFromFEN("8/1K6/1Q6/8/5r2/4rk2/8/8 w - a2")
	assert.Equal(t, bitboard.IndexA2, b.EnPassantTarget)

	b2 := BoardFromFEN(b.ToFEN())
	assert.Equal(t, b, b2)
	assert.Equal(t, b.ZobristHash, b2.ZobristHash)

	b2 = FromString(b.String())
	assert.NotEqual(t, b, b2)
	assert.NotEqual(t, b.ZobristHash, b2.ZobristHash)
	b2.EnPassantTarget = bitboard.IndexA2
	b2.ZobristHash = b2.Hash()
	assert.Equal(t, b, b2)
	assert.Equal(t, b.ZobristHash, b2.ZobristHash)
}
