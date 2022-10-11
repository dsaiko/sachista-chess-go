package bitboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounts(t *testing.T) {
	assert.Equal(t, 0, EmptyBoard.PopCount())
	assert.Equal(t, 64, UniverseBoard.PopCount())
	assert.Equal(t, 28, BoardFrame.PopCount())
}

func TestRanks(t *testing.T) {
	combined := EmptyBoard
	for _, b := range BoardRanks {
		assert.Equal(t, 8, b.PopCount())
		combined ^= b
	}

	assert.Equal(t, UniverseBoard, combined)
}

func TestFile(t *testing.T) {
	combined := EmptyBoard
	for _, b := range BoardFiles {
		assert.Equal(t, 8, b.PopCount())
		combined ^= b
	}

	assert.Equal(t, UniverseBoard, combined)
}

func TestA1H8(t *testing.T) {
	combined := EmptyBoard
	for _, b := range BoardA1H8 {
		combined ^= b
	}

	assert.Equal(t, UniverseBoard, combined)
}

func TestA8H1(t *testing.T) {
	combined := EmptyBoard
	for _, b := range BoardA8H1 {
		combined ^= b
	}

	assert.Equal(t, UniverseBoard, combined)
}

func TestFields(t *testing.T) {
	combined := EmptyBoard
	for _, b := range BoardFields {
		assert.Equal(t, 1, b.PopCount())
		combined ^= b
	}

	assert.Equal(t, UniverseBoard, combined)
}
