package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
)

type Castling int
type Piece int
type Color int

type Board struct {
	NextMove        Color
	Castling        [constants.NumberOfColors]Castling
	Pieces          [constants.NumberOfColors][constants.NumberOfPieces]bitboard.Board
	HalfMoveClock   int
	FullMoveNumber  int
	EnPassantTarget index.Index
	ZobristHash     uint64
}
