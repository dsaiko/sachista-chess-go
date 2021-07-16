package generator

import (
	"bytes"
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
)

func (b *Move) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(b.From.String())
	buffer.WriteString(b.To.String())

	if b.PromotionPiece > 0 { //this excludes NoPiece and King
		buffer.WriteString(string(b.PromotionPiece.Notation(chessboard.Black)))
	}

	return buffer.String()
}

//TODO test *Board
func Attacks(board *chessboard.Board, color chessboard.Color) bitboard.Board {
	return KnightAttacks(board, color) |
		PawnAttacks(board, color)
}

//TODO test *Board
func Moves(board *chessboard.Board, moves *[]Move) {
	KnightMoves(board, moves)
	PawnMoves(board, moves)
}
