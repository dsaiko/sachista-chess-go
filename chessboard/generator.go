package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

// attacks of all pieces
func attacks(board *Board, color Color) bitboard.Board {
	return knightAttacks(board, color) |
		pawnAttacks(board, color) |
		kingAttacks(board, color) |
		rookAttacks(board, color) |
		bishopAttacks(board, color)
}

// isBitmaskUnderAttack checks if certain squares are under attacks from opponent
func isBitmaskUnderAttack(board *Board, color Color, fields bitboard.Board) bool {
	switch {
	case
		rookAttacks(board, color)&fields != 0,
		bishopAttacks(board, color)&fields != 0,
		knightAttacks(board, color)&fields != 0,
		pawnAttacks(board, color)&fields != 0,
		kingAttacks(board, color)&fields != 0:
		return true
	default:
		return false
	}
}

// generatePseudoLegalMoves without checking legality of king check
func generatePseudoLegalMoves(b *Board, handler MoveHandler) {
	knightMoves(b, handler)
	pawnMoves(b, handler)
	kingMoves(b, handler)
	rookMoves(b, handler)
	bishopMoves(b, handler)
}

// GenerateLegalMoves ...
func GenerateLegalMoves(b *Board) []Move {
	const MovesCacheInitialCapacity = 32
	legalMoves := make([]Move, 0, MovesCacheInitialCapacity)

	generatePseudoLegalMoves(b, func(m Move) {
		if isOpponentsKingNotUnderCheck(m.ApplyTo(*b)) {
			legalMoves = append(legalMoves, m)
		}
	})

	return legalMoves
}

// isOpponentsKingNotUnderCheck for checking legality of the move
func isOpponentsKingNotUnderCheck(board *Board) bool {
	// check if opponent king is not under check by my pieces
	king := board.Pieces[board.OpponentColor()][King]

	if king == bitboard.EmptyBoard {
		return false
	}

	kingIndex := king.BitScan()
	pieces := board.Pieces[board.NextMove]
	allPieces := board.AllPieces()

	if pieces[Pawn]&pawnAttacksCache[board.OpponentColor()][kingIndex] != 0 {
		return false
	}

	if pieces[Knight]&knightMovesCache[kingIndex] != 0 {
		return false
	}

	if pieces[King]&kingMovesCache[kingIndex] != 0 {
		return false
	}

	rooks := pieces[Queen] | pieces[Rook]

	if rookMoveRankAttacks[kingIndex][(allPieces&rookMoveRankMask[kingIndex])>>rookMoveRankShift[kingIndex]]&rooks != 0 {
		return false
	}

	if rookMoveFileAttacks[kingIndex][((allPieces&rookMoveFileMask[kingIndex])*rookMoveFileMagic[kingIndex])>>57]&rooks != 0 {
		return false
	}

	bishops := pieces[Queen] | pieces[Bishop]

	if bishopMoveA8H1Attacks[kingIndex][((allPieces&bishopMoveA8H1Mask[kingIndex])*bishopMoveA8H1Magic[kingIndex])>>57]&bishops != 0 {
		return false
	}

	if bishopMoveA1H8Attacks[kingIndex][((allPieces&bishopMoveA1H8Mask[kingIndex])*bishopMoveA1H8Magic[kingIndex])>>57]&bishops != 0 {
		return false
	}

	return true
}
