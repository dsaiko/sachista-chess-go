package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
)

var knightMovesCache [constants.NumberOfSquares]bitboard.Board

func init() {
	for i := 0; i < constants.NumberOfSquares; i++ {
		piece := bitboard.FromIndex1(index.Index(i))
		knightMovesCache[i] =
			piece.Shift(2, 1) |
				piece.Shift(2, -1) |
				piece.Shift(1, 2) |
				piece.Shift(-1, 2) |
				piece.Shift(-2, 1) |
				piece.Shift(-2, -1) |
				piece.Shift(-1, -2) |
				piece.Shift(1, -2)
	}
}

func knightAttacks(board chessboard.Board, color chessboard.Color) bitboard.Board {
	pieces := board.Pieces[color][chessboard.Knight]
	attacks := bitboard.Empty

	for pieces != bitboard.Empty {
		attacks |= knightMovesCache[pieces.BitPop()]
	}

	return attacks
}

func KnightMoves(board chessboard.Board, moves *[]Move) {
	pieces := board.Pieces[board.NextMove][chessboard.Knight]

	for pieces != bitboard.Empty {
		sourceIndex := pieces.BitPop()
		target := knightMovesCache[sourceIndex] & board.BoardAvailableToAttack()

		for target != bitboard.Empty {
			targetIndex := target.BitPop()
			*moves = append(*moves, Move{Piece: chessboard.Knight, From: index.Index(sourceIndex), To: index.Index(targetIndex)})
		}
	}
}
