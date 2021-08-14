package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/chessboard"

	"saiko.cz/sachista/index"
)

func TestMove_String(t *testing.T) {
	move := Move{Piece: chessboard.Pawn, From: index.A2, To: index.A3}
	assert.Equal(t, "a2a3", move.String())

	move = Move{Piece: chessboard.Pawn, From: index.A7, To: index.B8, PromotionPiece: chessboard.Queen}
	assert.Equal(t, "a7b8q", move.String())
}

func testMovesFromString(t *testing.T, expectedCount int, stringBoard string) {
	board := chessboard.FromString(stringBoard)
	moves := GeneratePseudoLegalMoves(board)
	assert.Equal(t, expectedCount, len(moves))
}

func testMovesFromFEN(t *testing.T, expectedCount int, fen string) {
	board := chessboard.FromFen(fen)
	moves := GeneratePseudoLegalMoves(board)
	assert.Equal(t, expectedCount, len(moves))
}

func Test_isOpponentsKingNotUnderCheck(t *testing.T) {

	tests := []struct {
		name  string
		board *chessboard.Board
		want  bool
	}{
		{
			name: "No check",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
			`),
			want: true,
		},
		{
			name: "Check by King",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - K - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - - - - R 1
  a b c d e f g h
			`),
			want: false,
		},
		{
			name: "Check by Queen 1",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - Q - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
			`),
			want: false,
		},
		{
			name: "Check by Queen 2",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - Q - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
			`),
			want: false,
		},
		{
			name: "No check by Bishop",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - r - - 7
6 - - - - B - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
			`),
			want: true,
		},
		{
			name: "Check by Bishop",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - - - - 7
6 - - - - B - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
			`),
			want: false,
		},
		{
			name: "Check by Rook",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - R - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
			`),
			want: false,
		},
		{
			name: "No check by Rook",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - - r - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - R - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
			`),
			want: true,
		},
		{
			name: "Check by Knight",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - - - - 7
6 - - - - - N - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
			`),
			want: false,
		},
		{
			name: "Check by Pawn",
			board: chessboard.FromString(`
  a b c d e f g h
8 - - - - - - k - 8
7 - - - - - P - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
			`),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOpponentsKingNotUnderCheck(tt.board); got != tt.want {
				t.Errorf("isOpponentsKingNotUnderCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
