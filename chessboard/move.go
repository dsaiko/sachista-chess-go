package chessboard

import (
	"bytes"

	"saiko.cz/sachista/bitboard"
)

// MoveHandler is a callback function invoked for each generated move during move generation.
type MoveHandler func(Move)

// Move represents a chess move with source/target squares, piece type,
// and optional promotion or en passant flags.
type Move struct {
	Piece Piece
	From  bitboard.Index
	To    bitboard.Index

	IsEnPassant    bool
	PromotionPiece Piece
}

// String returns the move in long algebraic notation (e.g. "e2e4", "a7a8q" for promotion).
func (m *Move) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(m.From.String())
	buffer.WriteString(m.To.String())

	if m.PromotionPiece > 0 { // excludes NoPiece and King
		buffer.WriteString(m.PromotionPiece.String(Black))
	}

	return buffer.String()
}
