package generator

import (
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
