package chessboard

import (
	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/index"
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

func TestBoard_Stats(t *testing.T) {
	b := StandardBoard()

	assert.Equal(t, 32, b.AllPieces().PopCount())
	assert.Equal(t, 16, b.OpponentPieces().PopCount())
	assert.Equal(t, 64-16, b.BoardAvailable().PopCount())

	assert.Equal(t, index.E1, b.MyKingIndex())
	assert.Equal(t, index.E8, b.OpponentKingIndex())
}

func TestBoard_UpdateZobrist(t *testing.T) {
	b1 := StandardBoard()
	b2 := StandardBoard()

	assert.Equal(t, b1.ZobristHash, b2.ZobristHash)

	//no halfmove counters are relevant
	b2.HalfMoveClock = 99
	b2.FullMoveNumber = 99
	b2.UpdateZobrist()
	assert.Equal(t, b1.ZobristHash, b2.ZobristHash)

	b2.Pieces[White][Pawn] |= bitboard.A3
	b2.UpdateZobrist()
	assert.NotEqual(t, b1.ZobristHash, b2.ZobristHash)
}
