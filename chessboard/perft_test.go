package chessboard

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func BenchmarkPerfT(b *testing.B) {
	board := StandardBoard()

	for i := 0; i < b.N; i++ {
		PerfT(&board, 6)
	}
}
