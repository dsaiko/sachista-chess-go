package chessboard

import "strings"

type Piece int

// Piece index is used in ChessBoard.Pieces
const (
	King Piece = iota
	Queen
	Bishop
	Knight
	Rook
	Pawn
)

const NoPiece Piece = -1

// String representation of a piece for given color
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

// PieceFromNotation returns piece and color from a notation string like 'p' or 'P'
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
