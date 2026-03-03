package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

const CacheSize = 64 * 1024 * 1024

// PerfTCache uses lockless hashing (XOR trick) with two uint64 arrays.
// On ARM64/x86-64, aligned 64-bit loads/stores are naturally atomic;
// the XOR consistency check in get() detects any torn or stale reads.
type PerfTCache struct {
	keys   *[CacheSize]uint64 // stores hash XOR value (for consistency check)
	values *[CacheSize]uint64 // stores packed (depth << 56) | count
}

func newPerfTCache() *PerfTCache {
	return &PerfTCache{
		keys:   new([CacheSize]uint64),
		values: new([CacheSize]uint64),
	}
}

// set cache item using lockless XOR trick
func (c *PerfTCache) set(hash uint64, depth int, count uint64) {
	idx := (CacheSize - 1) & hash
	value := (uint64(depth) << 56) | (count & 0x00FFFFFFFFFFFFFF)
	c.values[idx] = value
	c.keys[idx] = hash ^ value
}

// get cache item - returns 0 on miss
func (c *PerfTCache) get(hash uint64, depth int) uint64 {
	idx := (CacheSize - 1) & hash
	value := c.values[idx]
	key := c.keys[idx]
	if key^value != hash {
		return 0
	}
	if value>>56 != uint64(depth) {
		return 0
	}
	return value & 0x00FFFFFFFFFFFFFF
}

var cache = newPerfTCache()

// perfT1 single threaded perft algorithm - takes Board by value to keep it on stack
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

// PerfT multithreading perfT algorithm
// goroutine are spawned on each of first set of legal moves
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
