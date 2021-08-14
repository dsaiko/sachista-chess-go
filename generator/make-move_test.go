package generator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"

	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/index"
)

func TestMove_MakeMove(t *testing.T) {
	tests := []struct {
		board chessboard.Board
		move  Move
		want  string
	}{
		{
			want: `
  a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`},
		{
			board: *chessboard.FromString(`
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
			move: Move{Piece: chessboard.Pawn, From: index.C7, To: index.D8, PromotionPiece: chessboard.Queen},
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
`},
		{
			board: *chessboard.FromString(`
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
			move: Move{Piece: chessboard.Pawn, From: index.H5, To: index.G6, IsEnPassant: true},
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
`},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			board2 := strings.TrimSpace(tc.move.MakeMove(tc.board).String())
			want := strings.TrimSpace(tc.want)
			if board2 != want {
				t.Errorf("MakeMove() =\n%v, want\n%v", board2, want)
			}
		})
	}
}

func TestZobristFailScenarion1(t *testing.T) {
	board := chessboard.FromFEN("r4rk1/p2pqpb1/bn2pnp1/2pP4/1p2P3/3N1Q1p/PPPBBPPP/RN2K2R w KQ c6 0 3")
	move := Move{Piece: chessboard.King, From: index.E1, To: index.G1, IsEnPassant: false}

	board2 := move.MakeMove(*board)

	assert.Equal(t, board2.ZobristHash, board2.ComputeBoardHash())
}

func TestZobrist(t *testing.T) {
	board := chessboard.FromFEN("r4rk1/p2pqpb1/bn2pnp1/2pP4/1p2P3/3N1Q1p/PPPBBPPP/RN2K2R w KQ c6 0 3")

	for i := 0; i < 1000; i++ {
		moves := GenerateLegalMoves(board)
		board = moves[0].MakeMove(*board)
	}

	assert.Equal(t, board.ZobristHash, board.ComputeBoardHash())
}
