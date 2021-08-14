package generator

import (
	"testing"
)

func TestPawnMoves(t *testing.T) {
	testMovesFromString(t, 2, `
a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - - 3
2 - - - - - - - P 2
1 - - - - - - - - 1
 a b c d e f g h
`)
	testMovesFromString(t, 0, `
  a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - - k 3
2 - - - - - - - P 2
1 - - - - - - - - 1
  a b c d e f g h
`)
	testMovesFromString(t, 3, `
  a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - - - - - n - 3
2 - - - - - - - P 2
1 - - - - - - - - 1
  a b c d e f g h
`)
	testMovesFromString(t, 4, `
  a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - - - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - n - n - n - 3
2 - - - P - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`)
	testMovesFromString(t, 2, `
  a b c d e f g h
8 - - - - - - - - 8
7 - - - - - - - - 7
6 - - n - - - - - 6
5 - - - P - - - - 5
4 - - - - - - - - 4
3 - - n - n - n - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`)
	testMovesFromString(t, 12, `
  a b c d e f g h
8 - - - n - n - - 8
7 - - - - P - - - 7
6 - - n - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - n - n - n - 3
2 - - - - - - - - 2
1 - - - - - - - - 1
  a b c d e f g h
`)

	testMovesFromString(t, 14, `
  a b c d e f g h
8 - - - n - n - - 8
7 - - - - P - - - 7
6 - - n - - - - - 6
5 - - - - - - - - 5
4 - - - - - - - - 4
3 - - n - n - n - 3
2 - - - - - - - - 2
1 N - - - - - - - 1
  a b c d e f g h
`)

	testMovesFromString(t, 21, `
  a b c d e f g h
8 - - - n - n - - 8
7 - - - - P - - - 7
6 - - n - - - - - 6
5 - - - - - N - - 5
4 - - - - - - - - 4
3 - - n - n - n - 3
2 - - - - - - - - 2
1 N - - - - - - - 1
  a b c d e f g h
`)

	testMovesFromFEN(t, 2, "111n1n111/11111111/11n11pP1/11111111/11111111/11n1n1n1/11111111/11111111 w KQkq f7 0 1")
}
