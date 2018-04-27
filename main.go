package main

import (
	"chess/board"
	"chess/algorithm"
)

func main() {
	testBoard := board.CreateBoard(8, 8)
	var solutions []board.Board
	var testedConfigurations []board.Board

	algorithm.Solution(testBoard, []rune{'K', 'K', 'Q', 'Q', 'B', 'B', 'N'}, &solutions, &testedConfigurations)

	println("Solutions")
	println(len(solutions))
}
