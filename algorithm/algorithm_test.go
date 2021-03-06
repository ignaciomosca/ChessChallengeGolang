package algorithm

import (
	"ChessChallengeGolang/board"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3x3Board2Kings1Rook(t *testing.T) {
	testBoard := board.CreateBoard(4, 4)
	var solutions []board.Board
	var testedConfigurations []board.Board
	Solution(testBoard, []rune{'K', 'K', 'R'}, &solutions, &testedConfigurations)

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
	testBoard := board.CreateBoard(5, 5)
	solutions := []board.Board{}
	testedConfigurations := []board.Board{}
	Solution(testBoard, []rune{'N', 'N', 'N', 'N', 'R', 'R'}, &solutions, &testedConfigurations)

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
	testBoard := board.CreateBoard(9, 9)
	var solutions []board.Board
	var testedConfigurations []board.Board
	Solution(testBoard, []rune{'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q'}, &solutions, &testedConfigurations)
	assert.Equal(t, len(solutions), 92, "Test N Queen Problem")
}

func BenchmarkSolution(b *testing.B) {
	testBoard := board.CreateBoard(8, 8)
	var solutions []board.Board
	var testedConfigurations []board.Board
	Solution(testBoard, []rune{'K', 'K', 'Q', 'Q', 'B', 'B', 'N'}, &solutions, &testedConfigurations)
	println("Solutions")
	println(len(solutions))
}
