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
