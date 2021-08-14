package generator

import (
	"bytes"
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

// String move notation
func (m *Move) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(m.From.String())
	buffer.WriteString(m.To.String())

	if m.PromotionPiece > 0 { //this excludes NoPiece and King
		buffer.WriteString(m.PromotionPiece.String(chessboard.Black))
	}

	return buffer.String()
}
