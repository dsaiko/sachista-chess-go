package chessboard

import (
	"saiko.cz/sachista/bitboard"
)

func (b *Board) AppliedMove(m Move) *Board {
	sourceIndex := m.From
	targetIndex := m.To
	board := *b

	if sourceIndex == targetIndex {
		return &board
	}

	sourceBitBoard := bitboard.BoardFromIndex(sourceIndex)
	targetBitBoard := bitboard.BoardFromIndex(targetIndex)
	opponentColor := board.OpponentColor()

	board.HalfMoveClock++

	// reset enPassant
	if board.EnPassantTarget != 0 {
		board.ZobristHash ^= ZobristKeys.EnPassant[board.EnPassantTarget]
		board.EnPassantTarget = 0
	}

	// reset castling
	if board.Castling[White] != 0 {
		board.ZobristHash ^= ZobristKeys.Castling[White][board.Castling[White]]
	}

	if board.Castling[Black] != 0 {
		board.ZobristHash ^= ZobristKeys.Castling[Black][board.Castling[Black]]
	}

	// make the move
	board.Pieces[board.NextMove][m.Piece] ^= sourceBitBoard | targetBitBoard
	board.ZobristHash ^= ZobristKeys.Pieces[board.NextMove][m.Piece][sourceIndex] ^ ZobristKeys.Pieces[board.NextMove][m.Piece][targetIndex]

	switch {
	case m.Piece == Rook:
		if board.NextMove == White {
			if sourceIndex == bitboard.IndexA1 {
				board.Castling[board.NextMove] &= ^CastlingQueenSide
			} else if sourceIndex == bitboard.IndexH1 {
				board.Castling[board.NextMove] &= ^CastlingKingSide
			}
		} else {
			if sourceIndex == bitboard.IndexA8 {
				board.RemovedCastling(board.NextMove, CastlingQueenSide)
			} else if sourceIndex == bitboard.IndexH8 {
				board.RemovedCastling(board.NextMove, CastlingKingSide)
			}
		}
	case m.Piece == King:
		board.Castling[board.NextMove] = CastlingNone
		if board.NextMove == White {
			if sourceIndex == bitboard.IndexE1 {
				// handle castling
				if targetIndex == bitboard.IndexC1 {
					board.Pieces[board.NextMove][Rook] ^= bitboard.BoardA1 | bitboard.BoardD1
					board.ZobristHash ^= ZobristKeys.Pieces[board.NextMove][Rook][bitboard.IndexA1] ^ ZobristKeys.Pieces[board.NextMove][Rook][bitboard.IndexD1]
				} else if targetIndex == bitboard.IndexG1 {
					board.Pieces[board.NextMove][Rook] ^= bitboard.BoardH1 | bitboard.BoardF1
					board.ZobristHash ^= ZobristKeys.Pieces[board.NextMove][Rook][bitboard.IndexH1] ^ ZobristKeys.Pieces[board.NextMove][Rook][bitboard.IndexF1]
				}
			}
		} else {
			if sourceIndex == bitboard.IndexE8 {
				// handle castling
				if targetIndex == bitboard.IndexC8 {
					board.Pieces[board.NextMove][Rook] ^= bitboard.BoardA8 | bitboard.BoardD8
					board.ZobristHash ^= ZobristKeys.Pieces[board.NextMove][Rook][bitboard.IndexA8] ^ ZobristKeys.Pieces[board.NextMove][Rook][bitboard.IndexD8]
				} else if targetIndex == bitboard.IndexG8 {
					board.Pieces[board.NextMove][Rook] ^= bitboard.BoardH8 | bitboard.BoardF8
					board.ZobristHash ^= ZobristKeys.Pieces[board.NextMove][Rook][bitboard.IndexH8] ^ ZobristKeys.Pieces[board.NextMove][Rook][bitboard.IndexF8]
				}
			}
		}
	case m.Piece == Pawn:
		board.HalfMoveClock = 0
		if absInt(int(targetIndex)-int(sourceIndex)) > 10 {
			var n bitboard.Index = 8
			if board.NextMove == Black {
				n = -8
			}
			board.EnPassantTarget = sourceIndex + n
		} else if m.PromotionPiece > 0 {
			board.Pieces[board.NextMove][Pawn] ^= targetBitBoard
			board.ZobristHash ^= ZobristKeys.Pieces[board.NextMove][Pawn][targetIndex]

			board.Pieces[board.NextMove][m.PromotionPiece] ^= targetBitBoard
			board.ZobristHash ^= ZobristKeys.Pieces[board.NextMove][m.PromotionPiece][targetIndex]
		}
	}

	isCapture := targetBitBoard&board.OpponentPieces() != 0
	if isCapture || m.IsEnPassant {
		// check capture
		board.HalfMoveClock = 0

		checkCapture := func(piece Piece) bool {
			if board.Pieces[opponentColor][piece]&targetBitBoard != 0 {
				board.Pieces[opponentColor][piece] ^= targetBitBoard
				board.ZobristHash ^= ZobristKeys.Pieces[opponentColor][piece][targetIndex]
				return true
			}
			return false
		}

		switch {
		case m.IsEnPassant:
			if board.NextMove == White {
				board.Pieces[Black][Pawn] ^= targetBitBoard.ShiftedOneSouth()
				board.ZobristHash ^= ZobristKeys.Pieces[Black][Pawn][targetIndex-8]
			} else {
				board.Pieces[White][Pawn] ^= targetBitBoard.ShiftedOneNorth()
				board.ZobristHash ^= ZobristKeys.Pieces[White][Pawn][targetIndex+8]
			}
		case checkCapture(Bishop):
		case checkCapture(Knight):
		case checkCapture(Pawn):
		case checkCapture(Queen):
		case checkCapture(Rook):
			if board.NextMove == White {
				if targetIndex == bitboard.IndexA8 {
					board.RemovedCastling(Black, CastlingQueenSide)
				} else if targetIndex == bitboard.IndexH8 {
					board.RemovedCastling(Black, CastlingKingSide)
				}
			} else {
				if targetIndex == bitboard.IndexA1 {
					board.RemovedCastling(White, CastlingQueenSide)
				} else if targetIndex == bitboard.IndexH1 {
					board.RemovedCastling(White, CastlingKingSide)
				}
			}
		}
	}

	if board.NextMove == Black {
		board.FullMoveNumber++
	}

	// update Zobrist
	board.NextMove = opponentColor
	board.ZobristHash ^= ZobristKeys.Side

	if board.Castling[White] != 0 {
		board.ZobristHash ^= ZobristKeys.Castling[White][board.Castling[White]]
	}

	if board.Castling[Black] != 0 {
		board.ZobristHash ^= ZobristKeys.Castling[Black][board.Castling[Black]]
	}

	if board.EnPassantTarget != 0 {
		board.ZobristHash ^= ZobristKeys.EnPassant[board.EnPassantTarget]
	}

	return &board
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
