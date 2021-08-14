package main

import (
	"fmt"
	"saiko.cz/sachista/chessboard"
	"saiko.cz/sachista/generator"
	"time"
)

//TODO try release build
//TODO create parametric main function
func main() {
	board := chessboard.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	depth := 7
	want := 3_195_901_860

	start := time.Now()
	result := generator.PerfT(board, depth)
	duration := time.Since(start)

	fmt.Println(result, want)
	fmt.Println(duration)
}
