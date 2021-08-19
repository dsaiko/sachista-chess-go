package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
)

const (
	WhiteCastleOOEmpty    = bitboard.F1 | bitboard.G1
	WhiteCastleOOAttacks  = bitboard.E1 | bitboard.F1 | bitboard.G1
	WhiteCastleOOOEmpty   = bitboard.B1 | bitboard.C1 | bitboard.D1
	WhiteCastleOOOAttacks = bitboard.C1 | bitboard.D1 | bitboard.E1

	BlackCastleOOEmpty    = bitboard.F8 | bitboard.G8
	BlackCastleOOAttacks  = bitboard.E8 | bitboard.F8 | bitboard.G8
	BlackCastleOOOEmpty   = bitboard.B8 | bitboard.C8 | bitboard.D8
	BlackCastleOOOAttacks = bitboard.C8 | bitboard.D8 | bitboard.E8
)

var kingMovesCache [constants.NumberOfSquares]bitboard.Board

func init() {
	for i := 0; i < constants.NumberOfSquares; i++ {
		piece := bitboard.FromIndex1(index.Index(i))

		kingMovesCache[i] =
			piece.Shift(1, -1) |
				piece.Shift(1, 0) |
				piece.Shift(1, 1) |
				piece.Shift(0, -1) |
				piece.Shift(0, 1) |
				piece.Shift(-1, -1) |
				piece.Shift(-1, 0) |
				piece.Shift(-1, 1)
	}
}

func kingAttacks(board chessboard.Board, color chessboard.Color) bitboard.Board {
	king := board.Pieces[color][chessboard.King]
	if king == bitboard.Empty {
		return bitboard.Empty
	}

	return kingMovesCache[king.BitScan()]
}

func kingMoves(board chessboard.Board, handler MoveHandler) {
	king := board.Pieces[board.NextMove][chessboard.King]
	if king == bitboard.Empty {
		return
	}

	kingIndex := index.Index(king.BitScan())
	movesBoard := kingMovesCache[kingIndex] & board.BoardAvailableToAttack()

	// for all moves
	for movesBoard != bitboard.Empty {
		toIndex := movesBoard.BitPop()
		handler(Move{Piece: chessboard.King, From: kingIndex, To: index.Index(toIndex)})
	}

	// check castling options
	if board.Castling[board.NextMove] == chessboard.CastlingNone {
		return
	}

	allPieces := board.BoardOfAllPieces()

	if board.NextMove == chessboard.White {
		// if castling available
		if (board.Castling[chessboard.White]&chessboard.CastlingKingSide != 0) && (allPieces&WhiteCastleOOEmpty == 0) {
			if !isBitmaskUnderAttack(board, chessboard.Black, WhiteCastleOOAttacks) {
				// add short castling move
				handler(Move{Piece: chessboard.King, From: kingIndex, To: index.G1})
			}
		}
		if (board.Castling[chessboard.White]&chessboard.CastlingQueenSide != 0) && (allPieces&WhiteCastleOOOEmpty == 0) {
			if !isBitmaskUnderAttack(board, chessboard.Black, WhiteCastleOOOAttacks) {
				// add long castling move
				handler(Move{Piece: chessboard.King, From: kingIndex, To: index.C1})
			}
		}
	} else {
		// if castling available
		if (board.Castling[chessboard.Black]&chessboard.CastlingKingSide != 0) && (allPieces&BlackCastleOOEmpty == 0) {
			if !isBitmaskUnderAttack(board, chessboard.White, BlackCastleOOAttacks) {
				// add short castling move
				handler(Move{Piece: chessboard.King, From: kingIndex, To: index.G8})
			}
		}
		if (board.Castling[chessboard.Black]&chessboard.CastlingQueenSide != 0) && (allPieces&BlackCastleOOOEmpty == 0) {
			if !isBitmaskUnderAttack(board, chessboard.White, BlackCastleOOOAttacks) {
				// add long castling move
				handler(Move{Piece: chessboard.King, From: kingIndex, To: index.C8})
			}
		}
	}
}
