package chessboard

import (
	"bytes"
	"strconv"

	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/constants"
	"saiko.cz/sachista/index"
)

// ToFEN converts board to FEN notation string
func (b Board) ToFEN() string {

	var buffer bytes.Buffer

	whiteKing := b.Pieces[White][King].MirrorVertical()
	whiteQueen := b.Pieces[White][Queen].MirrorVertical()
	whiteRook := b.Pieces[White][Rook].MirrorVertical()
	whiteKnight := b.Pieces[White][Knight].MirrorVertical()
	whiteBishop := b.Pieces[White][Bishop].MirrorVertical()
	whitePawn := b.Pieces[White][Pawn].MirrorVertical()
	blackKing := b.Pieces[Black][King].MirrorVertical()
	blackQueen := b.Pieces[Black][Queen].MirrorVertical()
	blackRook := b.Pieces[Black][Rook].MirrorVertical()
	blackKnight := b.Pieces[Black][Knight].MirrorVertical()
	blackBishop := b.Pieces[Black][Bishop].MirrorVertical()
	blackPawn := b.Pieces[Black][Pawn].MirrorVertical()

	spaces := 0

	outputCount := func() {
		if spaces > 0 {
			buffer.WriteString(strconv.Itoa(spaces))
			spaces = 0
		}
	}

	for i := 0; i < constants.NumberOfSquares; i++ {
		if (i % 8) == 0 {
			outputCount()
			if i > 0 {
				buffer.WriteString("/")
			}
		}

		c := ""
		test := bitboard.Board(1 << i)

		switch {
		case whiteKing&test != 0:
			c = King.String(White)
		case whiteQueen&test != 0:
			c = Queen.String(White)
		case whiteRook&test != 0:
			c = Rook.String(White)
		case whiteKnight&test != 0:
			c = Knight.String(White)
		case whiteBishop&test != 0:
			c = Bishop.String(White)
		case whitePawn&test != 0:
			c = Pawn.String(White)
		case blackKing&test != 0:
			c = King.String(Black)
		case blackQueen&test != 0:
			c = Queen.String(Black)
		case blackRook&test != 0:
			c = Rook.String(Black)
		case blackKnight&test != 0:
			c = Knight.String(Black)
		case blackBishop&test != 0:
			c = Bishop.String(Black)
		case blackPawn&test != 0:
			c = Pawn.String(Black)
		}

		if len(c) != 0 {
			outputCount()
			buffer.WriteString(c)
		} else {
			spaces++
		}
	}
	outputCount()

	// next move color
	buffer.WriteString(" ")
	buffer.WriteString(b.NextMove.String())
	buffer.WriteString(" ")

	// Castling
	if b.Castling[White]&CastlingKingSide != 0 {
		buffer.WriteString(King.String(White))
	}
	if b.Castling[White]&CastlingQueenSide != 0 {
		buffer.WriteString(Queen.String(White))
	}
	if b.Castling[Black]&CastlingKingSide != 0 {
		buffer.WriteString(King.String(Black))
	}
	if b.Castling[Black]&CastlingQueenSide != 0 {
		buffer.WriteString(Queen.String(Black))
	}
	if b.Castling[White]|b.Castling[Black] == 0 {
		buffer.WriteString("-")
	}

	// enPassant
	buffer.WriteString(" ")
	if b.EnPassantTarget != 0 {
		buffer.WriteString(b.EnPassantTarget.String())
	} else {
		buffer.WriteString("-")
	}

	// move number
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(b.HalfMoveClock))
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(b.FullMoveNumber))

	return buffer.String()
}

// FromFEN converts FEN string to board object
func FromFEN(fen string) Board {
	b := Empty()
	i := 0
	fenLength := len(fen)

	for ; i < fenLength; i++ {
		c := fen[i]
		if c == ' ' {
			break
		}

		if c == '/' {
			//nothing
			continue
		}

		if c >= '0' && c <= '9' {
			n := c - '0'

			// output number of empty fields
			for color := White; color <= Black; color++ {
				for piece := King; piece <= Pawn; piece++ {
					b.Pieces[color][piece] <<= n
				}
			}
		} else {
			// output a piece
			b.Pieces[White][King] <<= 1
			b.Pieces[White][Queen] <<= 1
			b.Pieces[White][Rook] <<= 1
			b.Pieces[White][Knight] <<= 1
			b.Pieces[White][Bishop] <<= 1
			b.Pieces[White][Pawn] <<= 1

			b.Pieces[Black][King] <<= 1
			b.Pieces[Black][Queen] <<= 1
			b.Pieces[Black][Rook] <<= 1
			b.Pieces[Black][Knight] <<= 1
			b.Pieces[Black][Bishop] <<= 1
			b.Pieces[Black][Pawn] <<= 1

			// set the new piece
			piece, color := PieceFromNotation(string(c))
			if piece != NoPiece {
				b.Pieces[color][piece] |= 1
			}
		}
	}

	// need to mirror the boards
	for color := White; color <= Black; color++ {
		for piece := King; piece <= Pawn; piece++ {
			b.Pieces[color][piece] = b.Pieces[color][piece].MirrorHorizontal()
		}
	}

	// next move
	i++ //skip space
	if i < fenLength {
		if fen[i] == 'w' {
			b.NextMove = White
		} else {
			b.NextMove = Black
		}
		i++
	}

	// castling
	i++ // skip space
	for ; i < fenLength; i++ {
		c := fen[i]
		if c == ' ' {
			break
		}

		piece, color := PieceFromNotation(string(c))

		castling := CastlingQueenSide
		if piece == King {
			castling = CastlingKingSide
		}

		b.Castling[color] |= castling
	}

	// enPassant
	i++ // skip space
	notation := ""
	for ; i < fenLength; i++ {
		c := fen[i]
		if c == ' ' {
			break
		}

		if c != '-' && len(notation) < 2 {
			notation += string(c)
		}

		if len(notation) == 2 {
			b.EnPassantTarget = index.FromNotation(notation)
		}
	}

	// half move clock
	i++ // skip space
	n := 0
	for ; i < fenLength; i++ {
		c := fen[i]
		if c == ' ' {
			break
		}

		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	if n != 0 {
		b.HalfMoveClock = n
	}

	i++ // skip space
	n = 0
	for ; i < fenLength; i++ {
		c := fen[i]
		if c == ' ' {
			break
		}

		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	if n != 0 {
		b.FullMoveNumber = n
	}

	// fix castling
	if (b.Pieces[White][Rook] & bitboard.A1) == 0 {
		b.RemoveCastling(White, CastlingQueenSide)
	}
	if (b.Pieces[White][Rook] & bitboard.H1) == 0 {
		b.RemoveCastling(White, CastlingKingSide)
	}
	if (b.Pieces[Black][Rook] & bitboard.A8) == 0 {
		b.RemoveCastling(Black, CastlingQueenSide)
	}
	if (b.Pieces[Black][Rook] & bitboard.H8) == 0 {
		b.RemoveCastling(Black, CastlingKingSide)
	}

	// if king is misplaced, remove castling availability
	if (b.Pieces[White][King] & bitboard.E1) == 0 {
		b.Castling[White] = CastlingNone
	}
	if (b.Pieces[Black][King] & bitboard.E8) == 0 {
		b.Castling[Black] = CastlingNone
	}

	b.ZobristHash = b.ComputeBoardHash()
	return b
}
