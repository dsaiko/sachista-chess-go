package generator

import (
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/index"
)

func (m *Move) MakeMove(board chessboard.Board) chessboard.Board {
	sourceIndex := m.From
	targetIndex := m.To

	if sourceIndex == targetIndex {
		return board
	}

	sourceBitBoard := bitboard.FromIndex1(sourceIndex)
	targetBitBoard := bitboard.FromIndex1(targetIndex)
	opponentColor := board.OpponentColor()

	board.HalfMoveClock++

	// reset enPassant
	if board.EnPassantTarget != 0 {
		board.ZobristHash ^= chessboard.ZobristKeys.EnPassant[board.EnPassantTarget]
		board.EnPassantTarget = 0
	}

	// reset castling
	if board.Castling[chessboard.White] != 0 {
		board.ZobristHash ^= chessboard.ZobristKeys.Castling[chessboard.White][board.Castling[chessboard.White]]
	}

	if board.Castling[chessboard.Black] != 0 {
		board.ZobristHash ^= chessboard.ZobristKeys.Castling[chessboard.Black][board.Castling[chessboard.Black]]
	}

	// make the move
	board.Pieces[board.NextMove][m.Piece] ^= sourceBitBoard | targetBitBoard
	board.ZobristHash ^= chessboard.ZobristKeys.Pieces[board.NextMove][m.Piece][sourceIndex] ^ chessboard.ZobristKeys.Pieces[board.NextMove][m.Piece][targetIndex]

	switch {
	case m.Piece == chessboard.Rook:
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
	case m.Piece == chessboard.King:
		board.Castling[board.NextMove] = chessboard.CastlingNone
		if board.NextMove == chessboard.White {
			if sourceIndex == index.E1 {
				// handle castling
				if targetIndex == index.C1 {
					board.Pieces[board.NextMove][chessboard.Rook] ^= bitboard.A1 | bitboard.D1
					board.ZobristHash ^= chessboard.ZobristKeys.Pieces[board.NextMove][chessboard.Rook][index.A1] ^ chessboard.ZobristKeys.Pieces[board.NextMove][chessboard.Rook][index.D1]
				} else if targetIndex == index.G1 {
					board.Pieces[board.NextMove][chessboard.Rook] ^= bitboard.H1 | bitboard.F1
					board.ZobristHash ^= chessboard.ZobristKeys.Pieces[board.NextMove][chessboard.Rook][index.H1] ^ chessboard.ZobristKeys.Pieces[board.NextMove][chessboard.Rook][index.F1]
				}
			}
		} else {
			if sourceIndex == index.E8 {
				// handle castling
				if targetIndex == index.C8 {
					board.Pieces[board.NextMove][chessboard.Rook] ^= bitboard.A8 | bitboard.D8
					board.ZobristHash ^= chessboard.ZobristKeys.Pieces[board.NextMove][chessboard.Rook][index.A8] ^ chessboard.ZobristKeys.Pieces[board.NextMove][chessboard.Rook][index.D8]
				} else if targetIndex == index.G8 {
					board.Pieces[board.NextMove][chessboard.Rook] ^= bitboard.H8 | bitboard.F8
					board.ZobristHash ^= chessboard.ZobristKeys.Pieces[board.NextMove][chessboard.Rook][index.H8] ^ chessboard.ZobristKeys.Pieces[board.NextMove][chessboard.Rook][index.F8]
				}
			}
		}
	case m.Piece == chessboard.Pawn:
		board.HalfMoveClock = 0
		if absInt(int(targetIndex)-int(sourceIndex)) > 10 {
			var n index.Index = 8
			if board.NextMove == chessboard.Black {
				n = -8
			}
			board.EnPassantTarget = sourceIndex + n
		} else if m.PromotionPiece > 0 {
			board.Pieces[board.NextMove][chessboard.Pawn] ^= targetBitBoard
			board.ZobristHash ^= chessboard.ZobristKeys.Pieces[board.NextMove][chessboard.Pawn][targetIndex]

			board.Pieces[board.NextMove][m.PromotionPiece] ^= targetBitBoard
			board.ZobristHash ^= chessboard.ZobristKeys.Pieces[board.NextMove][m.PromotionPiece][targetIndex]
		}
	}

	isCapture := targetBitBoard&board.BoardOfOpponentPieces() != 0
	if isCapture || m.IsEnPassant {
		// check capture
		board.HalfMoveClock = 0

		checkCapture := func(piece chessboard.Piece) bool {
			if board.Pieces[opponentColor][piece]&targetBitBoard != 0 {
				board.Pieces[opponentColor][piece] ^= targetBitBoard
				board.ZobristHash ^= chessboard.ZobristKeys.Pieces[opponentColor][piece][targetIndex]
				return true
			}
			return false
		}

		switch {
		case m.IsEnPassant:
			if board.NextMove == chessboard.White {
				board.Pieces[chessboard.Black][chessboard.Pawn] ^= targetBitBoard.OneSouth()
				board.ZobristHash ^= chessboard.ZobristKeys.Pieces[chessboard.Black][chessboard.Pawn][targetIndex-8]
			} else {
				board.Pieces[chessboard.White][chessboard.Pawn] ^= targetBitBoard.OneNorth()
				board.ZobristHash ^= chessboard.ZobristKeys.Pieces[chessboard.White][chessboard.Pawn][targetIndex+8]
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
		board.FullMoveNumber++
	}

	// update Zobrist
	board.NextMove = opponentColor
	board.ZobristHash ^= chessboard.ZobristKeys.Side

	if board.Castling[chessboard.White] != 0 {
		board.ZobristHash ^= chessboard.ZobristKeys.Castling[chessboard.White][board.Castling[chessboard.White]]
	}

	if board.Castling[chessboard.Black] != 0 {
		board.ZobristHash ^= chessboard.ZobristKeys.Castling[chessboard.Black][board.Castling[chessboard.Black]]
	}

	if board.EnPassantTarget != 0 {
		board.ZobristHash ^= chessboard.ZobristKeys.EnPassant[board.EnPassantTarget]
	}

	return board
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
