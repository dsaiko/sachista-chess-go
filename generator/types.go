package generator

import (
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/index"
)

type Move struct {
	Piece chessboard.Piece
	From  index.Index
	To    index.Index

	IsEnPassant    bool
	PromotionPiece chessboard.Piece
}
