package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove_String(t *testing.T) {
	move := Move{Piece: Pawn, From: bitboard.IndexA2, To: bitboard.IndexA3}
	assert.Equal(t, "a2a3", move.String())

	move = Move{Piece: Pawn, From: bitboard.IndexA7, To: bitboard.IndexB8, PromotionPiece: Queen}
	assert.Equal(t, "a7b8q", move.String())
}

func testMovesFromString(t *testing.T, expectedCount int, stringBoard string) {
	board := FromString(stringBoard)
	size := 0
	generatePseudoLegalMoves(&board, func(m Move) {
		size++
	})
	assert.Equal(t, expectedCount, size)
}

func testMovesFromFEN(t *testing.T, expectedCount int, fen string) {
	board := BoardFromFEN(fen)
	size := 0
	generatePseudoLegalMoves(&board, func(m Move) {
		size++
	})
	assert.Equal(t, expectedCount, size)
}

func Test_isOpponentsKingNotUnderCheck(t *testing.T) {

	tests := []struct {
		name  string
		board Board
		want  bool
	}{
		{
			name: "No check",
			board: FromString(`
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
			board: FromString(`
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
			board: FromString(`
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
			board: FromString(`
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
			board: FromString(`
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
			board: FromString(`
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
			board: FromString(`
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
			board: FromString(`
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
			board: FromString(`
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
			board: FromString(`
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
			if got := isOpponentsKingNotUnderCheck(&tt.board); got != tt.want {
				t.Errorf("isOpponentsKingNotUnderCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
