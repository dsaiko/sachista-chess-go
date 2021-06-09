package chessboard

import (
	"fmt"
	"testing"
)

func TestBoard_ToFEN(t *testing.T) {
	b := StandardBoard()

	fmt.Println(b.ToFEN())
}
