package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"saiko.cz/sachista/bitboard"
)

func TestBoard_RemoveCastling(t *testing.T) {
	b := EmptyBoard()

	assert.Equal(t, CastlingNone, b.Castling[White])
	assert.Equal(t, CastlingNone, b.Castling[Black])

	b.Castling[White] = CastlingBothSides
	b.RemovedCastling(White, CastlingQueenSide)

	assert.Equal(t, CastlingKingSide, b.Castling[White])
}

func TestBoard_Stats(t *testing.T) {
	b := StandardBoard()

	assert.Equal(t, 32, b.AllPieces().PopCount())
	assert.Equal(t, 16, b.OpponentPieces().PopCount())
	assert.Equal(t, 64-16, b.BoardAvailableToAttack().PopCount())

	assert.Equal(t, bitboard.IndexE1, b.MyKingIndex())
	assert.Equal(t, bitboard.IndexE8, b.OpponentKingIndex())
}

func TestBoard_OpponentColor(t *testing.T) {
	b := EmptyBoard()
	b.NextMove = White
	assert.Equal(t, Black, b.OpponentColor())
	b.NextMove = Black
	assert.Equal(t, White, b.OpponentColor())
}

func TestBoard_MyPieces(t *testing.T) {
	b := StandardBoard()
	b.NextMove = White
	assert.Equal(t, b.PiecesByColor(White), b.MyPieces())
	b.NextMove = Black
	assert.Equal(t, b.PiecesByColor(Black), b.MyPieces())
}

// assertOccupiedConsistent checks that the cached Occupied bitboards match the union of Pieces for each color.
func assertOccupiedConsistent(t *testing.T, b *Board) {
	t.Helper()
	for color := White; color <= Black; color++ {
		want := bitboard.EmptyBoard
		for piece := King; piece <= Pawn; piece++ {
			want |= b.Pieces[color][piece]
		}
		assert.Equal(t, want, b.Occupied[color], "Occupied[%v]", color)
	}
}

func TestBoard_RecomputeOccupied(t *testing.T) {
	b := StandardBoard()
	// Corrupt Occupied then recompute
	b.Occupied[White] = bitboard.EmptyBoard
	b.Occupied[Black] = bitboard.EmptyBoard
	b.RecomputeOccupied()
	assert.Equal(t, bitboard.BoardRanks[0]|bitboard.BoardRanks[1], b.Occupied[White])
	assert.Equal(t, bitboard.BoardRanks[6]|bitboard.BoardRanks[7], b.Occupied[Black])

	// asymmetric position: every color and piece type must land in the right Occupied entry
	b = BoardFromFEN("r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1")
	b.Occupied[White] = bitboard.UniverseBoard
	b.Occupied[Black] = bitboard.UniverseBoard
	b.RecomputeOccupied()
	assertOccupiedConsistent(t, &b)
	assert.Equal(t, 16, b.Occupied[White].PopCount())
	assert.Equal(t, 15, b.Occupied[Black].PopCount())
}

func TestBoard_UpdateZobrist(t *testing.T) {
	b1 := StandardBoard()
	b2 := StandardBoard()

	assert.Equal(t, b1.ZobristHash, b2.ZobristHash)

	// no halfmove counters are relevant
	b2.HalfMoveClock = 99
	b2.FullMoveNumber = 99
	b2.ZobristHash = b2.Hash()
	assert.Equal(t, b1.ZobristHash, b2.ZobristHash)

	b2.Pieces[White][Pawn] |= bitboard.BoardA3
	b2.ZobristHash = b2.Hash()
	assert.NotEqual(t, b1.ZobristHash, b2.ZobristHash)
}
