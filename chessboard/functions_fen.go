package chessboard

import (
	"bytes"
	"saiko.cz/sachista/bitboard"
	"strconv"
)

func (b *Board) ToFEN() string {

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

	for i := 0; i < bitboard.BitWidth; i++ {
		if (i % 8) == 0 {
			outputCount()
			if i > 0 {
				buffer.WriteString("/")
			}
		}

		c := byte(0)
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

		if c != 0 {
			outputCount()
			buffer.WriteByte(c)
		} else {
			spaces++
		}
	}
	outputCount()

	//next move color
	buffer.WriteString(" ")
	buffer.WriteString(b.NextMove.String())
	buffer.WriteString(" ")

	//Castling
	if b.Castling[White]&CastlingKingSide != 0 {
		buffer.WriteByte(King.Description(White))
	}
	if b.Castling[White]&CastlingQueenSide != 0 {
		buffer.WriteByte(Queen.Description(White))
	}
	if b.Castling[Black]&CastlingKingSide != 0 {
		buffer.WriteByte(King.Description(Black))
	}
	if b.Castling[Black]&CastlingQueenSide != 0 {
		buffer.WriteByte(Queen.Description(Black))
	}
	if b.Castling[White]|b.Castling[Black] == 0 {
		buffer.WriteString("-")
	}

	//enPassant
	buffer.WriteString(" ")
	if b.EnPassantTarget != 0 {
		buffer.WriteString(b.EnPassantTarget.String())
	} else {
		buffer.WriteString("-")
	}

	//move number
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(b.HalfMoveClock))
	buffer.WriteString(" ")
	buffer.WriteString(strconv.Itoa(b.FullMoveNumber))

	return buffer.String()
}
