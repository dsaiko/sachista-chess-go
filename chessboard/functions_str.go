package chessboard

import (
	"bytes"
	"saiko.cz/sachista/bitboard"
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
	for i := 0; i < bitboard.BitWidth; i++ {
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
			c = King.Description(White)
		case whiteQueen&test != 0:
			c = Queen.Description(White)
		case whiteRook&test != 0:
			c = Rook.Description(White)
		case whiteKnight&test != 0:
			c = Knight.Description(White)
		case whiteBishop&test != 0:
			c = Bishop.Description(White)
		case whitePawn&test != 0:
			c = Pawn.Description(White)
		case blackKing&test != 0:
			c = King.Description(Black)
		case blackQueen&test != 0:
			c = Queen.Description(Black)
		case blackRook&test != 0:
			c = Rook.Description(Black)
		case blackKnight&test != 0:
			c = Knight.Description(Black)
		case blackBishop&test != 0:
			c = Bishop.Description(Black)
		case blackPawn&test != 0:
			c = Pawn.Description(Black)
		}

		buffer.WriteByte(c)
		buffer.WriteString(" ")
	}

	buffer.WriteString("1\n")
	buffer.WriteString(bitboard.BoardHeader)

	return buffer.String()
}
