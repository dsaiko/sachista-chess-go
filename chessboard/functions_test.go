package chessboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_RemoveCastling(t *testing.T) {
	b := Empty()

	assert.Equal(t, CastlingNone, b.Castling[White])
	assert.Equal(t, CastlingNone, b.Castling[Black])

	b.Castling[White] = CastlingBothSides
	b.RemoveCastling(White, CastlingQueenSide)

	assert.Equal(t, CastlingKingSide, b.Castling[White])
}
