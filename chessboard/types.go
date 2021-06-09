package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/index"
)

type Castling int
type Piece int
type Color int

type Board struct {
	nextMove        Color
	castling        [2]Castling
	pieces          [2][6]bitboard.Board
	halfMoveClock   int
	fullMoveNumber  int
	enPassantTarget index.Index
}
