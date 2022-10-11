package bitboard

const BoardHeader = "  a b c d e f g h\n"

// Constants for Board of single piece on the bitboard
const (
	BoardA1 Board = 1
	BoardB1 Board = 1 << 1
	BoardC1 Board = 1 << 2
	BoardD1 Board = 1 << 3
	BoardE1 Board = 1 << 4
	BoardF1 Board = 1 << 5
	BoardG1 Board = 1 << 6
	BoardH1 Board = 1 << 7

	BoardA2 Board = 1 << 8
	BoardB2 Board = 1 << 9
	BoardC2 Board = 1 << 10
	BoardD2 Board = 1 << 11
	BoardE2 Board = 1 << 12
	BoardF2 Board = 1 << 13
	BoardG2 Board = 1 << 14
	BoardH2 Board = 1 << 15

	BoardA3 Board = 1 << 16
	BoardB3 Board = 1 << 17
	BoardC3 Board = 1 << 18
	BoardD3 Board = 1 << 19
	BoardE3 Board = 1 << 20
	BoardF3 Board = 1 << 21
	BoardG3 Board = 1 << 22
	BoardH3 Board = 1 << 23

	BoardA4 Board = 1 << 24
	BoardB4 Board = 1 << 25
	BoardC4 Board = 1 << 26
	BoardD4 Board = 1 << 27
	BoardE4 Board = 1 << 28
	BoardF4 Board = 1 << 29
	BoardG4 Board = 1 << 30
	BoardH4 Board = 1 << 31

	BoardA5 Board = 1 << 32
	BoardB5 Board = 1 << 33
	BoardC5 Board = 1 << 34
	BoardD5 Board = 1 << 35
	BoardE5 Board = 1 << 36
	BoardF5 Board = 1 << 37
	BoardG5 Board = 1 << 38
	BoardH5 Board = 1 << 39

	BoardA6 Board = 1 << 40
	BoardB6 Board = 1 << 41
	BoardC6 Board = 1 << 42
	BoardD6 Board = 1 << 43
	BoardE6 Board = 1 << 44
	BoardF6 Board = 1 << 45
	BoardG6 Board = 1 << 46
	BoardH6 Board = 1 << 47

	BoardA7 Board = 1 << 48
	BoardB7 Board = 1 << 49
	BoardC7 Board = 1 << 50
	BoardD7 Board = 1 << 51
	BoardE7 Board = 1 << 52
	BoardF7 Board = 1 << 53
	BoardG7 Board = 1 << 54
	BoardH7 Board = 1 << 55

	BoardA8 Board = 1 << 56
	BoardB8 Board = 1 << 57
	BoardC8 Board = 1 << 58
	BoardD8 Board = 1 << 59
	BoardE8 Board = 1 << 60
	BoardF8 Board = 1 << 61
	BoardG8 Board = 1 << 62
	BoardH8 Board = 1 << 63

	EmptyBoard Board = 0

	// UniverseBoard is Board with all pieces set
	UniverseBoard = ^EmptyBoard
)

var BoardFields = [...]Board{
	BoardA1, BoardB1, BoardC1, BoardD1, BoardE1, BoardF1, BoardG1, BoardH1,
	BoardA2, BoardB2, BoardC2, BoardD2, BoardE2, BoardF2, BoardG2, BoardH2,
	BoardA3, BoardB3, BoardC3, BoardD3, BoardE3, BoardF3, BoardG3, BoardH3,
	BoardA4, BoardB4, BoardC4, BoardD4, BoardE4, BoardF4, BoardG4, BoardH4,
	BoardA5, BoardB5, BoardC5, BoardD5, BoardE5, BoardF5, BoardG5, BoardH5,
	BoardA6, BoardB6, BoardC6, BoardD6, BoardE6, BoardF6, BoardG6, BoardH6,
	BoardA7, BoardB7, BoardC7, BoardD7, BoardE7, BoardF7, BoardG7, BoardH7,
	BoardA8, BoardB8, BoardC8, BoardD8, BoardE8, BoardF8, BoardG8, BoardH8,
}

var BoardRanks = [...]Board{
	BoardA1 | BoardB1 | BoardC1 | BoardD1 | BoardE1 | BoardF1 | BoardG1 | BoardH1,
	BoardA2 | BoardB2 | BoardC2 | BoardD2 | BoardE2 | BoardF2 | BoardG2 | BoardH2,
	BoardA3 | BoardB3 | BoardC3 | BoardD3 | BoardE3 | BoardF3 | BoardG3 | BoardH3,
	BoardA4 | BoardB4 | BoardC4 | BoardD4 | BoardE4 | BoardF4 | BoardG4 | BoardH4,
	BoardA5 | BoardB5 | BoardC5 | BoardD5 | BoardE5 | BoardF5 | BoardG5 | BoardH5,
	BoardA6 | BoardB6 | BoardC6 | BoardD6 | BoardE6 | BoardF6 | BoardG6 | BoardH6,
	BoardA7 | BoardB7 | BoardC7 | BoardD7 | BoardE7 | BoardF7 | BoardG7 | BoardH7,
	BoardA8 | BoardB8 | BoardC8 | BoardD8 | BoardE8 | BoardF8 | BoardG8 | BoardH8,
}

var BoardFiles = [...]Board{
	BoardA1 | BoardA2 | BoardA3 | BoardA4 | BoardA5 | BoardA6 | BoardA7 | BoardA8,
	BoardB1 | BoardB2 | BoardB3 | BoardB4 | BoardB5 | BoardB6 | BoardB7 | BoardB8,
	BoardC1 | BoardC2 | BoardC3 | BoardC4 | BoardC5 | BoardC6 | BoardC7 | BoardC8,
	BoardD1 | BoardD2 | BoardD3 | BoardD4 | BoardD5 | BoardD6 | BoardD7 | BoardD8,
	BoardE1 | BoardE2 | BoardE3 | BoardE4 | BoardE5 | BoardE6 | BoardE7 | BoardE8,
	BoardF1 | BoardF2 | BoardF3 | BoardF4 | BoardF5 | BoardF6 | BoardF7 | BoardF8,
	BoardG1 | BoardG2 | BoardG3 | BoardG4 | BoardG5 | BoardG6 | BoardG7 | BoardG8,
	BoardH1 | BoardH2 | BoardH3 | BoardH4 | BoardH5 | BoardH6 | BoardH7 | BoardH8,
}

var BoardFileA = BoardFiles[0]
var BoardFileH = BoardFiles[7]

var BoardRank1 = BoardRanks[0]
var BoardRank8 = BoardRanks[7]

var BoardFrame = BoardRank1 | BoardRank8 | BoardFileA | BoardFileH

var BoardA1H8 = [...]Board{
	BoardA8,
	BoardA7 | BoardB8,
	BoardA6 | BoardB7 | BoardC8,
	BoardA5 | BoardB6 | BoardC7 | BoardD8,
	BoardA4 | BoardB5 | BoardC6 | BoardD7 | BoardE8,
	BoardA3 | BoardB4 | BoardC5 | BoardD6 | BoardE7 | BoardF8,
	BoardA2 | BoardB3 | BoardC4 | BoardD5 | BoardE6 | BoardF7 | BoardG8,
	BoardA1 | BoardB2 | BoardC3 | BoardD4 | BoardE5 | BoardF6 | BoardG7 | BoardH8,
	BoardB1 | BoardC2 | BoardD3 | BoardE4 | BoardF5 | BoardG6 | BoardH7,
	BoardC1 | BoardD2 | BoardE3 | BoardF4 | BoardG5 | BoardH6,
	BoardD1 | BoardE2 | BoardF3 | BoardG4 | BoardH5,
	BoardE1 | BoardF2 | BoardG3 | BoardH4,
	BoardF1 | BoardG2 | BoardH3,
	BoardG1 | BoardH2,
	BoardH1,
}

var BoardA8H1 = [...]Board{
	BoardA1,
	BoardA2 | BoardB1,
	BoardA3 | BoardB2 | BoardC1,
	BoardA4 | BoardB3 | BoardC2 | BoardD1,
	BoardA5 | BoardB4 | BoardC3 | BoardD2 | BoardE1,
	BoardA6 | BoardB5 | BoardC4 | BoardD3 | BoardE2 | BoardF1,
	BoardA7 | BoardB6 | BoardC5 | BoardD4 | BoardE3 | BoardF2 | BoardG1,
	BoardA8 | BoardB7 | BoardC6 | BoardD5 | BoardE4 | BoardF3 | BoardG2 | BoardH1,
	BoardB8 | BoardC7 | BoardD6 | BoardE5 | BoardF4 | BoardG3 | BoardH2,
	BoardC8 | BoardD7 | BoardE6 | BoardF5 | BoardG4 | BoardH3,
	BoardD8 | BoardE7 | BoardF6 | BoardG5 | BoardH4,
	BoardE8 | BoardF7 | BoardG6 | BoardH5,
	BoardF8 | BoardG7 | BoardH6,
	BoardG8 | BoardH7,
	BoardH8,
}
