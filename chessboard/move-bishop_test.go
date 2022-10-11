package chessboard

import (
	"testing"
)

func TestBishopMoves(t *testing.T) {
	testMovesFromString(t, 7, `
   a b c d e f g h
 8 - - - - - - - - 8
 7 - - - - - - - - 7
 6 - - - - - - - - 6
 5 - - - - - - - - 5
 4 - - - - - - - - 4
 3 - - - - - - - - 3
 2 - - - - - - - - 2
 1 - - - - - - B - 1
   a b c d e f g h
`)
	testMovesFromString(t, 13, `
   a b c d e f g h
 8 - - - - - - - - 8
 7 - - - - - - - - 7
 6 - - - - - - - - 6
 5 - - - B - - - - 5
 4 - - - - - - - - 4
 3 - - - - - - - - 3
 2 - - - - - - - - 2
 1 - - - - - - - - 1
   a b c d e f g h
`)
	testMovesFromString(t, 4, `
   a b c d e f g h
 8 - - - - - - - - 8
 7 - - - - - - - - 7
 6 - - n - n - - - 6
 5 - - - B - - - - 5
 4 - - n - n - - - 4
 3 - - - - - - - - 3
 2 - - - - - - - - 2
 1 - - - - - - - - 1
   a b c d e f g h
`)

	testMovesFromString(t, 4, `
  a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - n n n - - - 6
5 - - n B n - - - 5
4 - - n n n - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`)
	testMovesFromString(t, 17, `
 a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - K - - - - - 6
5 - - - B - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
 a b c d e f g h
`)
	testMovesFromString(t, 7, `
  a b c d e f g h
8 B - - - - - - - 8
7 - K - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`)
	testMovesFromString(t, 10, `
  a b c d e f g h
8 - - - - - - - - 8
7 - B - - - - - - 7
6 - - K - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`)
}
