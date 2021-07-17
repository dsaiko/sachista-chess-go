package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
)

var pawnMoves [constants.NumberOfColors][constants.NumberOfSquares]bitboard.Board
var pawnDoubleMoves [constants.NumberOfColors][constants.NumberOfSquares]bitboard.Board
var pawnAttacks [constants.NumberOfColors][constants.NumberOfSquares]bitboard.Board

func init() {
	//TODO: perf test function instead of memory constants?
	for i := 0; i < constants.NumberOfSquares; i++ {
		piece := bitboard.FromIndex1(index.Index(i))

		pawnMoves[chessboard.White][i] = piece.Shift(0, 1)
		pawnDoubleMoves[chessboard.White][i] = piece.Shift(0, 2)
		pawnAttacks[chessboard.White][i] = piece.Shift(1, 1) | piece.Shift(-1, 1)

		pawnMoves[chessboard.Black][i] = piece.Shift(0, -1)
		pawnDoubleMoves[chessboard.Black][i] = piece.Shift(0, -2)
		pawnAttacks[chessboard.Black][i] = piece.Shift(1, -1) | piece.Shift(-1, -1)
	}
}

func PawnAttacks(board *chessboard.Board, color chessboard.Color) bitboard.Board {
	if color == chessboard.White {
		return board.Pieces[color][chessboard.Pawn].OneNorthEast() | board.Pieces[color][chessboard.Pawn].OneNorthWest()
	} else {
		return board.Pieces[color][chessboard.Pawn].OneSouthEast() | board.Pieces[color][chessboard.Pawn].OneSouthWest()
	}
}

func PawnMoves(board *chessboard.Board, moves *[]Move) {
	emptyBoard := ^board.AllPieces()

	whiteBaseRank := 16
	blackBaseRank := 999
	whitePromotionRank := 55
	blackPromotionRank := 0

	if board.NextMove == chessboard.Black {
		whiteBaseRank = 0
		blackBaseRank = 47
		whitePromotionRank = 999
		blackPromotionRank = 8
	}

	pawns := board.Pieces[board.NextMove][chessboard.Pawn]

	for pawns != bitboard.Empty {
		sourceIndex := pawns.BitPop()

		//get possible moves - moves minus my onw color
		//one step forward
		movesBoard := pawnMoves[board.NextMove][sourceIndex] & emptyBoard

		//if one step forward was successful and we are on base rank, try double move
		if movesBoard != emptyBoard && sourceIndex < whiteBaseRank {
			movesBoard |= movesBoard.OneNorth() & emptyBoard
		} else if movesBoard != emptyBoard && sourceIndex > blackBaseRank {
			movesBoard |= movesBoard.OneSouth() & emptyBoard
		}

		//get attacks, only against opponent pieces
		attacks := pawnAttacks[board.NextMove][sourceIndex]
		movesBoard |= attacks & board.OpponentPieces()

		//for all moves
		for movesBoard != bitboard.Empty {
			//get next move
			targetIndex := movesBoard.BitPop()

			//promotion?
			if targetIndex > whitePromotionRank || targetIndex < blackPromotionRank {
				*moves = append(*moves, Move{Piece: chessboard.Pawn, From: index.Index(sourceIndex), To: index.Index(targetIndex), PromotionPiece: chessboard.Bishop})
				*moves = append(*moves, Move{Piece: chessboard.Pawn, From: index.Index(sourceIndex), To: index.Index(targetIndex), PromotionPiece: chessboard.Knight})
				*moves = append(*moves, Move{Piece: chessboard.Pawn, From: index.Index(sourceIndex), To: index.Index(targetIndex), PromotionPiece: chessboard.Queen})
				*moves = append(*moves, Move{Piece: chessboard.Pawn, From: index.Index(sourceIndex), To: index.Index(targetIndex), PromotionPiece: chessboard.Rook})
			} else {
				//normal move/capture
				*moves = append(*moves, Move{Piece: chessboard.Pawn, From: index.Index(sourceIndex), To: index.Index(targetIndex)})
			}
		}

		//check enpassant capture
		if board.EnPassantTarget != 0 {
			movesBoard = attacks & bitboard.FromIndex1(board.EnPassantTarget)
			if movesBoard != bitboard.Empty {
				targetIndex := movesBoard.BitScan()
				*moves = append(*moves, Move{Piece: chessboard.Pawn, From: index.Index(sourceIndex), To: index.Index(targetIndex), IsEnPassant: true})
			}
		}
	}

}
