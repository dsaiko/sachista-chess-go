package engine

import "saiko.cz/sachista/index"
import "saiko.cz/sachista/board"

func BoardFromNotation(notations ...string) board.BitBoard {
	indices := make([]index.Index, len(notations))

	for i, n := range notations {
		indices[i] = index.FromNotation(n)
	}
	return BoardFromIndex(indices...)
}

func BoardFromIndex(indices ...index.Index) board.BitBoard {
	b := board.Empty

	for _, i := range indices {
		b |= IndexToBitBoard(i)
	}
	return b
}

func IndexToBitBoard(i index.Index) board.BitBoard {
	return 1 << i
}
