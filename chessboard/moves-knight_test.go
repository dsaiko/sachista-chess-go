package chessboard

import (
	"testing"
)

func TestKnightMoves(t *testing.T) {
	testMovesFromString(t, 6, `
       a b c d e f g h
     8 - - - - - - - - 8
     7 - - - - - - - - 7
     6 - - - - - - - - 6
     5 - - - - - - - - 5
     4 - - - - - - - - 4
     3 - - - - - - - - 3
     2 - - - - - - - - 2
     1 - N - - - - N - 1
       a b c d e f g h
`)

	testMovesFromString(t, 14, `
        a b c d e f g h
      8 - - - - - - - - 8
      7 - - - - - - - - 7
      6 - - - - N - - - 6
      5 - - - - - - - - 5
      4 - - - - - N - - 4
      3 - - - - - - - - 3
      2 - - - - - - - - 2
      1 - - - - - - - - 1
        a b c d e f g h
`)

	testMovesFromString(t, 23, `
        a b c d e f g h
      8 - N - - - - - - 8
      7 - - - - - - - - 7
      6 - - - - N - - - 6
      5 - - - - - - - - 5
      4 - - - - - N - - 4
      3 - N - - - - - - 3
      2 - - - - - - - - 2
      1 - - - - - - - - 1
        a b c d e f g h
`)
}
