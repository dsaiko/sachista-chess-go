package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
)

var rookMagicFile = [...]bitboard.Board{
	0x8040201008040200,
	0x4020100804020100,
	0x2010080402010080,
	0x1008040201008040,
	0x0804020100804020,
	0x0402010080402010,
	0x0201008040201008,
	0x0100804020100804,
}

var rookMoveRankShift [constants.NumberOfSquares]int
var rookMoveRankMask [constants.NumberOfSquares]bitboard.Board
var rookMoveRankAttacks [constants.NumberOfSquares][constants.NumberOfSquares]bitboard.Board
var rookMoveFileMask [constants.NumberOfSquares]bitboard.Board
var rookMoveFileMagic [constants.NumberOfSquares]bitboard.Board
var rookMoveFileAttacks [constants.NumberOfSquares][constants.NumberOfSquares]bitboard.Board

func init() {
	const fileAMask = bitboard.A2 | bitboard.A3 | bitboard.A4 | bitboard.A5 | bitboard.A6 | bitboard.A7

	for i := 0; i < constants.NumberOfSquares; i++ {

		fieldIndex := index.Index(i)

		//get 6-bit mask for a rank
		rookMoveRankMask[i] = bitboard.Board(126) << (fieldIndex.RankIndex() << 3)

		//compute needed rank shift
		rookMoveRankShift[i] = (fieldIndex.RankIndex() << 3) + 1

		//get 6-bit mask for a file
		rookMoveFileMask[i] = fileAMask << fieldIndex.FileIndex()

		//index magic number directly fo field
		rookMoveFileMagic[i] = rookMagicFile[fieldIndex.FileIndex()]
	}

	//precompute rank moves
	//for all pieces
	for i := 0; i < constants.NumberOfSquares; i++ {
		rankIndex := index.Index(i).RankIndex()

		//for all occupancy states
		for n := 0; n < constants.NumberOfSquares; n++ {
			//reconstruct occupancy state
			board := bitboard.Board(n).Shift(1, rankIndex)

			//generate available moves
			moves := bitboard.Empty

			//set piece in Ith position
			piece := bitboard.FromIndex1(index.Index(i))

			//move in one direction
			for piece != bitboard.Empty {
				piece = piece.OneWest()

				moves |= piece

				//end when there is another piece on the board (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					break
				}
			}

			//set piece back in Ith position
			piece = bitboard.FromIndex1(index.Index(i))

			//move in other direction
			for piece != bitboard.Empty {
				piece = piece.OneEast()
				moves |= piece

				//end when there is another piece on the board (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					break
				}
			}

			//remember the moves
			rookMoveRankAttacks[i][n] = moves
		}
	}

	//precompute file moves
	//for all pieces
	for i := 0; i < constants.NumberOfSquares; i++ {
		fileIndex := index.Index(i).FileIndex()

		//for all occupancy states
		for n := 0; n < constants.NumberOfSquares; n++ {

			//reconstuct the occupancy into file
			board := bitboard.Board(n).Shift(1, 0).MirrorHorizontal().FlipA1H8().Shift(fileIndex, 0)

			//generate available moves
			moves := bitboard.Empty

			//set piece back in Ith position
			piece := bitboard.FromIndex1(index.Index(i))

			//move piece in one direction
			for piece != bitboard.Empty {
				piece = piece.OneNorth()

				moves |= piece

				//end when there is another piece on the board (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					break
				}
			}

			//set piece back to original Ith index
			piece = bitboard.FromIndex1(index.Index(i))

			//move piece in other direction
			for piece != bitboard.Empty {
				piece = piece.OneSouth()

				moves |= piece

				//end when there is another piece on the board (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					break
				}
			}

			//remember file attacks
			rookMoveFileAttacks[i][n] = moves
		}
	}
}

func oneRookAttacks(sourceIndex int, allPieces bitboard.Board) bitboard.Board {
	//use magic multipliers to get occupancy state index
	stateIndexRank := (allPieces & rookMoveRankMask[sourceIndex]) >> rookMoveRankShift[sourceIndex]
	stateIndexFile := ((allPieces & rookMoveFileMask[sourceIndex]) * rookMoveFileMagic[sourceIndex]) >> 57

	//get possible attacks for field / occupancy state index
	return rookMoveRankAttacks[sourceIndex][stateIndexRank] | rookMoveFileAttacks[sourceIndex][stateIndexFile]
}

func RookAttacks(board *chessboard.Board, color chessboard.Color) bitboard.Board {
	pieces := board.Pieces[color][chessboard.Rook] | board.Pieces[color][chessboard.Queen]
	attacks := bitboard.Empty

	//for all rooks
	for pieces != bitboard.Empty {
		attacks |= oneRookAttacks(pieces.BitPop(), board.AllPieces()) //TODO: try to cache AllPieces
	}
	return attacks
}

func RookMoves(board *chessboard.Board, moves *[]Move) {
	movingPiece := chessboard.Rook
	rook := board.Pieces[board.NextMove][movingPiece]

	for i := 0; i < 2; i++ { // rooks and queens

		//for all rooks
		for rook != bitboard.Empty {
			//get next rook
			fromIndex := rook.BitPop()

			movesBoard := oneRookAttacks(fromIndex, board.AllPieces()) & board.BoardAvailable() //TODO: try to cache AllPieces and BoardAvailable

			//for all moves
			for movesBoard != bitboard.Empty {
				toIndex := movesBoard.BitPop()
				*moves = append(*moves, Move{Piece: movingPiece, From: index.Index(fromIndex), To: index.Index(toIndex)})
			}
		}
		//switch to queen
		movingPiece := chessboard.Queen
		rook = board.Pieces[board.NextMove][movingPiece]
	}
}
