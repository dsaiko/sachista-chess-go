package chessboard

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/bitboard"
)

func TestMove_MakeMove(t *testing.T) {
	tests := []struct {
		board Board
		move  Move
		want  string
	}{
		{
			// White king-side castling (O-O)
			board: BoardFromFEN("4k3/8/8/8/8/8/8/4K2R w K - 0 1"),
			move:  Move{Piece: King, From: bitboard.IndexE1, To: bitboard.IndexG1},
			want: `
  a b c d e f g h
8 - - - - k - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - R K - 1
  a b c d e f g h
`,
		},
		{
			// White queen-side castling (O-O-O)
			board: BoardFromFEN("4k3/8/8/8/8/8/8/R3K3 w Q - 0 1"),
			move:  Move{Piece: King, From: bitboard.IndexE1, To: bitboard.IndexC1},
			want: `
  a b c d e f g h
8 - - - - k - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - K R - - - - 1
  a b c d e f g h
`,
		},
		{
			// Black king-side castling (O-O)
			board: BoardFromFEN("4k2r/8/8/8/8/8/8/4K3 b k - 0 1"),
			move:  Move{Piece: King, From: bitboard.IndexE8, To: bitboard.IndexG8},
			want: `
  a b c d e f g h
8 - - - - - r k - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - K - - - 1
  a b c d e f g h
`,
		},
		{
			// Black queen-side castling (O-O-O)
			board: BoardFromFEN("r3k3/8/8/8/8/8/8/4K3 b q - 0 1"),
			move:  Move{Piece: King, From: bitboard.IndexE8, To: bitboard.IndexC8},
			want: `
  a b c d e f g h
8 - - k r - - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - K - - - 1
  a b c d e f g h
`,
		},
		{
			board: FromString(`
  a b c d e f g h
8 - - - r - - - - 8
7 - - P - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`),
			move: Move{Piece: Pawn, From: bitboard.IndexC7, To: bitboard.IndexD8, PromotionPiece: Queen},
			want: `
  a b c d e f g h
8 - - - Q - - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`,
		},
		{
			board: FromString(`
  a b c d e f g h
8 - - - r - - - - 8
7 - - P - - - - - 7
6 - - - - - - - - 6
5 - - - - - - p P 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`),
			move: Move{Piece: Pawn, From: bitboard.IndexH5, To: bitboard.IndexG6, IsEnPassant: true},
			want: `
  a b c d e f g h
8 - - - r - - - - 8
7 - - P - - - - - 7
6 - - - - - - P - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`,
		},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			result := tc.move.ApplyTo(tc.board)
			board2 := strings.TrimSpace(result.String())
			want := strings.TrimSpace(tc.want)
			if board2 != want {
				t.Errorf("MakeMove() =\n%v, want\n%v", board2, want)
			}
		})
	}
}

func TestApplyTo_RookMoveRevokesCastling(t *testing.T) {
	board := BoardFromFEN("4k3/8/8/8/8/8/8/R3K2R w KQ - 0 1")
	assert.Equal(t, CastlingBothSides, board.Castling[White])

	// Moving h1 rook revokes king-side castling only
	m := Move{Piece: Rook, From: bitboard.IndexH1, To: bitboard.IndexH4}
	result := m.ApplyTo(board)
	assert.Equal(t, CastlingQueenSide, result.Castling[White])

	// Moving a1 rook revokes queen-side castling only
	m = Move{Piece: Rook, From: bitboard.IndexA1, To: bitboard.IndexA4}
	result = m.ApplyTo(board)
	assert.Equal(t, CastlingKingSide, result.Castling[White])
}

func TestApplyTo_KingMoveRevokesCastling(t *testing.T) {
	board := BoardFromFEN("4k3/8/8/8/8/8/8/R3K2R w KQ - 0 1")
	assert.Equal(t, CastlingBothSides, board.Castling[White])

	// Any king move (non-castling) revokes both castling rights
	m := Move{Piece: King, From: bitboard.IndexE1, To: bitboard.IndexF1}
	result := m.ApplyTo(board)
	assert.Equal(t, CastlingNone, result.Castling[White])
}

func TestApplyTo_CaptureRookRevokesCastling(t *testing.T) {
	// White queen captures Black rook on h8, revoking Black's king-side castling right
	board := BoardFromFEN("r3k2r/8/8/8/7Q/8/8/R3K3 w Qkq - 0 1")
	assert.Equal(t, CastlingQueenSide, board.Castling[White])
	assert.Equal(t, CastlingBothSides, board.Castling[Black])

	m := Move{Piece: Queen, From: bitboard.IndexH4, To: bitboard.IndexH8}
	result := m.ApplyTo(board)
	assert.Equal(t, CastlingQueenSide, result.Castling[White])
	assert.Equal(t, CastlingQueenSide, result.Castling[Black]) // king-side lost
}

func TestApplyTo_EnPassantTargetSet(t *testing.T) {
	board := BoardFromFEN("4k3/8/8/8/8/8/4P3/4K3 w - - 0 1")
	assert.Equal(t, bitboard.Index(0), board.EnPassantTarget)

	// Double pawn push sets en passant target to the skipped square
	m := Move{Piece: Pawn, From: bitboard.IndexE2, To: bitboard.IndexE4}
	result := m.ApplyTo(board)
	assert.Equal(t, bitboard.IndexE3, result.EnPassantTarget)

	// Single pawn push does not set en passant target
	m = Move{Piece: Pawn, From: bitboard.IndexE2, To: bitboard.IndexE3}
	result = m.ApplyTo(board)
	assert.Equal(t, bitboard.Index(0), result.EnPassantTarget)
}

func TestApplyTo_FullMoveCounter(t *testing.T) {
	board := StandardBoard()
	assert.Equal(t, 1, board.FullMoveNumber)

	// After White's move the full-move counter stays at 1
	moves := GenerateLegalMoves(&board)
	board = moves[0].ApplyTo(board)
	assert.Equal(t, 1, board.FullMoveNumber)

	// After Black's move the full-move counter increments to 2
	moves = GenerateLegalMoves(&board)
	board = moves[0].ApplyTo(board)
	assert.Equal(t, 2, board.FullMoveNumber)
}

func TestZobristFailScenarion1(t *testing.T) {
	board := BoardFromFEN("r4rk1/p2pqpb1/bn2pnp1/2pP4/1p2P3/3N1Q1p/PPPBBPPP/RN2K2R w KQ c6 0 3")
	move := Move{Piece: King, From: bitboard.IndexE1, To: bitboard.IndexG1, IsEnPassant: false}

	board2 := move.ApplyTo(board)

	assert.Equal(t, board2.ZobristHash, board2.Hash())
}

func TestZobrist(t *testing.T) {
	board := BoardFromFEN("r4rk1/p2pqpb1/bn2pnp1/2pP4/1p2P3/3N1Q1p/PPPBBPPP/RN2K2R w KQ c6 0 3")

	for range 1000 {
		moves := GenerateLegalMoves(&board)
		board = moves[0].ApplyTo(board)
	}

	assert.Equal(t, board.ZobristHash, board.Hash())
}
