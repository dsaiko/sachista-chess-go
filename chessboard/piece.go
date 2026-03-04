package chessboard

import "strings"

// Piece represents a chess piece type, used as an index into Board.Pieces.
type Piece int

// Piece type constants, ordered to match their index in the Board.Pieces array.
const (
	King Piece = iota
	Queen
	Bishop
	Knight
	Rook
	Pawn
)

// NoPiece is a sentinel value indicating the absence of a piece.
const NoPiece Piece = -1

// String returns the algebraic notation letter for the piece (uppercase for White, lowercase for Black).
func (p Piece) String(color Color) string {
	c := "?"

	switch p {
	case King:
		c = "K"
	case Queen:
		c = "Q"
	case Bishop:
		c = "B"
	case Rook:
		c = "R"
	case Knight:
		c = "N"
	case Pawn:
		c = "P"
	}

	if color == Black {
		c = strings.ToLower(c)
	}

	return c
}

// PieceFromNotation returns the piece type and color from a single-character notation string
// (e.g. "P" -> Pawn, White; "p" -> Pawn, Black).
func PieceFromNotation(c string) (Piece, Color) {
	color := White

	if strings.ToLower(c) == c {
		color = Black
	}

	c = strings.ToLower(c)

	switch c {
	case "k":
		return King, color
	case "q":
		return Queen, color
	case "b":
		return Bishop, color
	case "r":
		return Rook, color
	case "n":
		return Knight, color
	case "p":
		return Pawn, color
	}

	return NoPiece, color
}
