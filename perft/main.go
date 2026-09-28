// Command perft runs a chess move path enumeration (perft) to verify move generation correctness
// and measure performance. It accepts an optional depth and FEN position as arguments.
package main

import (
	"errors"
	"fmt"
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

	board, depth, err := parseArgs(os.Args[1:])
	if err != nil {
		logger.err.Printf("Error: %v\n\n", err)
		logger.err.Printf("usage: [NO-ARGUMENTS] - running standard layout perft for the default depth of %v\n", defaultDepth)
		logger.err.Printf("usage: [DEPTH]        - running standard layout perft for the given depth\n")
		logger.err.Printf("usage: [DEPTH] [FEN]  - running custom board layout perft for the given depth\n")
		os.Exit(2)
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

// defaultDepth is the perft depth used when no depth argument is given.
const defaultDepth = 7

// parseArgs parses the command line arguments (without the program name) into the board and depth to run.
func parseArgs(args []string) (chessboard.Board, int, error) {
	board := chessboard.StandardBoard()
	depth := defaultDepth

	if len(args) > 2 {
		return board, 0, errors.New("too many arguments (quote the FEN string)")
	}

	if len(args) >= 1 {
		var err error
		if depth, err = strconv.Atoi(args[0]); err != nil || depth < 1 {
			return board, 0, fmt.Errorf("invalid depth argument: %v", args[0])
		}
	}

	if len(args) == 2 {
		board = chessboard.BoardFromFEN(args[1])
		// BoardFromFEN does not report errors; a position without exactly one king per side is unusable
		if board.Pieces[chessboard.White][chessboard.King].PopCount() != 1 || board.Pieces[chessboard.Black][chessboard.King].PopCount() != 1 {
			return board, 0, fmt.Errorf("invalid FEN string - can not create chess board: %v", args[1])
		}
	}

	return board, depth, nil
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
