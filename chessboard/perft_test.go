package chessboard

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPerfT(t *testing.T) {
	tests := []struct {
		board Board
		depth int
		want  uint64
	}{
		{
			board: StandardBoard(),
			depth: 5,
			want:  4_865_609,
		},
		{
			board: BoardFromFEN("8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - -"),
			depth: 5,
			want:  674_624,
		},
		{
			board: BoardFromFEN("r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1"),
			depth: 5,
			want:  15_833_292,
		},
		{
			board: BoardFromFEN("r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10"),
			depth: 5,
			want:  164_075_551,
		},
		{
			board: BoardFromFEN("r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1"),
			depth: 5,
			want:  15_833_292,
		},
		{
			board: BoardFromFEN("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -"),
			depth: 5,
			want:  193_690_690,
		},
		{
			board: BoardFromFEN("rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8"),
			depth: 5,
			want:  89_941_194,
		},
	}
	for _, tc := range tests {
		start := time.Now()
		got := PerfT(&tc.board, tc.depth)
		duration := time.Since(start)
		fmt.Printf("'%v': %v\n", tc.board.ToFEN(), duration)
		assert.Equal(t, tc.want, got)
	}
}

// TestPerfTShallow runs PerfT at small depths so that `make race` covers the goroutine fan-out
// and the shared cache without the cost of TestPerfT.
func TestPerfTShallow(t *testing.T) {
	tests := []struct {
		fen   string
		depth int
		want  uint64
	}{
		{fen: StandardBoardFEN, depth: 0, want: 1},
		{fen: StandardBoardFEN, depth: -1, want: 1},
		{fen: StandardBoardFEN, depth: 1, want: 20},
		{fen: StandardBoardFEN, depth: 3, want: 8_902},
		{fen: "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -", depth: 3, want: 97_862},
		{fen: "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - -", depth: 4, want: 43_238},
	}
	for _, tc := range tests {
		b := BoardFromFEN(tc.fen)
		assert.Equal(t, tc.want, PerfT(&b, tc.depth), "%v depth %v", tc.fen, tc.depth)
	}
}

func TestPerfTCache(t *testing.T) {
	c := newPerfTCache()
	const hash = uint64(0xDEADBEEF12345678)

	assert.Equal(t, uint64(0), c.get(hash, 3), "empty cache misses")

	c.set(hash, 3, 1234)
	assert.Equal(t, uint64(1234), c.get(hash, 3), "hit returns the stored count")
	assert.Equal(t, uint64(0), c.get(hash, 2), "depth mismatch misses")

	// a different hash mapping to the same slot must not see the entry
	other := hash + CacheSize
	assert.Equal(t, uint64(0), c.get(other, 3), "same slot, different hash misses")

	// the newer entry replaces the older one in the slot
	c.set(other, 3, 99)
	assert.Equal(t, uint64(99), c.get(other, 3))
	assert.Equal(t, uint64(0), c.get(hash, 3))

	// a key and value written by different set() calls fail the XOR check
	idx := (CacheSize - 1) & hash
	c.set(hash, 3, 1234)
	c.values[idx] ^= 1
	assert.Equal(t, uint64(0), c.get(hash, 3), "inconsistent key/value misses")

	// counts are stored in the low 56 bits
	c.set(hash, 5, cacheCountMask)
	assert.Equal(t, cacheCountMask, c.get(hash, 5))
}

func BenchmarkPerfT(b *testing.B) {
	board := StandardBoard()

	for b.Loop() {
		// start every iteration with an empty cache, otherwise all but the first run are pure cache hits
		b.StopTimer()
		cache = newPerfTCache()
		b.StartTimer()

		PerfT(&board, 6)
	}
}
