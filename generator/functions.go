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
		PawnAttacks(board, color) |
		KingAttacks(board, color) |
		RookAttacks(board, color) |
		BishopAttacks(board, color)
}

//TODO test *
func IsBitmaskUnderAttack(board *chessboard.Board, color chessboard.Color, fields bitboard.Board) bool {
	attacks := KnightAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	attacks = PawnAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	attacks = KingAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	attacks = RookAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	attacks = BishopAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	return false
}

//TODO test *Board
func Moves(board *chessboard.Board, moves *[]Move) {
	KnightMoves(board, moves)
	PawnMoves(board, moves)
	KingMoves(board, moves)
	RookMoves(board, moves)
	BishopMoves(board, moves)
}
