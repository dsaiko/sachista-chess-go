package zobrist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewZobrist(t *testing.T) {
	z := NewKeys()

	// Check that all generated random keys are unique

	set := make(map[uint64]bool)

	check := func(r uint64) {
		assert.NotZero(t, r)
		if _, ok := set[r]; ok {
			//key already exists !?
			assert.Fail(t, "Non unique zobrist random number detected!")
		}
		set[r] = true
	}

	for _, v1 := range z.Pieces {
		for _, v2 := range v1 {
			for _, v3 := range v2 {
				check(v3)
			}
		}
	}

	for _, v1 := range z.Castling {
		for _, v2 := range v1 {
			check(v2)
		}
	}

	for _, v1 := range z.EnPassant {
		check(v1)
	}

	check(z.Side)
}
