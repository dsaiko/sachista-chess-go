package bitboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert.Equal(t, 63, int(IndexH8-IndexA1))
}
