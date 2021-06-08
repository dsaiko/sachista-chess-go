package index

import (
	"fmt"
	"strings"
)

func (i Index) FileIndex() int {
	return int(i) % 8
}

func (i Index) RankIndex() int {
	return int(i) / 8
}

func (i Index) String() string {
	return fmt.Sprintf("%c%c", 'a'+i.RankIndex(), '1'+i.FileIndex())
}

func FromNotation(notation string) Index {
	notation = strings.ToLower(notation)
	return Index(notation[0] - 'a' + ((notation[1] - '1') << 3))
}
