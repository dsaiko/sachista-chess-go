package chessboard

// Castling represents castling availability as a bitmask (king-side=1, queen-side=2).
type Castling int

const (
	CastlingNone      Castling = 0
	CastlingKingSide  Castling = 1
	CastlingQueenSide Castling = 2
	CastlingBothSides Castling = 3
)
