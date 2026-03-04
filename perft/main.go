// Command perft runs a chess move path enumeration (perft) to verify move generation correctness
// and measure performance. It accepts an optional depth and FEN position as arguments.
package main

import (
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"saiko.cz/sachista/chessboard"
)

func main() {
	logger := NewLogger()

	logger.info.Printf("Welcome to sachista-chess-go %v perfT!\n\n", runtime.GOARCH)

	board := chessboard.StandardBoard()
	depth := 7
	var err error

	switch {
	case len(os.Args) == 2:
		if depth, err = strconv.Atoi(os.Args[1]); err != nil {
			logger.err.Fatalf("Error: Invalid depth argument: %v\n", os.Args[1])
		}
	case len(os.Args) == 3:
		if depth, err = strconv.Atoi(os.Args[1]); err != nil {
			logger.err.Fatalf("Error: Invalid depth argument: %v\n", os.Args[1])
		}
		board = chessboard.BoardFromFEN(os.Args[2])
		if board.ZobristHash == 0 {
			logger.err.Fatalf("Error: Invalid FEN String - can not create Chess board: %v\n", os.Args[2])
		}
	case len(os.Args) > 3:
		logger.err.Printf("usage: [NO-ARGUMENTS] - running standard layout perft for the default depth of %v\n", depth)
		logger.err.Printf("usage: [DEPTH]        - running standard layout perft for the given depth\n")
		logger.err.Printf("usage: [DEPTH] [FEN]  - running custom board layout perft for the given depth\n")
	}

	start := time.Now()
	result := chessboard.PerfT(&board, depth)
	duration := time.Since(start)

	count := strconv.FormatUint(result, 10)
	if result <= math.MaxInt64 {
		count = humanize.Comma(int64(result))
	}

	logger.info.Println("perfT finished:")
	logger.info.Println("   FEN:   ", board.ToFEN())
	logger.info.Println("   depth: ", depth)
	logger.info.Println("   count: ", count)
	logger.info.Println("   time:  ", duration)
}

// Logger provides separate log outputs for informational messages (stdout) and errors (stderr).
type Logger struct {
	err  *log.Logger
	info *log.Logger
}

// NewLogger creates a Logger with stdout for info and stderr for errors.
func NewLogger() *Logger {
	return &Logger{
		info: log.New(os.Stdout, "", 0),
		err:  log.New(os.Stderr, "", 0),
	}
}
