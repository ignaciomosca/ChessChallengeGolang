package main

import (
	"ChessChallengeGolang/algorithm"
	"ChessChallengeGolang/board"
	"ChessChallengeGolang/primitives"
)

func main() {
	testBoard := board.CreateBoard(8, 8)
	var solutions []board.Board
	var testedConfigurations map[int][]primitives.Attacks

	algorithm.Solution(testBoard, []rune{'K', 'K', 'Q', 'Q', 'B', 'B', 'N'}, &solutions, testedConfigurations)

	println("Solutions")
	println(len(solutions))
}
