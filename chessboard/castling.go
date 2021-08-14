package chessboard

type Castling int

const (
	CastlingNone      Castling = 0
	CastlingKingSide  Castling = 1
	CastlingQueenSide Castling = 2
	CastlingBothSides Castling = 3
)
