package generator

import (
	"fmt"
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

var cache PerfTCache

func allocateCache() {
	defer func() {
		if r := recover(); r != nil {
			cache.cacheSize /= 2
			fmt.Println(r, "Memory initialization failed. Reducing Cache Size:", cache.cacheSize)
			cache.entries = []PerfTCacheEntry{}
		}
	}()
	cache.entries = make([]PerfTCacheEntry, cache.cacheSize)
}

func init() {
	cache.cacheSize = 128 * 1024 * 1024
	for len(cache.entries) == 0 {
		allocateCache()
		if len(cache.entries) == 0 {
			if cache.cacheSize < 1024*1024 {
				panic("Memory initialization error")
			}
		}
	}
}

func PerfT(b *chessboard.Board, depth int) uint64 {
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

		if isKingMove || isCheck || (sourceBitBoard&attacks) != 0 || m.IsEnPassant {
			// need to validate move
			nextBoard := m.MakeMove(*b)
			if isOpponentsKingNotUnderCheck(nextBoard) {
				if depth == 1 {
					count++
				} else {
					count += PerfT(nextBoard, depth-1)
				}
			}
		} else {
			if depth == 1 {
				count++
			} else {
				// do not need to validate legality of the move
				nextBoard := m.MakeMove(*b)
				count += PerfT(nextBoard, depth-1)
			}
		}
	}

	// DEBUG OUTPUT FOR UTILS/PERFT-STOKFISH-CHECK.SH:
	// fmt.Printf("%v|%v|%v\n",b.ToFEN(), depth, count)

	cache.set(b.ZobristHash, depth, count)
	return count
}
