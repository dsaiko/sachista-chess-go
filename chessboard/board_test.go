package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"saiko.cz/sachista/bitboard"
)

func TestBoard_RemoveCastling(t *testing.T) {
	b := EmptyBoard()

	assert.Equal(t, CastlingNone, b.Castling[White])
	assert.Equal(t, CastlingNone, b.Castling[Black])

	b.Castling[White] = CastlingBothSides
	b.RemovedCastling(White, CastlingQueenSide)

	assert.Equal(t, CastlingKingSide, b.Castling[White])
}

func TestBoard_Stats(t *testing.T) {
	b := StandardBoard()

	assert.Equal(t, 32, b.AllPieces().PopCount())
	assert.Equal(t, 16, b.OpponentPieces().PopCount())
	assert.Equal(t, 64-16, b.BoardAvailableToAttack().PopCount())

	assert.Equal(t, bitboard.IndexE1, b.MyKingIndex())
	assert.Equal(t, bitboard.IndexE8, b.OpponentKingIndex())
}

func TestBoard_UpdateZobrist(t *testing.T) {
	b1 := StandardBoard()
	b2 := StandardBoard()

	assert.Equal(t, b1.ZobristHash, b2.ZobristHash)

	// no halfmove counters are relevant
	b2.HalfMoveClock = 99
	b2.FullMoveNumber = 99
	b2.ZobristHash = b2.Hash()
	assert.Equal(t, b1.ZobristHash, b2.ZobristHash)

	b2.Pieces[White][Pawn] |= bitboard.BoardA3
	b2.ZobristHash = b2.Hash()
	assert.NotEqual(t, b1.ZobristHash, b2.ZobristHash)
}
