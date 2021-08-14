package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/bitboard"

	"saiko.cz/sachista/index"
)

func TestBoard_ToFEN(t *testing.T) {
	b := Standard()
	b2 := FromFEN(b.ToFEN())
	b3 := FromFEN(StandardBoardFEN)

	assert.Equal(t, b.ZobristHash, b2.ZobristHash)
	assert.Equal(t, b.String(), b2.String())
	assert.Equal(t, b.ToFEN(), b2.ToFEN())
	assert.Equal(t, b.ToFEN(), StandardBoardFEN)
	assert.Equal(t, b.ZobristHash, b3.ZobristHash)

	assert.Equal(t, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", b.ToFEN())

	// incomplete FEN
	b2 = FromFEN("8/1K6/1Q6/8/5r2/4rk2/8/8 b - -")
	assert.Equal(t, 5, b2.BoardOfAllPieces().PopCount())
	assert.Equal(t, Black, b2.NextMove)
	assert.Equal(t, CastlingNone, b2.Castling[White])
	assert.Equal(t, CastlingNone, b2.Castling[Black])
	assert.Equal(t, index.Index(0), b2.EnPassantTarget)
	assert.Equal(t, 0, b2.HalfMoveClock)
	assert.Equal(t, 1, b2.FullMoveNumber)

	b2 = FromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQq a2 14 33")
	assert.Equal(t, 32, b2.BoardOfAllPieces().PopCount())
	assert.Equal(t, White, b2.NextMove)
	assert.Equal(t, CastlingBothSides, b2.Castling[White])
	assert.Equal(t, CastlingQueenSide, b2.Castling[Black])
	assert.Equal(t, index.A2, b2.EnPassantTarget)
	assert.Equal(t, 14, b2.HalfMoveClock)
	assert.Equal(t, 33, b2.FullMoveNumber)

	b2 = FromFEN("7B/6B1/5B2/4B3/3B4/2B5/1B6/B7 w - - 0 1")
	assert.Equal(t, bitboard.A1H8[7], b2.BoardOfAllPieces())
}

func TestBoard_ToFEN_Invalid(t *testing.T) {
	b2 := FromFEN("XXXXX")
	assert.Equal(t, uint64(0), b2.ZobristHash)
	assert.Equal(t, bitboard.Empty, b2.BoardOfAllPieces())

	b2 = FromFEN("")
	assert.Equal(t, uint64(0), b2.ZobristHash)
	assert.Equal(t, bitboard.Empty, b2.BoardOfAllPieces())
}
