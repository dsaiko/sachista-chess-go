package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
	"saiko.cz/sachista/zobrist"
)

var ZobristKeys = zobrist.NewKeys()

//goland:noinspection SpellCheckingInspection
const StandardBoardFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Board struct {
	NextMove        Color
	Castling        [constants.NumberOfColors]Castling
	Pieces          [constants.NumberOfColors][constants.NumberOfPieces]bitboard.Board
	HalfMoveClock   int
	FullMoveNumber  int
	EnPassantTarget index.Index
	ZobristHash     uint64
}

// Empty chessboard
func Empty() Board {
	return Board{FullMoveNumber: 1}
}

// Standard chessboard layout
func Standard() Board {
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

	b.ZobristHash = b.ComputeBoardHash()
	return b
}

// BoardOfColor bitboard of all pieces of one color
func (b Board) BoardOfColor(color Color) bitboard.Board {
	return b.Pieces[color][Queen] |
		b.Pieces[color][King] |
		b.Pieces[color][Rook] |
		b.Pieces[color][Bishop] |
		b.Pieces[color][Knight] |
		b.Pieces[color][Pawn]
}

// BoardOfAllPieces bitboard of all pieces
func (b Board) BoardOfAllPieces() bitboard.Board {
	return b.BoardOfColor(White) | b.BoardOfColor(Black)
}

// OpponentColor ...
func (b Board) OpponentColor() Color {
	if b.NextMove == White {
		return Black
	}
	return White
}

// BoardOfMyPieces ...
func (b Board) BoardOfMyPieces() bitboard.Board {
	return b.BoardOfColor(b.NextMove)
}

// BoardOfOpponentPieces ...
func (b Board) BoardOfOpponentPieces() bitboard.Board {
	return b.BoardOfColor(b.OpponentColor())
}

// BoardAvailableToAttack ..
func (b Board) BoardAvailableToAttack() bitboard.Board {
	return ^b.BoardOfMyPieces()
}

// MyKingIndex ...
func (b Board) MyKingIndex() index.Index {
	return index.Index(b.Pieces[b.NextMove][King].BitScan())
}

// OpponentKingIndex ...
func (b Board) OpponentKingIndex() index.Index {
	return index.Index(b.Pieces[b.OpponentColor()][King].BitScan())
}

// RemoveCastling option from a side
func (b *Board) RemoveCastling(color Color, castling Castling) {
	b.Castling[color] &= ^castling
}

// ComputeBoardHash hash
func (b Board) ComputeBoardHash() uint64 {
	hash := uint64(0)

	if b.NextMove != White {
		hash ^= ZobristKeys.Side
	}

	if b.Castling[White] != 0 {
		hash ^= ZobristKeys.Castling[White][b.Castling[White]]
	}

	if b.Castling[Black] != 0 {
		hash ^= ZobristKeys.Castling[Black][b.Castling[Black]]
	}

	if b.EnPassantTarget != 0 {
		hash ^= ZobristKeys.EnPassant[b.EnPassantTarget]
	}

	for color := 0; color < 2; color++ {
		for piece := 0; piece < 6; piece++ {
			pieces := b.Pieces[color][piece]

			for pieces > 0 {
				hash ^= ZobristKeys.Pieces[color][piece][pieces.BitPop()]
			}
		}
	}

	return hash
}
