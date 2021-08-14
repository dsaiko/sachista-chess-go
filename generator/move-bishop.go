package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
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

var bishopA1H8Index [constants.NumberOfSquares]int
var bishopA8H1Index [constants.NumberOfSquares]int

var bishopMoveA1H8Mask [constants.NumberOfSquares]bitboard.Board
var bishopMoveA1H8Magic [constants.NumberOfSquares]bitboard.Board
var bishopMoveA8H1Mask [constants.NumberOfSquares]bitboard.Board
var bishopMoveA8H1Magic [constants.NumberOfSquares]bitboard.Board

var bishopMoveA1H8Attacks [constants.NumberOfSquares][constants.NumberOfSquares]bitboard.Board
var bishopMoveA8H1Attacks [constants.NumberOfSquares][constants.NumberOfSquares]bitboard.Board

func init() {
	//for all fields
	for i := 0; i < constants.NumberOfSquares; i++ {
		//compute index of diagonal for the field
		bishopA8H1Index[i] = index.Index(i).FileIndex() + index.Index(i).RankIndex()%8
		bishopA1H8Index[i] = index.Index(i).FileIndex() + 7 - index.Index(i).RankIndex()%8

		//compute 6-bit diagonal for the field
		bishopMoveA8H1Mask[i] = bitboard.A8H1[bishopA8H1Index[i]] & ^bitboard.Frame
		bishopMoveA1H8Mask[i] = bitboard.A1H8[bishopA1H8Index[i]] & ^bitboard.Frame

		//index magic multiplier for the field
		bishopMoveA8H1Magic[i] = bishopMagicA8H1[bishopA8H1Index[i]]
		bishopMoveA1H8Magic[i] = bishopMagicA1H8[bishopA1H8Index[i]]
	}

	//precompute A1H8 moves
	// i is field index
	// n is 6 bit configuration

	//for all fields
	for i := 0; i < constants.NumberOfSquares; i++ {
		//for all possible diagonal states
		for n := 0; n < constants.NumberOfSquares; n++ {

			//get the diagonal
			diagonal := bitboard.A1H8[bishopA1H8Index[i]]

			//reconstruct the state (number) into the diagonal

			//get the left/bottom bit - start of diagonal
			for diagonal.OneSouthWest() != bitboard.Empty {
				diagonal = diagonal.OneSouthWest()
			}

			board := bitboard.Empty

			m := n
			//traverse diagonal and set bits according to N
			for diagonal != bitboard.Empty {
				//shift up by one
				diagonal = diagonal.OneNorthEast()
				if (m & 1) != 0 {
					board |= diagonal
				}
				m >>= 1
			}

			//make it 6-bit only
			board &= ^bitboard.Frame

			//compute possible moves
			moves := bitboard.Empty

			//set piece to Ith index
			piece := bitboard.FromIndex1(index.Index(i))

			//move in one direction
			for piece != bitboard.Empty {
				piece = piece.OneNorthEast()
				moves |= piece

				//end when there is another piece (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					piece = bitboard.Empty
				}
			}

			//set piece back to Ith index
			piece = bitboard.FromIndex1(index.Index(i))

			//move in the other direction
			for piece != bitboard.Empty {
				piece = piece.OneSouthWest()
				moves |= piece

				//end when there is another piece (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					piece = bitboard.Empty
				}
			}

			//remember the moves for Ith field within Nth occupancy state
			bishopMoveA1H8Attacks[i][n] = moves
		}
	}

	//precompute A8H1 moves
	// i is field index
	// n is 6 bit configuration
	//for all fields
	for i := 0; i < constants.NumberOfSquares; i++ {
		//for all possible diagonal states
		for n := 0; n < constants.NumberOfSquares; n++ {

			//get the diagonal
			diagonal := bitboard.A8H1[bishopA8H1Index[i]]

			//get the left/top bit - start of the diagonal
			for diagonal.OneNorthWest() != bitboard.Empty {
				diagonal = diagonal.OneNorthWest()
			}

			//traverse diagonal and set bits according to N
			board := bitboard.Empty

			m := n
			for diagonal != bitboard.Empty {
				//shift down by one
				diagonal = diagonal.OneSouthEast()
				if (m & 1) != 0 {
					board |= diagonal
				}
				m >>= 1
			}

			//make it 6-bit only
			board &= ^bitboard.Frame

			//pre-compute moves
			moves := bitboard.Empty

			//set the piece to Ith position
			piece := bitboard.FromIndex1(index.Index(i))

			//move one direction
			for piece != bitboard.Empty {
				piece = piece.OneNorthWest()
				moves |= piece
				//end when there is another piece (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					piece = 0
				}
			}

			//set the piece back to Ith position
			piece = bitboard.FromIndex1(index.Index(i))

			//move the other direction
			for piece != bitboard.Empty {
				piece = piece.OneSouthEast()
				moves |= piece
				//end when there is another piece (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					piece = 0
				}
			}
			bishopMoveA8H1Attacks[i][n] = moves
		}
	}
}

func oneBishopAttacks(sourceIndex int, allPieces bitboard.Board) bitboard.Board {
	stateIndexA8H1 := ((allPieces & bishopMoveA8H1Mask[sourceIndex]) * bishopMoveA8H1Magic[sourceIndex]) >> 57
	stateIndexA1H8 := ((allPieces & bishopMoveA1H8Mask[sourceIndex]) * bishopMoveA1H8Magic[sourceIndex]) >> 57

	//add attacks
	return bishopMoveA8H1Attacks[sourceIndex][stateIndexA8H1] | bishopMoveA1H8Attacks[sourceIndex][stateIndexA1H8]
}

func BishopAttacks(board *chessboard.Board, color chessboard.Color) bitboard.Board {
	pieces := board.Pieces[color][chessboard.Bishop] | board.Pieces[color][chessboard.Queen]
	attacks := bitboard.Empty

	//for all rooks
	for pieces != bitboard.Empty {
		attacks |= oneBishopAttacks(pieces.BitPop(), board.BoardOfAllPieces()) //TODO: try to cache BoardOfAllPieces
	}
	return attacks
}

func BishopMoves(board *chessboard.Board, moves *[]Move) {
	movingPiece := chessboard.Bishop
	bishop := board.Pieces[board.NextMove][movingPiece]

	for i := 0; i < 2; i++ { // bishops and queens

		//for all rooks
		for bishop != bitboard.Empty {
			//get next rook
			fromIndex := bishop.BitPop()

			movesBoard := oneBishopAttacks(fromIndex, board.BoardOfAllPieces()) & board.BoardAvailableToAttack() //TODO: try to cache BoardOfAllPieces and BoardAvailableToAttack

			//for all moves
			for movesBoard != bitboard.Empty {
				toIndex := movesBoard.BitPop()
				*moves = append(*moves, Move{Piece: movingPiece, From: index.Index(fromIndex), To: index.Index(toIndex)})
			}
		}
		//switch to queen
		movingPiece = chessboard.Queen
		bishop = board.Pieces[board.NextMove][movingPiece]
	}
}
