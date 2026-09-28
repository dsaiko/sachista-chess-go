package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"saiko.cz/sachista/chessboard"
)

func TestParseArgs(t *testing.T) {
	const fen = "8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1"

	tests := []struct {
		name      string
		args      []string
		wantDepth int
		wantFEN   string
		wantErr   bool
	}{
		{name: "defaults", args: nil, wantDepth: defaultDepth, wantFEN: chessboard.StandardBoardFEN},
		{name: "depth", args: []string{"5"}, wantDepth: 5, wantFEN: chessboard.StandardBoardFEN},
		{name: "depth and FEN", args: []string{"4", fen}, wantDepth: 4, wantFEN: fen},
		{name: "non-numeric depth", args: []string{"x"}, wantErr: true},
		{name: "zero depth", args: []string{"0"}, wantErr: true},
		{name: "negative depth", args: []string{"-3"}, wantErr: true},
		{name: "invalid FEN", args: []string{"3", "XXXXX"}, wantErr: true},
		{name: "FEN without black king", args: []string{"3", "8/8/8/8/8/8/8/4K3 w - - 0 1"}, wantErr: true},
		{name: "unquoted FEN", args: []string{"3", "8/8/8/8/8/8/8/4K3", "w", "-"}, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			board, depth, err := parseArgs(tc.args)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantDepth, depth)
			assert.Equal(t, tc.wantFEN, board.ToFEN())
		})
	}
}
