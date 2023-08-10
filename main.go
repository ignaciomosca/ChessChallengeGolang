package main

import (
	"chesschallengegolang/algorithm"
)

func main() {
	pieceCounts := map[rune]int{'K': 2, 'Q': 2, 'B': 2, 'R': 0, 'N': 1}

	solutions := algorithm.Solution(8, 8, pieceCounts)

	println("Solutions")
	println(len(solutions))
}
