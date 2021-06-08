package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounts(t *testing.T) {
	assert.Equal(t, 0, Empty.PopCount())
	assert.Equal(t, 64, Universe.PopCount())
	assert.Equal(t, 28, Frame.PopCount())
}

func TestRanks(t *testing.T) {
	combined := Empty
	for _, b := range Rank {
		assert.Equal(t, 8, b.PopCount())
		combined ^= b
	}

	assert.Equal(t, Universe, combined)
}

func TestFile(t *testing.T) {
	combined := Empty
	for _, b := range File {
		assert.Equal(t, 8, b.PopCount())
		combined ^= b
	}

	assert.Equal(t, Universe, combined)
}

func TestA1H8(t *testing.T) {
	combined := Empty
	for _, b := range A1H8 {
		combined ^= b
	}

	assert.Equal(t, Universe, combined)
}

func TestA8H1(t *testing.T) {
	combined := Empty
	for _, b := range A8H1 {
		combined ^= b
	}

	assert.Equal(t, Universe, combined)
}

func TestFields(t *testing.T) {
	combined := Empty
	for _, b := range Fields {
		assert.Equal(t, 1, b.PopCount())
		combined ^= b
	}

	assert.Equal(t, Universe, combined)
}
