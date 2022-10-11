package chessboard

import (
	"bytes"
	"saiko.cz/sachista/bitboard"
)

type MoveHandler func(Move)

type Move struct {
	Piece Piece
	From  bitboard.Index
	To    bitboard.Index

	IsEnPassant    bool
	PromotionPiece Piece
}

// String move notation
func (m Move) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(m.From.String())
	buffer.WriteString(m.To.String())

	if m.PromotionPiece > 0 { // this excludes NoPiece and King
		buffer.WriteString(m.PromotionPiece.String(Black))
	}

	return buffer.String()
}
