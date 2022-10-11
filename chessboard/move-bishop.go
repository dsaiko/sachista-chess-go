package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

var bishopMagicA8H1 = [...]bitboard.Board{
	0x0,
	0x0,
	0x0101010101010100,
	0x0101010101010100,
	0x0101010101010100,
	0x0101010101010100,
	0x0101010101010100,
	0x0101010101010100,
	0x0080808080808080,
	0x0040404040404040,
	0x0020202020202020,
	0x0010101010101010,
	0x0008080808080808,
	0x0,
	0x0,
}

var bishopMagicA1H8 = [...]bitboard.Board{
	0x0,
	0x0,
	0x0101010101010100,
	0x0101010101010100,
	0x0101010101010100,
	0x0101010101010100,
	0x0101010101010100,
	0x0101010101010100,
	0x8080808080808000,
	0x4040404040400000,
	0x2020202020000000,
	0x1010101000000000,
	0x0808080000000000,
	0x0,
	0x0,
}

var bishopA1H8Index [bitboard.NumberOfSquares]int
var bishopA8H1Index [bitboard.NumberOfSquares]int

var bishopMoveA1H8Mask [bitboard.NumberOfSquares]bitboard.Board
var bishopMoveA1H8Magic [bitboard.NumberOfSquares]bitboard.Board
var bishopMoveA8H1Mask [bitboard.NumberOfSquares]bitboard.Board
var bishopMoveA8H1Magic [bitboard.NumberOfSquares]bitboard.Board

var bishopMoveA1H8Attacks [bitboard.NumberOfSquares][bitboard.NumberOfSquares]bitboard.Board
var bishopMoveA8H1Attacks [bitboard.NumberOfSquares][bitboard.NumberOfSquares]bitboard.Board

func init() {
	// for all fields
	for i := 0; i < bitboard.NumberOfSquares; i++ {
		// compute index of diagonal for the field
		bishopA8H1Index[i] = bitboard.Index(i).File() + bitboard.Index(i).Rank()%8
		bishopA1H8Index[i] = bitboard.Index(i).File() + 7 - bitboard.Index(i).Rank()%8

		// compute 6-bit diagonal for the field
		bishopMoveA8H1Mask[i] = bitboard.BoardA8H1[bishopA8H1Index[i]] & ^bitboard.BoardFrame
		bishopMoveA1H8Mask[i] = bitboard.BoardA1H8[bishopA1H8Index[i]] & ^bitboard.BoardFrame

		// index magic multiplier for the field
		bishopMoveA8H1Magic[i] = bishopMagicA8H1[bishopA8H1Index[i]]
		bishopMoveA1H8Magic[i] = bishopMagicA1H8[bishopA1H8Index[i]]
	}

	// precompute A1H8 moves
	// i is field index
	// n is 6 bit configuration

	// for all fields
	for i := 0; i < bitboard.NumberOfSquares; i++ {
		// for all possible diagonal states
		for n := 0; n < bitboard.NumberOfSquares; n++ {
			// get the diagonal
			diagonal := bitboard.BoardA1H8[bishopA1H8Index[i]]

			// reconstruct the state (number) into the diagonal

			// get the left/bottom bit - start of diagonal
			for diagonal.ShiftedOneSouthWest() != bitboard.EmptyBoard {
				diagonal = diagonal.ShiftedOneSouthWest()
			}

			board := bitboard.EmptyBoard

			m := n
			// traverse diagonal and set bits according to N
			for diagonal != bitboard.EmptyBoard {
				// shift up by one
				diagonal = diagonal.ShiftedOneNorthEast()
				if (m & 1) != 0 {
					board |= diagonal
				}
				m >>= 1
			}

			// make it 6-bit only
			board &= ^bitboard.BoardFrame

			// compute possible moves
			moves := bitboard.EmptyBoard

			// set piece to Ith index
			piece := bitboard.BoardFromIndex(bitboard.Index(i))

			// move in one direction
			for piece != bitboard.EmptyBoard {
				piece = piece.ShiftedOneNorthEast()
				moves |= piece

				// end when there is another piece (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					piece = bitboard.EmptyBoard
				}
			}

			// set piece back to Ith index
			piece = bitboard.BoardFromIndex(bitboard.Index(i))

			// move in the other direction
			for piece != bitboard.EmptyBoard {
				piece = piece.ShiftedOneSouthWest()
				moves |= piece

				// end when there is another piece (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					piece = bitboard.EmptyBoard
				}
			}

			// remember the moves for Ith field within Nth occupancy state
			bishopMoveA1H8Attacks[i][n] = moves
		}
	}

	// precompute A8H1 moves
	// i is field index
	// n is 6 bit configuration
	// for all fields
	for i := 0; i < bitboard.NumberOfSquares; i++ {
		// for all possible diagonal states
		for n := 0; n < bitboard.NumberOfSquares; n++ {
			// get the diagonal
			diagonal := bitboard.BoardA8H1[bishopA8H1Index[i]]

			// get the left/top bit - start of the diagonal
			for diagonal.ShiftedOneNorthWest() != bitboard.EmptyBoard {
				diagonal = diagonal.ShiftedOneNorthWest()
			}

			// traverse diagonal and set bits according to N
			board := bitboard.EmptyBoard

			m := n
			for diagonal != bitboard.EmptyBoard {
				// shift down by one
				diagonal = diagonal.ShiftedOneSouthEast()
				if (m & 1) != 0 {
					board |= diagonal
				}
				m >>= 1
			}

			// make it 6-bit only
			board &= ^bitboard.BoardFrame

			// pre-compute moves
			moves := bitboard.EmptyBoard

			// set the piece to Ith position
			piece := bitboard.BoardFromIndex(bitboard.Index(i))

			// move one direction
			for piece != bitboard.EmptyBoard {
				piece = piece.ShiftedOneNorthWest()
				moves |= piece
				// end when there is another piece (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					piece = 0
				}
			}

			// set the piece back to Ith position
			piece = bitboard.BoardFromIndex(bitboard.Index(i))

			// move the other direction
			for piece != bitboard.EmptyBoard {
				piece = piece.ShiftedOneSouthEast()
				moves |= piece
				// end when there is another piece (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					piece = 0
				}
			}
			bishopMoveA8H1Attacks[i][n] = moves
		}
	}
}

func oneBishopAttacks(sourceIndex bitboard.Index, allPieces bitboard.Board) bitboard.Board {
	stateIndexA8H1 := ((allPieces & bishopMoveA8H1Mask[sourceIndex]) * bishopMoveA8H1Magic[sourceIndex]) >> 57
	stateIndexA1H8 := ((allPieces & bishopMoveA1H8Mask[sourceIndex]) * bishopMoveA1H8Magic[sourceIndex]) >> 57

	// add attacks
	return bishopMoveA8H1Attacks[sourceIndex][stateIndexA8H1] | bishopMoveA1H8Attacks[sourceIndex][stateIndexA1H8]
}

func bishopAttacks(board *Board, color Color) bitboard.Board {
	pieces := board.Pieces[color][Bishop] | board.Pieces[color][Queen]
	attacks := bitboard.EmptyBoard

	// for all rooks
	allPieces := board.AllPieces()
	var i bitboard.Index
	for pieces != bitboard.EmptyBoard {
		i, pieces = pieces.BitPop()
		attacks |= oneBishopAttacks(i, allPieces)
	}
	return attacks
}

func bishopMoves(board *Board, handler MoveHandler) {
	movingPiece := Bishop
	bishop := board.Pieces[board.NextMove][movingPiece]

	allPieces := board.AllPieces()
	boardAvailable := board.BoardAvailableToAttack()

	var fromIndex, toIndex bitboard.Index

	for i := 0; i < 2; i++ { // bishops and queens
		// for all rooks
		for bishop != bitboard.EmptyBoard {
			// get next rook

			fromIndex, bishop = bishop.BitPop()
			movesBoard := oneBishopAttacks(fromIndex, allPieces) & boardAvailable

			// for all moves
			for movesBoard != bitboard.EmptyBoard {
				toIndex, movesBoard = movesBoard.BitPop()
				handler(Move{Piece: movingPiece, From: fromIndex, To: toIndex})
			}
		}
		// switch to queen
		movingPiece = Queen
		bishop = board.Pieces[board.NextMove][movingPiece]
	}
}
