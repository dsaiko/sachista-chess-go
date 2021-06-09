package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/index"
)

type Castling int
type Piece int
type Color int

type Board struct {
	NextMove        Color
	Castling        [2]Castling
	Pieces          [2][6]bitboard.Board
	HalfMoveClock   int
	FullMoveNumber  int
	EnPassantTarget index.Index
	ZobristHash     uint64
}
