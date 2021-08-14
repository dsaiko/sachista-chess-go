package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
)

//goland:noinspection SpellCheckingInspection
const StandardBoardFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Castling int

const (
	CastlingNone      Castling = 0
	CastlingKingSide  Castling = 1
	CastlingQueenSide Castling = 2
	CastlingBothSides Castling = 3
)

type Color int

const (
	White Color = 0
	Black Color = 1
)

type Board struct {
	NextMove        Color
	Castling        [constants.NumberOfColors]Castling
	Pieces          [constants.NumberOfColors][constants.NumberOfPieces]bitboard.Board
	HalfMoveClock   int
	FullMoveNumber  int
	EnPassantTarget index.Index
	ZobristHash     uint64
}
