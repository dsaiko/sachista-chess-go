package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitPop(t *testing.T) {
	b := Universe

	count := 64
	for i := 0; i < 64; i++ {
		assert.Equal(t, count, b.PopCount())
		assert.Equal(t, i, b.BitPop())
		count--
	}
}
