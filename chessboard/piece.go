package chessboard

type Piece int

// Piece index is used in ChessBoard.Pieces
const (
	King    Piece = 0
	Queen   Piece = 1
	Bishop  Piece = 2
	Knight  Piece = 3
	Rook    Piece = 4
	Pawn    Piece = 5
	NoPiece Piece = -1
)

// String representation of a piece for given color
func (p Piece) String(color Color) string {
	if color == White {
		switch p {
		case King:
			return "K"
		case Queen:
			return "Q"
		case Bishop:
			return "B"
		case Rook:
			return "R"
		case Knight:
			return "N"
		case Pawn:
			return "P"
		}
	} else {
		switch p {
		case King:
			return "k"
		case Queen:
			return "q"
		case Bishop:
			return "b"
		case Rook:
			return "r"
		case Knight:
			return "n"
		case Pawn:
			return "p"
		}
	}
	return "?"
}

func PieceFromNotation(c string) (Piece, Color) {
	switch c {
	case "K":
		return King, White
	case "Q":
		return Queen, White
	case "B":
		return Bishop, White
	case "R":
		return Rook, White
	case "N":
		return Knight, White
	case "P":
		return Pawn, White

	case "k":
		return King, Black
	case "q":
		return Queen, Black
	case "b":
		return Bishop, Black
	case "r":
		return Rook, Black
	case "n":
		return Knight, Black
	case "p":
		return Pawn, Black
	}

	return NoPiece, White
}
