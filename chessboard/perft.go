package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

// CacheSize is the number of entries in the PerfT transposition table (64M entries, ~1 GB).
const CacheSize = 64 * 1024 * 1024

// cacheCountMask extracts the lower 56-bit count value from a packed cache entry.
const cacheCountMask = uint64(0x00FFFFFFFFFFFFFF)

// cacheDepthShift is the number of bits the depth value is shifted in a packed cache entry.
const cacheDepthShift = 56

// PerfTCache uses lockless hashing (XOR trick) with two uint64 arrays.
// On ARM64/x86-64, aligned 64-bit loads/stores are naturally atomic;
// the XOR consistency check in get() detects any torn or stale reads.
type PerfTCache struct {
	keys   *[CacheSize]uint64 // stores hash XOR value (for consistency check)
	values *[CacheSize]uint64 // stores packed (depth << 56) | count
}

// newPerfTCache allocates a new empty transposition cache.
func newPerfTCache() *PerfTCache {
	return &PerfTCache{
		keys:   new([CacheSize]uint64),
		values: new([CacheSize]uint64),
	}
}

// set stores a perft result in the cache using the lockless XOR trick (Robert Hyatt's method).
func (c *PerfTCache) set(hash uint64, depth int, count uint64) {
	idx := (CacheSize - 1) & hash
	value := (uint64(depth) << cacheDepthShift) | (count & cacheCountMask) //nolint:gosec
	c.values[idx] = value
	c.keys[idx] = hash ^ value
}

// get retrieves a cached perft result. Returns 0 on cache miss or consistency check failure.
func (c *PerfTCache) get(hash uint64, depth int) uint64 {
	idx := (CacheSize - 1) & hash
	value := c.values[idx]
	key := c.keys[idx]
	if key^value != hash {
		return 0
	}
	if value>>cacheDepthShift != uint64(depth) { //nolint:gosec
		return 0
	}
	return value & cacheCountMask
}

var cache = newPerfTCache()

// perfT1 is the single-threaded recursive perft algorithm with transposition caching.
// Takes Board by value to keep it on the stack and avoid heap allocations.
func perfT1(b Board, depth int) uint64 {
	if depth <= 0 {
		return 1
	}

	// if found in cache
	count := cache.get(b.ZobristHash, depth)
	if count != 0 {
		return count
	}

	attacks := attacks(&b, b.OpponentColor())
	isCheck := attacks&b.Pieces[b.NextMove][King] != 0

	handler := func(m Move) {
		sourceBitBoard := bitboard.BoardFromIndex(m.From)
		isKingMove := m.Piece == King

		// need to validate legality of move only in following cases
		needToValidate := isKingMove || isCheck || sourceBitBoard&attacks != 0 || m.IsEnPassant

		if depth == 1 {
			if !needToValidate {
				count++
			} else {
				nb := m.ApplyTo(b)
				if isOpponentsKingNotUnderCheck(&nb) {
					count++
				}
			}
		} else {
			nextBoard := m.ApplyTo(b)
			if !needToValidate || isOpponentsKingNotUnderCheck(&nextBoard) {
				count += perfT1(nextBoard, depth-1)
			}
		}
	}

	// generate pseudo legal moves
	generatePseudoLegalMoves(&b, handler)

	cache.set(b.ZobristHash, depth, count)
	return count
}

// PerfT runs a multi-threaded perft (performance test, move path enumeration) to the given depth.
// A goroutine is spawned for each top-level legal move to utilize all available CPU cores.
func PerfT(b *Board, depth int) uint64 {
	moves := GenerateLegalMoves(b)
	results := make(chan uint64, len(moves))

	// for each legal move, create a goroutine
	for _, m := range moves {
		nb := m.ApplyTo(*b)
		go func(b Board) {
			results <- perfT1(b, depth-1)
		}(nb)
	}

	// count results
	count := uint64(0)
	for range moves {
		count += <-results
	}

	return count
}
