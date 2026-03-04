package chessboard

// Castling represents castling availability as a bitmask (king-side=1, queen-side=2).
type Castling int

const (
	CastlingNone      Castling = 0 // no castling available
	CastlingKingSide  Castling = 1 // king-side (short, O-O) castling available
	CastlingQueenSide Castling = 2 // queen-side (long, O-O-O) castling available
	CastlingBothSides Castling = 3 // both castling options available
)
