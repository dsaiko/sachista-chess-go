package board

type BitBoard uint64

const (
	A1 BitBoard = 1
	B1 BitBoard = 1 << 1
	C1 BitBoard = 1 << 2
	D1 BitBoard = 1 << 3
	E1 BitBoard = 1 << 4
	F1 BitBoard = 1 << 5
	G1 BitBoard = 1 << 6
	H1 BitBoard = 1 << 7

	A2 BitBoard = 1 << 8
	B2 BitBoard = 1 << 9
	C2 BitBoard = 1 << 10
	D2 BitBoard = 1 << 11
	E2 BitBoard = 1 << 12
	F2 BitBoard = 1 << 13
	G2 BitBoard = 1 << 14
	H2 BitBoard = 1 << 15

	A3 BitBoard = 1 << 16
	B3 BitBoard = 1 << 17
	C3 BitBoard = 1 << 18
	D3 BitBoard = 1 << 19
	E3 BitBoard = 1 << 20
	F3 BitBoard = 1 << 21
	G3 BitBoard = 1 << 22
	H3 BitBoard = 1 << 23

	A4 BitBoard = 1 << 24
	B4 BitBoard = 1 << 25
	C4 BitBoard = 1 << 26
	D4 BitBoard = 1 << 27
	E4 BitBoard = 1 << 28
	F4 BitBoard = 1 << 29
	G4 BitBoard = 1 << 30
	H4 BitBoard = 1 << 31

	A5 BitBoard = 1 << 32
	B5 BitBoard = 1 << 33
	C5 BitBoard = 1 << 34
	D5 BitBoard = 1 << 35
	E5 BitBoard = 1 << 36
	F5 BitBoard = 1 << 37
	G5 BitBoard = 1 << 38
	H5 BitBoard = 1 << 39

	A6 BitBoard = 1 << 40
	B6 BitBoard = 1 << 41
	C6 BitBoard = 1 << 42
	D6 BitBoard = 1 << 43
	E6 BitBoard = 1 << 44
	F6 BitBoard = 1 << 45
	G6 BitBoard = 1 << 46
	H6 BitBoard = 1 << 47

	A7 BitBoard = 1 << 48
	B7 BitBoard = 1 << 49
	C7 BitBoard = 1 << 50
	D7 BitBoard = 1 << 51
	E7 BitBoard = 1 << 52
	F7 BitBoard = 1 << 53
	G7 BitBoard = 1 << 54
	H7 BitBoard = 1 << 55

	A8 BitBoard = 1 << 56
	B8 BitBoard = 1 << 57
	C8 BitBoard = 1 << 58
	D8 BitBoard = 1 << 59
	E8 BitBoard = 1 << 60
	F8 BitBoard = 1 << 61
	G8 BitBoard = 1 << 62
	H8 BitBoard = 1 << 63

	Empty    BitBoard = 0
	Universe          = ^Empty
)

var Fields = [...]BitBoard{
	A1, B1, C1, D1, E1, F1, G1, H1,
	A2, B2, C2, D2, E2, F2, G2, H2,
	A3, B3, C3, D3, E3, F3, G3, H3,
	A4, B4, C4, D4, E4, F4, G4, H4,
	A5, B5, C5, D5, E5, F5, G5, H5,
	A6, B6, C6, D6, E6, F6, G6, H6,
	A7, B7, C7, D7, E7, F7, G7, H7,
	A8, B8, C8, D8, E8, F8, G8, H8,
}

var Rank = [...]BitBoard{
	A1 | B1 | C1 | D1 | E1 | F1 | G1 | H1,
	A2 | B2 | C2 | D2 | E2 | F2 | G2 | H2,
	A3 | B3 | C3 | D3 | E3 | F3 | G3 | H3,
	A4 | B4 | C4 | D4 | E4 | F4 | G4 | H4,
	A5 | B5 | C5 | D5 | E5 | F5 | G5 | H5,
	A6 | B6 | C6 | D6 | E6 | F6 | G6 | H6,
	A7 | B7 | C7 | D7 | E7 | F7 | G7 | H7,
	A8 | B8 | C8 | D8 | E8 | F8 | G8 | H8,
}

var File = [...]BitBoard{
	A1 | A2 | A3 | A4 | A5 | A6 | A7 | A8,
	B1 | B2 | B3 | B4 | B5 | B6 | B7 | B8,
	C1 | C2 | C3 | C4 | C5 | C6 | C7 | C8,
	D1 | D2 | D3 | D4 | D5 | D6 | D7 | D8,
	E1 | E2 | E3 | E4 | E5 | E6 | E7 | E8,
	F1 | F2 | F3 | F4 | F5 | F6 | F7 | F8,
	G1 | G2 | G3 | G4 | G5 | G6 | G7 | G8,
	H1 | H2 | H3 | H4 | H5 | H6 | H7 | H8,
}

var FileA = File[0]
var FileH = File[7]

var Rank1 = Rank[0]
var Rank8 = Rank[7]

var Frame = Rank1 | Rank8 | FileA | FileH

var A1H8 = [...]BitBoard{
	A8,
	A7 | B8,
	A6 | B7 | C8,
	A5 | B6 | C7 | D8,
	A4 | B5 | C6 | D7 | E8,
	A3 | B4 | C5 | D6 | E7 | F8,
	A2 | B3 | C4 | D5 | E6 | F7 | G8,
	A1 | B2 | C3 | D4 | E5 | F6 | G7 | H8,
	B1 | C2 | D3 | E4 | F5 | G6 | H7,
	C1 | D2 | E3 | F4 | G5 | H6,
	D1 | E2 | F3 | G4 | H5,
	E1 | F2 | G3 | H4,
	F1 | G2 | H3,
	G1 | H2,
	H1,
}

var A8H1 = [...]BitBoard{
	A1,
	A2 | B1,
	A3 | B2 | C1,
	A4 | B3 | C2 | D1,
	A5 | B4 | C3 | D2 | E1,
	A6 | B5 | C4 | D3 | E2 | F1,
	A7 | B6 | C5 | D4 | E3 | F2 | G1,
	A8 | B7 | C6 | D5 | E4 | F3 | G2 | H1,
	B8 | C7 | D6 | E5 | F4 | G3 | H2,
	C8 | D7 | E6 | F5 | G4 | H3,
	D8 | E7 | F6 | G5 | H4,
	E8 | F7 | G6 | H5,
	F8 | G7 | H6,
	G8 | H7,
	H8,
}

const header = "  a b c d e f g h\n"
