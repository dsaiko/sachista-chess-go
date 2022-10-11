package bitboard

import (
	"fmt"
	"strings"
)

type Index int

// File return index of File (column) of the piece
func (i Index) File() int {
	return int(i) % 8
}

// Rank return index of Rank (row) of the piece
func (i Index) Rank() int {
	return int(i) / 8
}

// String returns notation representation of the position, for example a1 or b3
func (i Index) String() string {
	return fmt.Sprintf("%c%c", 'a'+i.File(), '1'+i.Rank())
}

// IndexFromNotation returns position index from string representation
func IndexFromNotation(notation string) Index {
	notation = strings.ToLower(notation)
	return Index(notation[0] - 'a' + ((notation[1] - '1') << 3))
}
