package generator

import (
	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
	"testing"
)

func TestMove_String(t *testing.T) {
	move := Move{Piece: chessboard.Pawn, From: index.A2, To: index.A3}
	assert.Equal(t, "a2a3", move.String())

	move = Move{Piece: chessboard.Pawn, From: index.A7, To: index.B8, PromotionPiece: chessboard.Queen}
	assert.Equal(t, "a7b8q", move.String())
}

func testMovesFromString(t *testing.T, expectedCount int, stringBoard string) {
	board := chessboard.FromString(stringBoard)
	moves := make([]Move, 0, constants.MaxMoves)

	Moves(board, &moves)
	assert.Equal(t, expectedCount, len(moves))
}

func testMovesFromFEN(t *testing.T, expectedCount int, fen string) {
	board := chessboard.FromFen(fen)
	moves := make([]Move, 0, constants.MaxMoves)

	Moves(board, &moves)
	assert.Equal(t, expectedCount, len(moves))
}
