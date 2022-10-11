package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/zobrist"
)

var ZobristKeys = zobrist.NewKeys()

//goland:noinspection SpellCheckingInspection
const StandardBoardFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type Board struct {
	NextMove        Color
	Castling        [bitboard.NumberOfColors]Castling
	Pieces          [bitboard.NumberOfColors][bitboard.NumberOfPieces]bitboard.Board
	HalfMoveClock   int
	FullMoveNumber  int
	EnPassantTarget bitboard.Index
	ZobristHash     uint64
}

type Color int

const (
	White Color = iota
	Black
)

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

// PiecesByColor bitboard of all pieces of one color
func (b *Board) PiecesByColor(color Color) bitboard.Board {
	return b.Pieces[color][Queen] |
		b.Pieces[color][King] |
		b.Pieces[color][Rook] |
		b.Pieces[color][Bishop] |
		b.Pieces[color][Knight] |
		b.Pieces[color][Pawn]
}

// AllPieces bitboard of all pieces
func (b *Board) AllPieces() bitboard.Board {
	return b.PiecesByColor(White) | b.PiecesByColor(Black)
}

// OpponentColor ...
func (b *Board) OpponentColor() Color {
	if b.NextMove == White {
		return Black
	}
	return White
}

// MyPieces ...
func (b *Board) MyPieces() bitboard.Board {
	return b.PiecesByColor(b.NextMove)
}

// OpponentPieces ...
func (b *Board) OpponentPieces() bitboard.Board {
	return b.PiecesByColor(b.OpponentColor())
}

// BoardAvailableToAttack ...
func (b *Board) BoardAvailableToAttack() bitboard.Board {
	return ^b.MyPieces()
}

// MyKingIndex ...
func (b *Board) MyKingIndex() bitboard.Index {
	return b.Pieces[b.NextMove][King].BitScan()
}

// OpponentKingIndex ...
func (b *Board) OpponentKingIndex() bitboard.Index {
	return b.Pieces[b.OpponentColor()][King].BitScan()
}

// RemovedCastling option from a side
func (b *Board) RemovedCastling(color Color, castling Castling) {
	b.Castling[color] &= ^castling
}

// Hash ...
func (b *Board) Hash() uint64 {
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

	var i bitboard.Index

	for color := 0; color < 2; color++ {
		for piece := 0; piece < 6; piece++ {
			pieces := b.Pieces[color][piece]
			for pieces > 0 {
				i, pieces = pieces.BitPop()
				hash ^= ZobristKeys.Pieces[color][piece][i]
			}
		}
	}

	return hash
}

// EmptyBoard chessboard
func EmptyBoard() Board {
	return Board{FullMoveNumber: 1}
}

// StandardBoard chessboard layout
func StandardBoard() Board {
	b := EmptyBoard()
	b.Pieces[White][Rook] = bitboard.BoardA1 | bitboard.BoardH1
	b.Pieces[White][Knight] = bitboard.BoardB1 | bitboard.BoardG1
	b.Pieces[White][Bishop] = bitboard.BoardC1 | bitboard.BoardF1
	b.Pieces[White][Queen] = bitboard.BoardD1
	b.Pieces[White][King] = bitboard.BoardE1
	b.Pieces[White][Pawn] = bitboard.BoardA2 | bitboard.BoardB2 | bitboard.BoardC2 | bitboard.BoardD2 | bitboard.BoardE2 | bitboard.BoardF2 | bitboard.BoardG2 | bitboard.BoardH2

	b.Pieces[Black][Rook] = bitboard.BoardA8 | bitboard.BoardH8
	b.Pieces[Black][Knight] = bitboard.BoardB8 | bitboard.BoardG8
	b.Pieces[Black][Bishop] = bitboard.BoardC8 | bitboard.BoardF8
	b.Pieces[Black][Queen] = bitboard.BoardD8
	b.Pieces[Black][King] = bitboard.BoardE8
	b.Pieces[Black][Pawn] = bitboard.BoardA7 | bitboard.BoardB7 | bitboard.BoardC7 | bitboard.BoardD7 | bitboard.BoardE7 | bitboard.BoardF7 | bitboard.BoardG7 | bitboard.BoardH7

	b.Castling[White] = CastlingBothSides
	b.Castling[Black] = CastlingBothSides

	b.ZobristHash = b.Hash()
	return b
}
