package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/index"
)

//TODO: test return new board vs apply to existing
//TODO: try undo move
//TODO: try recompute zobrist vs modify zobrist
//TODO: test pointers
func (m *Move) MakeMove(board chessboard.Board) *chessboard.Board {

	sourceIndex := m.From
	targetIndex := m.To

	//TODO: test perf if removed
	if sourceIndex == targetIndex {
		return &board
	}

	sourceBitBoard := bitboard.FromIndex1(sourceIndex)
	targetBitBoard := bitboard.FromIndex1(targetIndex)
	opponentColor := board.OpponentColor()

	board.HalfMoveClock += 1
	board.EnPassantTarget = 0

	board.Pieces[board.NextMove][m.Piece] ^= sourceBitBoard | targetBitBoard

	if m.Piece == chessboard.Rook {
		if board.NextMove == chessboard.White {
			if sourceIndex == index.A1 {
				board.Castling[board.NextMove] &= ^chessboard.CastlingQueenSide
			} else if sourceIndex == index.H1 {
				board.Castling[board.NextMove] &= ^chessboard.CastlingKingSide
			}
		} else {
			if sourceIndex == index.A8 {
				board.RemoveCastling(board.NextMove, chessboard.CastlingQueenSide)
			} else if sourceIndex == index.H8 {
				board.RemoveCastling(board.NextMove, chessboard.CastlingKingSide)
			}
		}
	} else if m.Piece == chessboard.King {
		board.Castling[board.NextMove] = chessboard.CastlingNone

		if board.NextMove == chessboard.White {
			if sourceIndex == index.E1 {
				//handle castling
				if targetIndex == index.C1 {
					board.Pieces[board.NextMove][chessboard.Rook] ^= bitboard.A1 | bitboard.D1
				} else if targetIndex == index.G1 {
					board.Pieces[board.NextMove][chessboard.Rook] ^= bitboard.H1 | bitboard.F1
				}
			}
		} else {
			if sourceIndex == index.E8 {
				//handle castling
				if targetIndex == index.C8 {
					board.Pieces[board.NextMove][chessboard.Rook] ^= bitboard.A8 | bitboard.D8
				} else if targetIndex == index.G8 {
					board.Pieces[board.NextMove][chessboard.Rook] ^= bitboard.H8 | bitboard.F8
				}
			}
		}
	} else if m.Piece == chessboard.Pawn {
		board.HalfMoveClock = 0

		//initial pawn double move
		if absInt(int(targetIndex)-int(sourceIndex)) > 10 {
			var n index.Index = 8
			if board.NextMove == chessboard.Black {
				n = -8
			}
			board.EnPassantTarget = sourceIndex + n
		} else if m.PromotionPiece > 0 {
			board.Pieces[board.NextMove][chessboard.Pawn] ^= targetBitBoard
			board.Pieces[board.NextMove][m.PromotionPiece] ^= targetBitBoard
		}
	}

	isCapture := targetBitBoard&board.OpponentPieces() != 0
	if isCapture || m.IsEnPassant {
		board.HalfMoveClock = 0

		checkCapture := func(piece chessboard.Piece) bool {
			if board.Pieces[opponentColor][piece]&targetBitBoard != 0 {
				board.Pieces[opponentColor][piece] ^= targetBitBoard
				return true
			}
			return false
		}

		switch {
		case m.IsEnPassant:
			if board.NextMove == chessboard.White {
				board.Pieces[chessboard.Black][chessboard.Pawn] ^= targetBitBoard.OneSouth()
			} else {
				board.Pieces[chessboard.White][chessboard.Pawn] ^= targetBitBoard.OneNorth()
			}
		case checkCapture(chessboard.Bishop):
		case checkCapture(chessboard.Knight):
		case checkCapture(chessboard.Pawn):
		case checkCapture(chessboard.Queen):
		case checkCapture(chessboard.Rook):
			if board.NextMove == chessboard.White {
				if targetIndex == index.A8 {
					board.RemoveCastling(chessboard.Black, chessboard.CastlingQueenSide)
				} else if targetIndex == index.H8 {
					board.RemoveCastling(chessboard.Black, chessboard.CastlingKingSide)
				}
			} else {
				if targetIndex == index.A1 {
					board.RemoveCastling(chessboard.White, chessboard.CastlingQueenSide)
				} else if targetIndex == index.H1 {
					board.RemoveCastling(chessboard.White, chessboard.CastlingKingSide)
				}
			}
		}
	}

	if board.NextMove == chessboard.Black {
		board.FullMoveNumber += 1
	}

	board.NextMove = opponentColor

	board.UpdateZobrist()
	return &board
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
