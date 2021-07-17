package index

import (
	"fmt"
	"strings"
)

// FileIndex return index of File (column) of the piece
func (i Index) FileIndex() int {
	return int(i) % 8
}

// RankIndex return index of Rank (row) of the piece
func (i Index) RankIndex() int {
	return int(i) / 8
}

// String returns notation representation of the position, for example a1 or b3
func (i Index) String() string {
	return fmt.Sprintf("%c%c", 'a'+i.FileIndex(), '1'+i.RankIndex())
}

// FromNotation returns position index from string representation
func FromNotation(notation string) Index {
	notation = strings.ToLower(notation)
	return Index(notation[0] - 'a' + ((notation[1] - '1') << 3))
}
