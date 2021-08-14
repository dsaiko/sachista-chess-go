package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/bitboard"

	"saiko.cz/sachista/index"
)

func TestBoard_String(t *testing.T) {
	b := Standard()

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
	b := Standard()

	assert.Equal(t, 16, b.BoardOfColor(White).PopCount())
	assert.Equal(t, 16, b.BoardOfColor(Black).PopCount())
	assert.Equal(t, 32, b.BoardOfAllPieces().PopCount())

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

	b = Empty()
	hash = b.ZobristHash
	assert.Zero(t, hash)

	b.ZobristHash = b.ComputeBoardHash()
	assert.Zero(t, b.ZobristHash)
}

func TestFromString(t *testing.T) {
	b := FromFen("8/1K6/1Q6/8/5r2/4rk2/8/8 w - a2")
	assert.Equal(t, index.A2, b.EnPassantTarget)

	b2 := FromFen(b.ToFEN())
	assert.Equal(t, b, b2)
	assert.Equal(t, b.ZobristHash, b2.ZobristHash)

	b2 = FromString(b.String())
	assert.NotEqual(t, b, b2)
	assert.NotEqual(t, b.ZobristHash, b2.ZobristHash)
	b2.EnPassantTarget = index.A2
	b2.ZobristHash = b2.ComputeBoardHash()
	assert.Equal(t, b, b2)
	assert.Equal(t, b.ZobristHash, b2.ZobristHash)
}
