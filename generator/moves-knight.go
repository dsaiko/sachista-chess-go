package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
)

var knightMoveCache [constants.NumberOfSquares]bitboard.Board

func init() {
	for i := 0; i < constants.NumberOfSquares; i++ {
		piece := bitboard.FromIndex1(index.Index(i))
		knightMoveCache[i] =
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

func KnightAttacks(board *chessboard.Board, color chessboard.Color) bitboard.Board {
	pieces := board.Pieces[color][chessboard.Knight]
	attacks := bitboard.Empty

	for pieces != bitboard.Empty {
		attacks |= knightMoveCache[pieces.BitPop()]
	}

	return attacks
}

//TODO test with pointer to Board
func KnightMoves(board *chessboard.Board, moves *[]Move) {
	pieces := board.Pieces[board.NextMove][chessboard.Knight]

	for pieces != bitboard.Empty {
		sourceIndex := pieces.BitPop()
		target := knightMoveCache[sourceIndex] & board.BoardAvailable()

		for target != bitboard.Empty {
			targetIndex := target.BitPop()
			//TODO: optimize?
			*moves = append(*moves, Move{Piece: chessboard.Knight, From: index.Index(sourceIndex), To: index.Index(targetIndex)})
		}
	}

}
