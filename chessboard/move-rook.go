package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

// rookMagicFile holds magic multiplier constants for each file,
// used to map file occupancy states to a 6-bit lookup index.
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

var (
	rookMoveRankShift   [bitboard.NumberOfSquares]int
	rookMoveRankMask    [bitboard.NumberOfSquares]bitboard.Board
	rookMoveRankAttacks [bitboard.NumberOfSquares][bitboard.NumberOfSquares]bitboard.Board
	rookMoveFileMask    [bitboard.NumberOfSquares]bitboard.Board
	rookMoveFileMagic   [bitboard.NumberOfSquares]bitboard.Board
	rookMoveFileAttacks [bitboard.NumberOfSquares][bitboard.NumberOfSquares]bitboard.Board
)

func init() {
	const fileAMask = bitboard.BoardA2 | bitboard.BoardA3 | bitboard.BoardA4 | bitboard.BoardA5 | bitboard.BoardA6 | bitboard.BoardA7

	for i := range bitboard.NumberOfSquares {
		fieldIndex := bitboard.Index(i)

		// get 6-bit mask for a rank
		rookMoveRankMask[i] = bitboard.Board(126) << (fieldIndex.Rank() << 3) //nolint:revive

		// compute needed rank shift
		rookMoveRankShift[i] = (fieldIndex.Rank() << 3) + 1

		// get 6-bit mask for a file
		rookMoveFileMask[i] = fileAMask << fieldIndex.File()

		// index magic number directly for the field
		rookMoveFileMagic[i] = rookMagicFile[fieldIndex.File()]
	}

	// precompute rank moves
	// for all pieces
	for i := range bitboard.NumberOfSquares {
		rankIndex := bitboard.Index(i).Rank()

		// for all occupancy states
		for n := range bitboard.NumberOfSquares {
			// reconstruct occupancy state
			board := bitboard.Board(n).Shifted(1, rankIndex)

			// generate available moves
			moves := bitboard.EmptyBoard

			// set piece in Ith position
			piece := bitboard.BoardFromIndex(bitboard.Index(i))

			// move in one direction
			for piece != bitboard.EmptyBoard {
				piece = piece.ShiftedOneWest()

				moves |= piece

				// end when there is another piece on the board (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					break
				}
			}

			// set piece back in Ith position
			piece = bitboard.BoardFromIndex(bitboard.Index(i))

			// move in other direction
			for piece != bitboard.EmptyBoard {
				piece = piece.ShiftedOneEast()
				moves |= piece

				// end when there is another piece on the board (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					break
				}
			}

			// remember the moves
			rookMoveRankAttacks[i][n] = moves
		}
	}

	// precompute file moves
	// for all pieces
	for i := range bitboard.NumberOfSquares {
		fileIndex := bitboard.Index(i).File()

		// for all occupancy states
		for n := range bitboard.NumberOfSquares {
			// reconstruct the occupancy into file
			board := bitboard.Board(n).Shifted(1, 0).MirroredHorizontal().FlippedA1H8().Shifted(fileIndex, 0)

			// generate available moves
			moves := bitboard.EmptyBoard

			// set piece back in Ith position
			piece := bitboard.BoardFromIndex(bitboard.Index(i))

			// move piece in one direction
			for piece != bitboard.EmptyBoard {
				piece = piece.ShiftedOneNorth()

				moves |= piece

				// end when there is another piece on the board (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					break
				}
			}

			// set piece back to original Ith index
			piece = bitboard.BoardFromIndex(bitboard.Index(i))

			// move piece in other direction
			for piece != bitboard.EmptyBoard {
				piece = piece.ShiftedOneSouth()

				moves |= piece

				// end when there is another piece on the board (either color, own color will have to be stripped out)
				if (piece & board) != 0 {
					break
				}
			}

			// remember file attacks
			rookMoveFileAttacks[i][n] = moves
		}
	}
}

// oneRookAttacks computes the attack bitboard for a single rook using magic bitboard lookups.
func oneRookAttacks(sourceIndex bitboard.Index, allPieces bitboard.Board) bitboard.Board {
	// use magic multipliers to get occupancy state index
	stateIndexRank := (allPieces & rookMoveRankMask[sourceIndex]) >> rookMoveRankShift[sourceIndex]
	stateIndexFile := ((allPieces & rookMoveFileMask[sourceIndex]) * rookMoveFileMagic[sourceIndex]) >> 57 //nolint:revive

	// get possible attacks for field / occupancy state index
	return rookMoveRankAttacks[sourceIndex][stateIndexRank] | rookMoveFileAttacks[sourceIndex][stateIndexFile]
}

// rookAttacks returns the combined attack bitboard for all rooks and queens (rank/file) of the given color.
func rookAttacks(board *Board, color Color) bitboard.Board {
	pieces := board.Pieces[color][Rook] | board.Pieces[color][Queen]
	attacks := bitboard.EmptyBoard

	allPieces := board.AllPieces()

	var i bitboard.Index
	// for all rooks
	for pieces != bitboard.EmptyBoard {
		i, pieces = pieces.BitPop()
		attacks |= oneRookAttacks(i, allPieces)
	}
	return attacks
}

// rookMoves generates all pseudo-legal rook and queen (rank/file) moves for the current side.
func rookMoves(board *Board, handler MoveHandler) {
	movingPiece := Rook
	rook := board.Pieces[board.NextMove][movingPiece]

	allPieces := board.AllPieces()
	boardAvailable := board.BoardAvailableToAttack()

	var fromIndex, toIndex bitboard.Index

	for range 2 { // rooks and queens
		// for all rooks
		for rook != bitboard.EmptyBoard {
			// get next rook
			fromIndex, rook = rook.BitPop()
			movesBoard := oneRookAttacks(fromIndex, allPieces) & boardAvailable

			// for all moves
			for movesBoard != bitboard.EmptyBoard {
				toIndex, movesBoard = movesBoard.BitPop()
				handler(Move{Piece: movingPiece, From: fromIndex, To: toIndex})
			}
		}
		// switch to queen
		movingPiece = Queen
		rook = board.Pieces[board.NextMove][movingPiece]
	}
}
