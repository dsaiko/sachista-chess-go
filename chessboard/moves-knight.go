package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

var knightMovesCache [bitboard.NumberOfSquares]bitboard.Board

func init() {
	for i := 0; i < bitboard.NumberOfSquares; i++ {
		piece := bitboard.BoardFromIndex(bitboard.Index(i))
		knightMovesCache[i] =
			piece.Shifted(2, 1) |
				piece.Shifted(2, -1) |
				piece.Shifted(1, 2) |
				piece.Shifted(-1, 2) |
				piece.Shifted(-2, 1) |
				piece.Shifted(-2, -1) |
				piece.Shifted(-1, -2) |
				piece.Shifted(1, -2)
	}
}

func knightAttacks(board *Board, color Color) bitboard.Board {
	pieces := board.Pieces[color][Knight]
	attacks := bitboard.EmptyBoard

	var i bitboard.Index
	for pieces != bitboard.EmptyBoard {
		i, pieces = pieces.BitPop()
		attacks |= knightMovesCache[i]
	}

	return attacks
}

func knightMoves(board *Board, handler MoveHandler) {
	pieces := board.Pieces[board.NextMove][Knight]

	var fromIndex, toIndex bitboard.Index

	for pieces != bitboard.EmptyBoard {
		fromIndex, pieces = pieces.BitPop()
		target := knightMovesCache[fromIndex] & board.BoardAvailableToAttack()

		for target != bitboard.EmptyBoard {
			toIndex, target = target.BitPop()
			handler(Move{Piece: Knight, From: fromIndex, To: toIndex})
		}
	}
}
