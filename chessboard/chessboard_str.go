package chessboard

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"

	"saiko.cz/sachista/bitboard"
)

// String returns an ASCII representation of the board with piece letters and rank/file labels.
// Note: castling rights and en passant state are not included in the output.
func (b *Board) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(bitboard.BoardHeader)

	whiteKing := b.Pieces[White][King].MirroredVertical()
	whiteQueen := b.Pieces[White][Queen].MirroredVertical()
	whiteRook := b.Pieces[White][Rook].MirroredVertical()
	whiteKnight := b.Pieces[White][Knight].MirroredVertical()
	whiteBishop := b.Pieces[White][Bishop].MirroredVertical()
	whitePawn := b.Pieces[White][Pawn].MirroredVertical()
	blackKing := b.Pieces[Black][King].MirroredVertical()
	blackQueen := b.Pieces[Black][Queen].MirroredVertical()
	blackRook := b.Pieces[Black][Rook].MirroredVertical()
	blackKnight := b.Pieces[Black][Knight].MirroredVertical()
	blackBishop := b.Pieces[Black][Bishop].MirroredVertical()
	blackPawn := b.Pieces[Black][Pawn].MirroredVertical()

	// print all 64 Pieces
	for i := range bitboard.NumberOfSquares {
		if (i % 8) == 0 {
			if i > 0 {
				buffer.WriteString(strconv.Itoa(9 - (i / 8)))
				buffer.WriteString("\n")
			}

			buffer.WriteString(strconv.Itoa(8 - (i / 8)))
			buffer.WriteString(" ")
		}

		c := "-"
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

		buffer.WriteString(c)
		buffer.WriteString(" ")
	}

	buffer.WriteString("1\n")
	buffer.WriteString(bitboard.BoardHeader)

	return buffer.String()
}

// FromString parses an ASCII board representation (as produced by String) back into a Board.
// Default values are used for castling rights, en passant, and move counters.
func FromString(str string) Board {
	var fenBuilder strings.Builder
	reHeader := regexp.MustCompile("a b c d e f g h")
	str = reHeader.ReplaceAllString(str, "")

	// create FEN string from board pieces
	for _, c := range str {
		piece, _ := PieceFromNotation(string(c))
		if piece != NoPiece {
			fenBuilder.WriteRune(c)
		}
		if c == '-' {
			fenBuilder.WriteString("1")
		}
	}

	fen := fenBuilder.String()
	if len(fen) < 64 {
		fen += "/"
	}
	fen += " w KQkq - 0 1"

	return BoardFromFEN(fen)
}
