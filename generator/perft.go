package generator

import (
	"sync"

	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
)

type PerfTCacheEntry struct {
	hash  uint64
	depth int
	count uint64
}

type PerfTCache struct {
	entries   []PerfTCacheEntry
	cacheSize uint64
	mux       sync.Mutex
}

func (c *PerfTCache) set(hash uint64, depth int, count uint64) {
	c.mux.Lock()
	defer c.mux.Unlock()

	i := (c.cacheSize - 1) & hash

	c.entries[i].hash = hash
	c.entries[i].depth = depth
	c.entries[i].count = count
}

func (c *PerfTCache) get(hash uint64, depth int) uint64 {
	c.mux.Lock()
	defer c.mux.Unlock()

	i := (c.cacheSize - 1) & hash

	if c.entries[i].hash == hash && c.entries[i].depth == depth {
		return c.entries[i].count
	}

	return 0
}

func newCache(size uint64) *PerfTCache {
	var cache PerfTCache
	cache.cacheSize = size
	cache.entries = make([]PerfTCacheEntry, cache.cacheSize)
	return &cache
}

func perfT1(cache *PerfTCache, b *chessboard.Board, depth int) uint64 {
	count := cache.get(b.ZobristHash, depth)
	if count != 0 {
		return count
	}

	moves := make([]Move, 0, constants.MovesCacheInitialCapacity)
	GeneratePseudoLegalMoves(b, &moves)

	attacks := Attacks(b, b.OpponentColor())
	isCheck := attacks&b.Pieces[b.NextMove][chessboard.King] != 0

	for _, m := range moves {
		sourceBitBoard := bitboard.FromIndex1(m.From)
		isKingMove := m.Piece == chessboard.King

		needToValidate := isKingMove || isCheck || sourceBitBoard&attacks != 0 || m.IsEnPassant

		if depth == 1 {
			if needToValidate {
				// need to validate move
				nextBoard := m.MakeMove(*b)
				if isOpponentsKingNotUnderCheck(nextBoard) {
					count++
				}
			} else {
				count++
			}
		} else {
			nextBoard := m.MakeMove(*b)
			if needToValidate {
				if isOpponentsKingNotUnderCheck(nextBoard) {
					count += perfT1(cache, nextBoard, depth-1)
				}
			} else {
				count += perfT1(cache, nextBoard, depth-1)
			}
		}
	}

	// DEBUG OUTPUT FOR UTILS/PERFT-STOKFISH-CHECK.SH:
	// fmt.Printf("%v|%v|%v\n",b.ToFEN(), depth, count)

	cache.set(b.ZobristHash, depth, count)
	return count
}

func PerfT(b *chessboard.Board, depth int) uint64 {
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

	for _, m := range moves {
		wg.Add(1)
		nextBoard := m.MakeMove(*b)
		go func(cache *PerfTCache, b *chessboard.Board) {
			defer wg.Done()
			results <- perfT1(cache, b, depth-1)
		}(cache, nextBoard)
	}

	wg.Wait()
	close(results)

	count := uint64(0)
	for res := range results {
		count += res
	}

	return count
}
