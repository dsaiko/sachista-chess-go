package index

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex_String(t *testing.T) {
	f7 := fmt.Sprintf("%v", F7)
	a1 := fmt.Sprintf("%v", A1)
	h8 := fmt.Sprintf("%v", H8)

	assert.Equal(t, "f7", f7)
	assert.Equal(t, "a1", a1)
	assert.Equal(t, "h8", h8)
}

func TestFromNotation(t *testing.T) {
	notation := "A1"

	a1 := FromNotation(notation)
	h8 := FromNotation("h8")

	assert.Equal(t, "A1", notation)
	assert.Equal(t, A1, a1)
	assert.Equal(t, H8, h8)
}
