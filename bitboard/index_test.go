package bitboard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex_String(t *testing.T) {
	f7 := fmt.Sprintf("%v", IndexF7)
	a1 := fmt.Sprintf("%v", IndexA1)
	h8 := fmt.Sprintf("%v", IndexH8)

	assert.Equal(t, "f7", f7)
	assert.Equal(t, "a1", a1)
	assert.Equal(t, "h8", h8)
}

func TestFromNotation(t *testing.T) {
	notation := "A1"

	a1 := IndexFromNotation(notation)
	h8 := IndexFromNotation("h8")

	assert.Equal(t, "A1", notation)
	assert.Equal(t, IndexA1, a1)
	assert.Equal(t, IndexH8, h8)
}

func TestIndex_FileRank(t *testing.T) {
	assert.Equal(t, 0, IndexA1.File())
	assert.Equal(t, 0, IndexA1.Rank())

	assert.Equal(t, 7, IndexH8.File())
	assert.Equal(t, 7, IndexH8.Rank())

	// e4: e-file = 4 (0-indexed), rank 4 = 3 (0-indexed)
	assert.Equal(t, 4, IndexE4.File())
	assert.Equal(t, 3, IndexE4.Rank())

	// f7: f-file = 5, rank 7 = 6
	assert.Equal(t, 5, IndexF7.File())
	assert.Equal(t, 6, IndexF7.Rank())

	assert.Equal(t, 7, IndexH1.File())
	assert.Equal(t, 0, IndexH1.Rank())

	assert.Equal(t, 0, IndexA8.File())
	assert.Equal(t, 7, IndexA8.Rank())
}
