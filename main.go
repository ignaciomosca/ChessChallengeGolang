package main

import (
	"chesschallengegolang/algorithm"
	"chesschallengegolang/board"
)

func main() {
	testBoard := board.CreateBoard(9, 9)
	solutions := make(map[string]bool)
	testedConfigurations := make(map[string]bool)

	algorithm.Solution(testBoard, []rune{'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q'}, &solutions, &testedConfigurations)

	println("Solutions")
	println(len(solutions))
}
