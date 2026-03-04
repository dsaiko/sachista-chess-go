package bitboard

import (
	"fmt"
	"strings"
)

// Index represents a square position on the board (0=a1, 7=h1, 56=a8, 63=h8).
type Index int

// File returns the file (column) of the square, ranging from 0 (a-file) to 7 (h-file).
func (i Index) File() int {
	return int(i) % 8
}

// Rank returns the rank (row) of the square, ranging from 0 (1st rank) to 7 (8th rank).
func (i Index) Rank() int {
	return int(i) / 8
}

// String returns the algebraic notation of the square (e.g. "a1", "h8").
func (i Index) String() string {
	return fmt.Sprintf("%c%c", 'a'+i.File(), '1'+i.Rank())
}

// IndexFromNotation returns the square index from algebraic notation (e.g. "e4" -> 28).
func IndexFromNotation(notation string) Index {
	notation = strings.ToLower(notation)
	return Index(notation[0] - 'a' + ((notation[1] - '1') << 3))
}
