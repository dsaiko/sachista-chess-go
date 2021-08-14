package generator

import (
	"sync"

	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
)

// PerfTCacheEntry cache record structure
type PerfTCacheEntry struct {
	hash  uint64
	depth int
	count uint64
}

// PerfTCache cache for repeated moves
type PerfTCache struct {
	entries   []PerfTCacheEntry
	cacheSize uint64
	mux       sync.Mutex
}

// set cache item - synchronized method
func (c *PerfTCache) set(hash uint64, depth int, count uint64) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry := &c.entries[(c.cacheSize-1)&hash]

	entry.hash = hash
	entry.depth = depth
	entry.count = count
}

// get cache item - synchronized method
func (c *PerfTCache) get(hash uint64, depth int) uint64 {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry := &c.entries[(c.cacheSize-1)&hash]

	if entry.hash == hash && entry.depth == depth {
		return entry.count
	}

	return 0
}

// newCache of requested size
func newCache(size uint64) *PerfTCache {
	var cache PerfTCache
	cache.cacheSize = size
	cache.entries = make([]PerfTCacheEntry, size)
	return &cache
}

// perfT1 single threaded min/max algorithm for searching the moves
func perfT1(cache *PerfTCache, b chessboard.Board, depth int) uint64 {

	// if found in cache
	count := cache.get(b.ZobristHash, depth)
	if count != 0 {
		return count
	}

	// generate pseudo legal moves
	moves := generatePseudoLegalMoves(b)

	attacks := attacks(b, b.OpponentColor())
	isCheck := attacks&b.Pieces[b.NextMove][chessboard.King] != 0

	for _, m := range moves {
		sourceBitBoard := bitboard.FromIndex1(m.From)
		isKingMove := m.Piece == chessboard.King

		// need to validate legality of move only in following cases
		needToValidate := isKingMove || isCheck || sourceBitBoard&attacks != 0 || m.IsEnPassant

		if depth == 1 {
			if !needToValidate || isOpponentsKingNotUnderCheck(m.MakeMove(b)) {
				count++
			}
		} else {
			nextBoard := m.MakeMove(b)
			if !needToValidate || isOpponentsKingNotUnderCheck(nextBoard) {
				count += perfT1(cache, nextBoard, depth-1)
			}
		}
	}

	// DEBUG OUTPUT FOR UTILS/PERFT-STOKFISH-CHECK.SH:
	// fmt.Printf("%v|%v|%v\n",b.ToFEN(), depth, count)

	cache.set(b.ZobristHash, depth, count)
	return count
}

// PerfT multithreading perfT algorithm
// goroutine are spawned on each of first set of legal moves
func PerfT(b chessboard.Board, depth int) uint64 {
	if depth <= 0 {
		return 1
	}

	moves := GenerateLegalMoves(b)
	if depth == 1 {
		return uint64(len(moves))
	}

	cache := newCache(16 * 1024 * 1024)
	results := make(chan uint64, len(moves))
	var wg sync.WaitGroup

	// for each legal move, create a goroutine
	for _, m := range moves {
		wg.Add(1)
		go func(b chessboard.Board) {
			defer wg.Done()
			results <- perfT1(cache, b, depth-1)
		}(m.MakeMove(b))
	}

	// wait for all and close results
	wg.Wait()
	close(results)

	// count results
	count := uint64(0)
	for res := range results {
		count += res
	}

	return count
}
