package algorithm

import (
	"chesschallengegolang/board"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3x3Board2Kings1Rook(t *testing.T) {
	pieceCounts := map[rune]int{'K': 2, 'Q': 0, 'B': 0, 'R': 1, 'N': 0}
	solutions := Solution(4, 4, pieceCounts)

	keys := []board.Board{}
	for _, k := range solutions {
		keys = append(keys, k)
	}

	for _, b := range keys {
		board.ShowBoard(b)
	}

	assert.Equal(t, 4, len(solutions), "Test 3x3 Problem")
}

func Test4x4Board2Rooks4Knights(t *testing.T) {
	pieceCounts := map[rune]int{'K': 0, 'Q': 0, 'B': 0, 'R': 2, 'N': 4}
	solutions := Solution(5, 5, pieceCounts)

	keys := []board.Board{}
	for _, k := range solutions {
		keys = append(keys, k)
	}

	for _, b := range keys {
		board.ShowBoard(b)
	}

	assert.Equal(t, len(solutions), 8, "Test 4x4 Problem")
}

func TestNQueen(t *testing.T) {
	pieceCounts := map[rune]int{'K': 0, 'Q': 8, 'B': 0, 'R': 0, 'N': 0}
	solutions := Solution(9, 9, pieceCounts)
	assert.Equal(t, len(solutions), 92, "Test N Queen Problem")
}

func BenchmarkSolution(b *testing.B) {
	pieceCounts := map[rune]int{'K': 2, 'Q': 2, 'B': 2, 'R': 0, 'N': 1}
	solutions := Solution(8, 8, pieceCounts)
	println("Solutions")
	println(len(solutions))
}
