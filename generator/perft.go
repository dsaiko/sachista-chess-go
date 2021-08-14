package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/constants"
)

// TODO: implement cache
// TODO: compare cache and non cached results
func PerfT(b *chessboard.Board, depth int) uint64 {
	moves := make([]Move, 0, constants.MovesCacheInitialCapacity)

	GeneratePseudoLegalMoves(b, &moves)

	var count uint64 = 0
	attacks := Attacks(b, b.OpponentColor())
	isCheck := attacks&b.Pieces[b.NextMove][chessboard.King] != 0

	for _, m := range moves {
		sourceBitBoard := bitboard.FromIndex1(m.From)
		isKingMove := m.Piece == chessboard.King

		if isKingMove || isCheck || (sourceBitBoard&attacks) != 0 || m.IsEnPassant {
			//need to validate move
			nextBoard := m.MakeMove(*b)
			if isOpponentsKingNotUnderCheck(nextBoard) {
				if depth == 1 {
					count += 1
				} else {
					count += PerfT(nextBoard, depth-1)
				}
			}
		} else {
			if depth == 1 {
				count += 1
			} else {
				//do not need to validate legality of the move
				nextBoard := m.MakeMove(*b)
				count += PerfT(nextBoard, depth-1)
			}
		}
	}

	// DEBUG OUTPUT :
	// fmt.Printf("%v|%v|%v\n",b.ToFEN(), depth, count)
	return count
}
