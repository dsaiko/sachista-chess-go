package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/zobrist"
)

// ZobristKeys holds the pre-computed random keys used for Zobrist hashing of board positions.
var ZobristKeys = zobrist.NewKeys()

// StandardBoardFEN is the FEN string for the standard chess starting position.
//
//goland:noinspection SpellCheckingInspection
const StandardBoardFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// Board represents a complete chess position, including piece placement,
// side to move, castling rights, en passant target, and move counters.
type Board struct {
	Pieces          [bitboard.NumberOfColors][bitboard.NumberOfPieces]bitboard.Board
	Occupied        [bitboard.NumberOfColors]bitboard.Board // cached PiecesByColor
	ZobristHash     uint64
	NextMove        Color
	Castling        [bitboard.NumberOfColors]Castling
	EnPassantTarget bitboard.Index
	HalfMoveClock   int
	FullMoveNumber  int
}

// Color represents a player's side (White or Black).
type Color int

const (
	White Color = iota
	Black
)

// String returns "w" for White, "b" for Black.
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

// PiecesByColor returns the bitboard of all pieces of the given color.
func (b *Board) PiecesByColor(color Color) bitboard.Board {
	return b.Occupied[color]
}

// AllPieces returns the bitboard of all pieces on the board regardless of color.
func (b *Board) AllPieces() bitboard.Board {
	return b.Occupied[White] | b.Occupied[Black]
}

// RecomputeOccupied recomputes the cached Occupied bitboards from Pieces.
func (b *Board) RecomputeOccupied() {
	b.Occupied[White] = b.Pieces[White][Queen] |
		b.Pieces[White][King] |
		b.Pieces[White][Rook] |
		b.Pieces[White][Bishop] |
		b.Pieces[White][Knight] |
		b.Pieces[White][Pawn]
	b.Occupied[Black] = b.Pieces[Black][Queen] |
		b.Pieces[Black][King] |
		b.Pieces[Black][Rook] |
		b.Pieces[Black][Bishop] |
		b.Pieces[Black][Knight] |
		b.Pieces[Black][Pawn]
}

// OpponentColor returns the opposite color using branchless XOR.
func (b *Board) OpponentColor() Color {
	return 1 ^ b.NextMove
}

// MyPieces returns the bitboard of all pieces belonging to the side to move.
func (b *Board) MyPieces() bitboard.Board {
	return b.PiecesByColor(b.NextMove)
}

// OpponentPieces returns the bitboard of all pieces belonging to the opponent.
func (b *Board) OpponentPieces() bitboard.Board {
	return b.PiecesByColor(b.OpponentColor())
}

// BoardAvailableToAttack returns a bitmask of squares the current side may move to
// (all squares except those occupied by own pieces).
func (b *Board) BoardAvailableToAttack() bitboard.Board {
	return ^b.MyPieces()
}

// MyKingIndex returns the square index of the current side's king.
func (b *Board) MyKingIndex() bitboard.Index {
	return b.Pieces[b.NextMove][King].BitScan()
}

// OpponentKingIndex returns the square index of the opponent's king.
func (b *Board) OpponentKingIndex() bitboard.Index {
	return b.Pieces[b.OpponentColor()][King].BitScan()
}

// RemovedCastling clears the specified castling right for the given color.
func (b *Board) RemovedCastling(color Color, castling Castling) {
	b.Castling[color] &= ^castling
}

// Hash computes the full Zobrist hash of the current board position from scratch.
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

	for color := range 2 {
		for piece := range 6 {
			pieces := b.Pieces[color][piece]
			for pieces > 0 {
				i, pieces = pieces.BitPop()
				hash ^= ZobristKeys.Pieces[color][piece][i]
			}
		}
	}

	return hash
}

// EmptyBoard returns a board with no pieces and the full move counter set to 1.
func EmptyBoard() Board {
	return Board{FullMoveNumber: 1}
}

// StandardBoard returns a board set up in the standard chess starting position.
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

	b.RecomputeOccupied()
	b.ZobristHash = b.Hash()
	return b
}
