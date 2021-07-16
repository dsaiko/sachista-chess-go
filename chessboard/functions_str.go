package chessboard

import (
	"bytes"
	"regexp"
	"saiko.cz/sachista/bitboard"
	"saiko.cz/sachista/constants"
	"strconv"
)

func (b *Board) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(bitboard.BoardHeader)

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

	// print all 64 Pieces
	for i := 0; i < constants.NumberOfSquares; i++ {
		if (i % 8) == 0 {
			if i > 0 {
				buffer.WriteString(strconv.Itoa(9 - (i / 8)))
				buffer.WriteString("\n")
			}

			buffer.WriteString(strconv.Itoa(8 - (i / 8)))
			buffer.WriteString(" ")
		}

		c := byte('-')
		test := bitboard.Board(1 << i)

		switch {
		case whiteKing&test != 0:
			c = King.Notation(White)
		case whiteQueen&test != 0:
			c = Queen.Notation(White)
		case whiteRook&test != 0:
			c = Rook.Notation(White)
		case whiteKnight&test != 0:
			c = Knight.Notation(White)
		case whiteBishop&test != 0:
			c = Bishop.Notation(White)
		case whitePawn&test != 0:
			c = Pawn.Notation(White)
		case blackKing&test != 0:
			c = King.Notation(Black)
		case blackQueen&test != 0:
			c = Queen.Notation(Black)
		case blackRook&test != 0:
			c = Rook.Notation(Black)
		case blackKnight&test != 0:
			c = Knight.Notation(Black)
		case blackBishop&test != 0:
			c = Bishop.Notation(Black)
		case blackPawn&test != 0:
			c = Pawn.Notation(Black)
		}

		buffer.WriteByte(c)
		buffer.WriteString(" ")
	}

	buffer.WriteString("1\n")
	buffer.WriteString(bitboard.BoardHeader)

	return buffer.String()
}

func FromString(str string) *Board {
	fen := ""
	reHeader := regexp.MustCompile("a b c d e f g h")
	str = reHeader.ReplaceAllString(str, "")

	//create FEN string from board pieces
	for _, c := range str {
		piece, _ := PieceFromNotation(byte(c))
		if piece != NoPiece {
			fen += string(c)
		}
		if c == '-' {
			fen += "1"
		}
	}

	if len(fen) < 64 {
		fen += "/"
	}
	fen += " w KQkq - 0 1"

	return FromFen(fen)
}
