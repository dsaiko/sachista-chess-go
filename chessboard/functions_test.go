package chessboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_RemoveCastling(t *testing.T) {
	b := Empty()

	assert.Equal(t, CastlingNone, b.castling[White])
	assert.Equal(t, CastlingNone, b.castling[Black])

	b.castling[White] = CastlingBothSides
	b.RemoveCastling(White, CastlingQueenSide)

	assert.Equal(t, CastlingKingSide, b.castling[White])
}
