package generator

import (
	"bytes"
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
)

func (m *Move) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(m.From.String())
	buffer.WriteString(m.To.String())

	if m.PromotionPiece > 0 { //this excludes NoPiece and King
		buffer.WriteString(m.PromotionPiece.String(chessboard.Black))
	}

	return buffer.String()
}

//TODO test *Board
func Attacks(board *chessboard.Board, color chessboard.Color) bitboard.Board {
	return KnightAttacks(board, color) |
		PawnAttacks(board, color) |
		KingAttacks(board, color) |
		RookAttacks(board, color) |
		BishopAttacks(board, color)
}

//TODO test *
func IsBitmaskUnderAttack(board *chessboard.Board, color chessboard.Color, fields bitboard.Board) bool {
	attacks := KnightAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	attacks = PawnAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	attacks = KingAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	attacks = RookAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	attacks = BishopAttacks(board, color)
	if attacks&fields != 0 {
		return true
	}

	return false
}

//TODO reorganize
//TODO test *Board
func GeneratePseudoLegalMoves(b *chessboard.Board) []Move {
	moves := make([]Move, 0, constants.MovesCacheInitialCapacity)

	KnightMoves(b, &moves)
	PawnMoves(b, &moves)
	RookMoves(b, &moves)
	BishopMoves(b, &moves)

	return moves
}

func GenerateLegalMoves(b *chessboard.Board) []Move {
	moves := GeneratePseudoLegalMoves(b)
	legalMoves := make([]Move, 0, len(moves))

	for _, m := range moves {
		nextBoard := m.MakeMove(*b)
		if isOpponentsKingNotUnderCheck(nextBoard) {
			legalMoves = append(legalMoves, m)
		}
	}

	return legalMoves
}

func isOpponentsKingNotUnderCheck(board *chessboard.Board) bool {
	//TODO try to cache oponent color and mycolor

	//check if opponent king is not under check by my pieces
	king := board.Pieces[board.OpponentColor()][chessboard.King]

	if king == bitboard.Empty {
		return false
	}

	kingIndex := king.BitScan()
	pieces := board.Pieces[board.NextMove]

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

	if rookMoveRankAttacks[kingIndex][(board.BoardOfAllPieces()&rookMoveRankMask[kingIndex])>>rookMoveRankShift[kingIndex]]&rooks != 0 {
		return false
	}

	if rookMoveFileAttacks[kingIndex][((board.BoardOfAllPieces()&rookMoveFileMask[kingIndex])*rookMoveFileMagic[kingIndex])>>57]&rooks != 0 {
		return false
	}

	bishops := pieces[chessboard.Queen] | pieces[chessboard.Bishop]

	if bishopMoveA8H1Attacks[kingIndex][((board.BoardOfAllPieces()&bishopMoveA8H1Mask[kingIndex])*bishopMoveA8H1Magic[kingIndex])>>57]&bishops != 0 {
		return false
	}

	if bishopMoveA1H8Attacks[kingIndex][((board.BoardOfAllPieces()&bishopMoveA1H8Mask[kingIndex])*bishopMoveA1H8Magic[kingIndex])>>57]&bishops != 0 {
		return false
	}

	return true
}
