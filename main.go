package main

import (
	"ChessChallengeGolang/algorithm"
	"ChessChallengeGolang/board"
)

func main() {
	testBoard := board.CreateBoard(9, 9)
	var solutions []board.Board
	var testedConfigurations []board.Board

	algorithm.Solution(testBoard, []rune{'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q'}, &solutions, &testedConfigurations)

	println("Solutions")
	println(len(solutions))
}
