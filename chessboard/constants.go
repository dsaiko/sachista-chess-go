package chessboard

//goland:noinspection SpellCheckingInspection
const StandardBoardFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

const (
	CastlingNone      Castling = 0
	CastlingKingSide  Castling = 1
	CastlingQueenSide Castling = 2
	CastlingBothSides Castling = 3
)

//Piece index is used in ChessBoard.pieces as index
const (
	King    Piece = 0
	Queen   Piece = 1
	Bishop  Piece = 2
	Knight  Piece = 3
	Rook    Piece = 4
	Pawn    Piece = 5
	NoPiece Piece = -1
)

const (
	White Color = 0
	Black Color = 1
)
