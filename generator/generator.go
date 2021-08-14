package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
)

// attacks of all pieces
func attacks(board chessboard.Board, color chessboard.Color) bitboard.Board {
	return KnightAttacks(board, color) |
		PawnAttacks(board, color) |
		KingAttacks(board, color) |
		RookAttacks(board, color) |
		BishopAttacks(board, color)
}

// isBitmaskUnderAttack checks if certain squares are under attacks from opponent
func isBitmaskUnderAttack(board chessboard.Board, color chessboard.Color, fields bitboard.Board) bool {
	switch {
	case
		RookAttacks(board, color)&fields != 0,
		BishopAttacks(board, color)&fields != 0,
		KnightAttacks(board, color)&fields != 0,
		PawnAttacks(board, color)&fields != 0,
		KingAttacks(board, color)&fields != 0:
		return true
	default:
		return false
	}
}

// generatePseudoLegalMoves without checking legality of king check
func generatePseudoLegalMoves(b chessboard.Board) []Move {
	moves := make([]Move, 0, constants.MovesCacheInitialCapacity)
	KnightMoves(b, &moves)
	PawnMoves(b, &moves)
	KingMoves(b, &moves)
	RookMoves(b, &moves)
	BishopMoves(b, &moves)
	return moves
}

// GenerateLegalMoves ...
func GenerateLegalMoves(b chessboard.Board) []Move {
	moves := generatePseudoLegalMoves(b)
	legalMoves := make([]Move, 0, len(moves))

	for _, m := range moves {
		if isOpponentsKingNotUnderCheck(m.MakeMove(b)) {
			legalMoves = append(legalMoves, m)
		}
	}

	return legalMoves
}

// isOpponentsKingNotUnderCheck for checking legality of the move
func isOpponentsKingNotUnderCheck(board chessboard.Board) bool {
	// check if opponent king is not under check by my pieces
	king := board.Pieces[board.OpponentColor()][chessboard.King]

	if king == bitboard.Empty {
		return false
	}

	kingIndex := king.BitScan()
	pieces := board.Pieces[board.NextMove]
	allPieces := board.BoardOfAllPieces()

	if pieces[chessboard.Pawn]&pawnAttacksCache[board.OpponentColor()][kingIndex] != 0 {
		return false
	}

	if pieces[chessboard.Knight]&knightMovesCache[kingIndex] != 0 {
		return false
	}

	if pieces[chessboard.King]&kingMovesCache[kingIndex] != 0 {
		return false
	}

	rooks := pieces[chessboard.Queen] | pieces[chessboard.Rook]

	if rookMoveRankAttacks[kingIndex][(allPieces&rookMoveRankMask[kingIndex])>>rookMoveRankShift[kingIndex]]&rooks != 0 {
		return false
	}

	if rookMoveFileAttacks[kingIndex][((allPieces&rookMoveFileMask[kingIndex])*rookMoveFileMagic[kingIndex])>>57]&rooks != 0 {
		return false
	}

	bishops := pieces[chessboard.Queen] | pieces[chessboard.Bishop]

	if bishopMoveA8H1Attacks[kingIndex][((allPieces&bishopMoveA8H1Mask[kingIndex])*bishopMoveA8H1Magic[kingIndex])>>57]&bishops != 0 {
		return false
	}

	if bishopMoveA1H8Attacks[kingIndex][((allPieces&bishopMoveA1H8Mask[kingIndex])*bishopMoveA1H8Magic[kingIndex])>>57]&bishops != 0 {
		return false
	}

	return true
}
