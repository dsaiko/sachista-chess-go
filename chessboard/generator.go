package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

// attacks returns a combined attack bitboard for all pieces of the given color.
func attacks(board *Board, color Color) bitboard.Board {
	return knightAttacks(board, color) |
		pawnAttacks(board, color) |
		kingAttacks(board, color) |
		rookAttacks(board, color) |
		bishopAttacks(board, color)
}

// isBitmaskUnderAttack reports whether any of the given squares are attacked by the specified color.
// Uses short-circuit evaluation: returns as soon as any piece type attacks the squares.
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

// generatePseudoLegalMoves generates all pseudo-legal moves (without filtering for king safety).
func generatePseudoLegalMoves(b *Board, handler MoveHandler) {
	knightMoves(b, handler)
	pawnMoves(b, handler)
	kingMoves(b, handler)
	rookMoves(b, handler)
	bishopMoves(b, handler)
}

// GenerateLegalMoves returns all legal moves for the current position by generating
// pseudo-legal moves and filtering out those that leave the king in check.
func GenerateLegalMoves(b *Board) []Move {
	const MovesCacheInitialCapacity = 32
	legalMoves := make([]Move, 0, MovesCacheInitialCapacity)

	generatePseudoLegalMoves(b, func(m Move) {
		nb := m.ApplyTo(*b)
		if isOpponentsKingNotUnderCheck(&nb) {
			legalMoves = append(legalMoves, m)
		}
	})

	return legalMoves
}

// isOpponentsKingNotUnderCheck reports whether the opponent's king (the side that just moved)
// is not in check. Used to verify move legality after applying a move.
func isOpponentsKingNotUnderCheck(board *Board) bool {
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

	if rookMoveFileAttacks[kingIndex][((allPieces&rookMoveFileMask[kingIndex])*rookMoveFileMagic[kingIndex])>>57]&rooks != 0 { //nolint:revive
		return false
	}

	bishops := pieces[Queen] | pieces[Bishop]

	if bishopMoveA8H1Attacks[kingIndex][((allPieces&bishopMoveA8H1Mask[kingIndex])*bishopMoveA8H1Magic[kingIndex])>>57]&bishops != 0 { //nolint:revive
		return false
	}

	if bishopMoveA1H8Attacks[kingIndex][((allPieces&bishopMoveA1H8Mask[kingIndex])*bishopMoveA1H8Magic[kingIndex])>>57]&bishops != 0 { //nolint:revive
		return false
	}

	return true
}
