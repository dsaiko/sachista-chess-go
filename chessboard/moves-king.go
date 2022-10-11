package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

const (
	WhiteCastleOOEmpty    = bitboard.BoardF1 | bitboard.BoardG1
	WhiteCastleOOAttacks  = bitboard.BoardE1 | bitboard.BoardF1 | bitboard.BoardG1
	WhiteCastleOOOEmpty   = bitboard.BoardB1 | bitboard.BoardC1 | bitboard.BoardD1
	WhiteCastleOOOAttacks = bitboard.BoardC1 | bitboard.BoardD1 | bitboard.BoardE1

	BlackCastleOOEmpty    = bitboard.BoardF8 | bitboard.BoardG8
	BlackCastleOOAttacks  = bitboard.BoardE8 | bitboard.BoardF8 | bitboard.BoardG8
	BlackCastleOOOEmpty   = bitboard.BoardB8 | bitboard.BoardC8 | bitboard.BoardD8
	BlackCastleOOOAttacks = bitboard.BoardC8 | bitboard.BoardD8 | bitboard.BoardE8
)

var kingMovesCache [bitboard.NumberOfSquares]bitboard.Board

func init() {
	for i := 0; i < bitboard.NumberOfSquares; i++ {
		piece := bitboard.BoardFromIndex(bitboard.Index(i))

		kingMovesCache[i] =
			piece.Shifted(1, -1) |
				piece.Shifted(1, 0) |
				piece.Shifted(1, 1) |
				piece.Shifted(0, -1) |
				piece.Shifted(0, 1) |
				piece.Shifted(-1, -1) |
				piece.Shifted(-1, 0) |
				piece.Shifted(-1, 1)
	}
}

func kingAttacks(board *Board, color Color) bitboard.Board {
	king := board.Pieces[color][King]
	if king == bitboard.EmptyBoard {
		return bitboard.EmptyBoard
	}

	return kingMovesCache[king.BitScan()]
}

func kingMoves(board *Board, handler MoveHandler) {
	king := board.Pieces[board.NextMove][King]
	if king == bitboard.EmptyBoard {
		return
	}

	kingIndex := king.BitScan()
	movesBoard := kingMovesCache[kingIndex] & board.BoardAvailableToAttack()

	var toIndex bitboard.Index
	// for all moves
	for movesBoard != bitboard.EmptyBoard {
		toIndex, movesBoard = movesBoard.BitPop()
		handler(Move{Piece: King, From: kingIndex, To: toIndex})
	}

	// check castling options
	if board.Castling[board.NextMove] == CastlingNone {
		return
	}

	allPieces := board.AllPieces()

	if board.NextMove == White {
		// if castling available
		if (board.Castling[White]&CastlingKingSide != 0) && (allPieces&WhiteCastleOOEmpty == 0) {
			if !isBitmaskUnderAttack(board, Black, WhiteCastleOOAttacks) {
				// add short castling move
				handler(Move{Piece: King, From: kingIndex, To: bitboard.IndexG1})
			}
		}
		if (board.Castling[White]&CastlingQueenSide != 0) && (allPieces&WhiteCastleOOOEmpty == 0) {
			if !isBitmaskUnderAttack(board, Black, WhiteCastleOOOAttacks) {
				// add long castling move
				handler(Move{Piece: King, From: kingIndex, To: bitboard.IndexC1})
			}
		}
	} else {
		// if castling available
		if (board.Castling[Black]&CastlingKingSide != 0) && (allPieces&BlackCastleOOEmpty == 0) {
			if !isBitmaskUnderAttack(board, White, BlackCastleOOAttacks) {
				// add short castling move
				handler(Move{Piece: King, From: kingIndex, To: bitboard.IndexG8})
			}
		}
		if (board.Castling[Black]&CastlingQueenSide != 0) && (allPieces&BlackCastleOOOEmpty == 0) {
			if !isBitmaskUnderAttack(board, White, BlackCastleOOOAttacks) {
				// add long castling move
				handler(Move{Piece: King, From: kingIndex, To: bitboard.IndexC8})
			}
		}
	}
}
