package index

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	a1 := fmt.Sprintf("%v", A1)
	h8 := fmt.Sprintf("%v", H8)
	assert.Equal(t, "a1", a1)
	assert.Equal(t, "h8", h8)
}

func TestIndexFromNotation(t *testing.T) {
	notation := "A1"

	a1 := FromNotation(notation)
	h8 := FromNotation("h8")

	assert.Equal(t, "A1", notation)
	assert.Equal(t, A1, a1)
	assert.Equal(t, H8, h8)
}
