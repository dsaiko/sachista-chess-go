package chessboard

import "testing"

func TestKingMoves(t *testing.T) {
	testMovesFromString(t, 8, `
  a b c d e f g h
8 - n - - - - - - 8
7 - - - - - - - - 7
6 - - - - n - - - 6
5 - - - - - - - - 5
4 - - - - - n - - 4
3 - n - - - - - - 3
2 - - - K - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`)

	testMovesFromString(t, 26, `
  a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
  a b c d e f g h
`)

	testMovesFromString(t, 24, `
 a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - n - - 3
2 - - - - - - - - 2
1 R - - - K - - R 1
 a b c d e f g h
`)
}
