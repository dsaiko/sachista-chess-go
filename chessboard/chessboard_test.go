package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/index"
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
	b := Standard()

	assert.Equal(t, 32, b.BoardOfAllPieces().PopCount())
	assert.Equal(t, 16, b.BoardOfOpponentPieces().PopCount())
	assert.Equal(t, 64-16, b.BoardAvailableToAttack().PopCount())

	assert.Equal(t, index.E1, b.MyKingIndex())
	assert.Equal(t, index.E8, b.OpponentKingIndex())
}

func TestBoard_UpdateZobrist(t *testing.T) {
	b1 := Standard()
	b2 := Standard()

	assert.Equal(t, b1.ZobristHash, b2.ZobristHash)

	//no halfmove counters are relevant
	b2.HalfMoveClock = 99
	b2.FullMoveNumber = 99
	b2.ZobristHash = b2.ComputeZobrist()
	assert.Equal(t, b1.ZobristHash, b2.ZobristHash)

	b2.Pieces[White][Pawn] |= bitboard.A3
	b2.ZobristHash = b2.ComputeZobrist()
	assert.NotEqual(t, b1.ZobristHash, b2.ZobristHash)
}
