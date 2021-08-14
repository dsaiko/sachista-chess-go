package main

import (
	"fmt"
	"os"
	"runtime"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/generator"
	"strconv"
	"time"
)

func main() {

	fmt.Printf("Welcome to sachista-chess-go %v perfT!\n\n", runtime.GOARCH)

	board := chessboard.Standard()
	depth := 7
	var err error

	if len(os.Args) < 2 {
		fmt.Printf("usage: [NO-ARGUMENTS] - running standard layout perft for the default depth of %v\n", depth)
		fmt.Printf("usage: [DEPTH]        - running standard layout perft for the given depth\n")
		fmt.Printf("usage: [DEPTH] [FEN]  - running custom board layout perft for the given depth\n")
	}

	if len(os.Args) > 1 {
		if depth, err = strconv.Atoi(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid depth argument: %v\n", os.Args[1])
			return
		}
	}

	if len(os.Args) > 2 {
		board = chessboard.FromFEN(os.Args[2])
		if board.ZobristHash == 0 {
			fmt.Fprintf(os.Stderr, "Error: Invalid FEN String - can not create Chess board: %v\n", os.Args[2])
			return
		}
	}

	start := time.Now()
	result := generator.PerfT(board, depth)
	duration := time.Since(start)

	fmt.Println("perfT finished:")
	fmt.Println("   FEN:   ", board.ToFEN())
	fmt.Println("   depth: ", depth)
	fmt.Println("   count: ", formatNumber(result))
	fmt.Println("   time:  ", duration)
}

// formatNumber ...
// https://stackoverflow.com/questions/13020308/how-to-fmt-printf-an-integer-with-thousands-comma
func formatNumber(n uint64) string {
	in := strconv.FormatUint(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
