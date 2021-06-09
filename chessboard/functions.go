package chessboard

import "saiko.cz/sachista/bitboard"

func Empty() *Board {
	return &Board{fullMoveNumber: 1}
}

func StandardBoard() *Board {
	b := Empty()
	b.pieces[White][Rook] = bitboard.A1 | bitboard.H1
	b.pieces[White][Knight] = bitboard.B1 | bitboard.G1
	b.pieces[White][Bishop] = bitboard.C1 | bitboard.F1
	b.pieces[White][Queen] = bitboard.D1
	b.pieces[White][King] = bitboard.E1
	b.pieces[White][Pawn] = bitboard.A2 | bitboard.B2 | bitboard.C2 | bitboard.D2 | bitboard.E2 | bitboard.F2 | bitboard.G2 | bitboard.H2

	b.pieces[Black][Rook] = bitboard.A8 | bitboard.H8
	b.pieces[Black][Knight] = bitboard.B8 | bitboard.G8
	b.pieces[Black][Bishop] = bitboard.C8 | bitboard.F8
	b.pieces[Black][Queen] = bitboard.D8
	b.pieces[Black][King] = bitboard.E8
	b.pieces[Black][Pawn] = bitboard.A7 | bitboard.B7 | bitboard.C7 | bitboard.D7 | bitboard.E7 | bitboard.F7 | bitboard.G7 | bitboard.H7

	//TODO test fen of standard board
	//TODO update zobrist of standard board
	return b
}

func (b *Board) Pieces(color Color) bitboard.Board {
	return b.pieces[color][Queen] |
		b.pieces[color][King] |
		b.pieces[color][Rook] |
		b.pieces[color][Bishop] |
		b.pieces[color][Knight] |
		b.pieces[color][Pawn]
}

func (b *Board) AllPieces() bitboard.Board {
	return b.Pieces(White) | b.Pieces(Black)
}

func (b *Board) RemoveCastling(color Color, castling Castling) {
	b.castling[color] = b.castling[color] & ^castling
}
