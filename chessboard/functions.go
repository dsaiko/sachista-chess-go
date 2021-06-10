package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/zobrist"
)

func Empty() *Board {
	return &Board{FullMoveNumber: 1}
}

func StandardBoard() *Board {
	b := Empty()
	b.Pieces[White][Rook] = bitboard.A1 | bitboard.H1
	b.Pieces[White][Knight] = bitboard.B1 | bitboard.G1
	b.Pieces[White][Bishop] = bitboard.C1 | bitboard.F1
	b.Pieces[White][Queen] = bitboard.D1
	b.Pieces[White][King] = bitboard.E1
	b.Pieces[White][Pawn] = bitboard.A2 | bitboard.B2 | bitboard.C2 | bitboard.D2 | bitboard.E2 | bitboard.F2 | bitboard.G2 | bitboard.H2

	b.Pieces[Black][Rook] = bitboard.A8 | bitboard.H8
	b.Pieces[Black][Knight] = bitboard.B8 | bitboard.G8
	b.Pieces[Black][Bishop] = bitboard.C8 | bitboard.F8
	b.Pieces[Black][Queen] = bitboard.D8
	b.Pieces[Black][King] = bitboard.E8
	b.Pieces[Black][Pawn] = bitboard.A7 | bitboard.B7 | bitboard.C7 | bitboard.D7 | bitboard.E7 | bitboard.F7 | bitboard.G7 | bitboard.H7

	b.Castling[White] = CastlingBothSides
	b.Castling[Black] = CastlingBothSides

	b.UpdateZobrist()
	return b
}

func (b *Board) PiecesOfColor(color Color) bitboard.Board {
	return b.Pieces[color][Queen] |
		b.Pieces[color][King] |
		b.Pieces[color][Rook] |
		b.Pieces[color][Bishop] |
		b.Pieces[color][Knight] |
		b.Pieces[color][Pawn]
}

func (b *Board) AllPieces() bitboard.Board {
	return b.PiecesOfColor(White) | b.PiecesOfColor(Black)
}

func (b *Board) RemoveCastling(color Color, castling Castling) {
	b.Castling[color] = b.Castling[color] & ^castling
}

func (p Piece) Notation(color Color) byte {
	if color == White {
		switch p {
		case King:
			return 'K'
		case Queen:
			return 'Q'
		case Bishop:
			return 'B'
		case Rook:
			return 'R'
		case Knight:
			return 'N'
		case Pawn:
			return 'P'
		}
	} else {
		switch p {
		case King:
			return 'k'
		case Queen:
			return 'q'
		case Bishop:
			return 'b'
		case Rook:
			return 'r'
		case Knight:
			return 'n'
		case Pawn:
			return 'p'
		}
	}
	return '?'
}

func PieceFromNotation(c byte) (Piece, Color) {
	switch c {
	case 'K':
		return King, White
	case 'Q':
		return Queen, White
	case 'B':
		return Bishop, White
	case 'R':
		return Rook, White
	case 'N':
		return Knight, White
	case 'P':
		return Pawn, White

	case 'k':
		return King, Black
	case 'q':
		return Queen, Black
	case 'b':
		return Bishop, Black
	case 'r':
		return Rook, Black
	case 'n':
		return Knight, Black
	case 'p':
		return Pawn, Black
	}

	return NoPiece, White
}

func (c Color) String() string {
	switch c {
	case White:
		return "w"
	case Black:
		return "b"
	default:
		return "?"
	}
}

var rndZobrist = zobrist.NewZobrist()

func (b *Board) UpdateZobrist() {
	hash := uint64(0)

	if b.NextMove != White {
		hash ^= rndZobrist.RndSide
	}

	if b.Castling[White] != 0 {
		hash ^= rndZobrist.RndCastling[White][b.Castling[White]]
	}

	if b.Castling[Black] != 0 {
		hash ^= rndZobrist.RndCastling[Black][b.Castling[Black]]
	}

	if b.EnPassantTarget != 0 {
		hash ^= rndZobrist.RndEnPassant[b.EnPassantTarget]
	}

	for color := 0; color < 2; color++ {
		for piece := 0; piece < 6; piece++ {
			pieces := b.Pieces[color][piece]

			for pieces > 0 {
				hash ^= rndZobrist.RndPieces[color][piece][pieces.BitPop()]
			}
		}
	}

	b.ZobristHash = hash
}
