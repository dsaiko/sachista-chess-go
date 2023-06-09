package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

var pawnAttacksCache [bitboard.NumberOfColors][bitboard.NumberOfSquares]bitboard.Board

func init() {
	for i := 0; i < bitboard.NumberOfSquares; i++ {
		piece := bitboard.BoardFromIndex(bitboard.Index(i))

		pawnAttacksCache[White][i] = piece.ShiftedOneNorthEast() | piece.ShiftedOneNorthWest()
		pawnAttacksCache[Black][i] = piece.ShiftedOneSouthEast() | piece.ShiftedOneSouthWest()
	}
}

func pawnAttacks(board *Board, color Color) bitboard.Board {
	if color == White {
		return board.Pieces[color][Pawn].ShiftedOneNorthEast() | board.Pieces[color][Pawn].ShiftedOneNorthWest()
	}
	return board.Pieces[color][Pawn].ShiftedOneSouthEast() | board.Pieces[color][Pawn].ShiftedOneSouthWest()
}

func pawnMoves(board *Board, handler MoveHandler) {
	emptyBoard := ^board.AllPieces()

	pawns := board.Pieces[board.NextMove][Pawn]

	var fromIndex, toIndex bitboard.Index

	for pawns != bitboard.EmptyBoard {
		fromIndex, pawns = pawns.BitPop()
		fromBitBoard := bitboard.BoardFromIndex(fromIndex)

		// get possible moves - moves minus my own color
		// one step forward
		var movesBoard bitboard.Board
		switch board.NextMove {
		case White:
			movesBoard = fromBitBoard.ShiftedOneNorth() & emptyBoard
			if fromIndex < bitboard.IndexA3 {
				// double move
				movesBoard |= movesBoard.ShiftedOneNorth() & emptyBoard
			}
		case Black:
			movesBoard = fromBitBoard.ShiftedOneSouth() & emptyBoard
			if fromIndex > bitboard.IndexH6 {
				// double move
				movesBoard |= movesBoard.ShiftedOneSouth() & emptyBoard
			}
		}

		// get attacks, only against opponent pieces
		var attacks bitboard.Board
		switch board.NextMove {
		case White:
			attacks = fromBitBoard.ShiftedOneNorthEast() | fromBitBoard.ShiftedOneNorthWest()
		default:
			attacks = fromBitBoard.ShiftedOneSouthEast() | fromBitBoard.ShiftedOneSouthWest()
		}

		//attacks := pawnAttacksCache[board.NextMove][fromIndex]
		movesBoard |= attacks & board.OpponentPieces()

		// for all moves
		for movesBoard != bitboard.EmptyBoard {
			// get next move
			toIndex, movesBoard = movesBoard.BitPop()

			// promotion?
			if toIndex > bitboard.IndexH7 || toIndex < bitboard.IndexA2 {
				handler(Move{Piece: Pawn, From: fromIndex, To: toIndex, PromotionPiece: Bishop})
				handler(Move{Piece: Pawn, From: fromIndex, To: toIndex, PromotionPiece: Knight})
				handler(Move{Piece: Pawn, From: fromIndex, To: toIndex, PromotionPiece: Queen})
				handler(Move{Piece: Pawn, From: fromIndex, To: toIndex, PromotionPiece: Rook})
			} else {
				// normal move/capture
				handler(Move{Piece: Pawn, From: fromIndex, To: toIndex})
			}
		}

		// check enpassant capture
		if board.EnPassantTarget != 0 {
			movesBoard = attacks & bitboard.BoardFromIndex(board.EnPassantTarget)
			if movesBoard != bitboard.EmptyBoard {
				targetIndex := movesBoard.BitScan()
				handler(Move{Piece: Pawn, From: fromIndex, To: targetIndex, IsEnPassant: true})
			}
		}
	}
}
