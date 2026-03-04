package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/bitboard"
)

func TestMove_String(t *testing.T) {
	move := Move{Piece: Pawn, From: bitboard.IndexA2, To: bitboard.IndexA3}
	assert.Equal(t, "a2a3", move.String())

	move = Move{Piece: Pawn, From: bitboard.IndexA7, To: bitboard.IndexB8, PromotionPiece: Queen}
	assert.Equal(t, "a7b8q", move.String())
}

func testMovesFromString(t *testing.T, expectedCount int, stringBoard string) {
	size := 0
	generatePseudoLegalMoves(new(FromString(stringBoard)), func(_ Move) {
		size++
	})
	assert.Equal(t, expectedCount, size)
}

func testMovesFromFEN(t *testing.T, expectedCount int, fen string) {
	size := 0
	generatePseudoLegalMoves(new(BoardFromFEN(fen)), func(_ Move) {
		size++
	})
	assert.Equal(t, expectedCount, size)
}

func TestGenerateLegalMoves(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  int
	}{
		{
			// Standard start: 16 pawn + 4 knight = 20
			name:  "Standard starting position",
			board: StandardBoard(),
			want:  20,
		},
		{
			// Kiwipete: well-known position with 48 legal moves at depth 1
			name:  "Kiwipete",
			board: BoardFromFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -"),
			want:  48,
		},
		{
			// Black rook on a1 checks the white king on e1; only d2/e2/f2 escape
			name:  "King in check by rook",
			board: BoardFromFEN("8/8/8/8/8/8/8/r3K2k w - - 0 1"),
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moves := GenerateLegalMoves(&tt.board)
			assert.Equal(t, tt.want, len(moves))
		})
	}
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
