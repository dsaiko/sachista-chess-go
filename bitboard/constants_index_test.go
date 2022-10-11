package bitboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	assert.Equal(t, 63, int(IndexH8-IndexA1))
}
