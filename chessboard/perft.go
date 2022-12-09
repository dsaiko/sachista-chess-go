package chessboard

import (
	"saiko.cz/sachista/bitboard"
	"sync/atomic"
)

// PerfTCacheEntry cache record structure
type PerfTCacheEntry struct {
	hash  uint64
	depth int
	count uint64
}

const CacheSize = 64 * 1024 * 1024

// PerfTCache cache for repeated moves
type PerfTCache [CacheSize]atomic.Value

// set cache item - atomic method
func (c *PerfTCache) set(hash uint64, depth int, count uint64) {
	c[(CacheSize-1)&hash].Store(PerfTCacheEntry{
		hash:  hash,
		depth: depth,
		count: count,
	})
}

// get cache item - atomic method
func (c *PerfTCache) get(hash uint64, depth int) uint64 {
	entry, ok := c[(CacheSize-1)&hash].Load().(PerfTCacheEntry)
	if !ok || entry.hash != hash || entry.depth != depth {
		return 0
	}
	return entry.count
}

var cache = PerfTCache{}

// perfT1 single threaded min/max algorithm for searching the moves
func perfT1(b *Board, depth int) uint64 {
	if depth <= 0 {
		return 1
	}

	// if found in cache
	count := cache.get(b.ZobristHash, depth)
	if count != 0 {
		return count
	}

	attacks := attacks(b, b.OpponentColor())
	isCheck := attacks&b.Pieces[b.NextMove][King] != 0

	handler := func(m Move) {
		sourceBitBoard := bitboard.BoardFromIndex(m.From)
		isKingMove := m.Piece == King

		// need to validate legality of move only in following cases
		needToValidate := isKingMove || isCheck || sourceBitBoard&attacks != 0 || m.IsEnPassant

		if depth == 1 {
			if !needToValidate || isOpponentsKingNotUnderCheck(m.ApplyTo(*b)) {
				count++
			}
		} else {
			nextBoard := m.ApplyTo(*b)
			if !needToValidate || isOpponentsKingNotUnderCheck(nextBoard) {
				count += perfT1(nextBoard, depth-1)
			}
		}
	}

	// generate pseudo legal moves
	generatePseudoLegalMoves(b, handler)

	// DEBUG OUTPUT FOR UTILS/PERFT-STOKFISH-CHECK.SH:
	// fmt.Printf("%v|%v|%v\n",b.ToFEN(), depth, count)

	cache.set(b.ZobristHash, depth, count)
	return count
}

// PerfT multithreading perfT algorithm
// goroutine are spawned on each of first set of legal moves
func PerfT(b *Board, depth int) uint64 {
	moves := GenerateLegalMoves(b)
	results := make(chan uint64, len(moves))

	// for each legal move, create a goroutine
	for _, m := range moves {
		go func(b *Board) {
			results <- perfT1(b, depth-1)
		}(m.ApplyTo(*b))
	}

	// count results
	count := uint64(0)
	for i := 0; i < len(moves); i++ {
		count += <-results
	}

	return count
}
