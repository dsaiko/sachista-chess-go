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
	board := chessboard.FromFen("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq -")
	depth := 5
	want := 193_690_690

	start := time.Now()
	result := generator.PerfT(board, depth)
	duration := time.Since(start)

	fmt.Println(result, want)
	fmt.Println(duration)
}
