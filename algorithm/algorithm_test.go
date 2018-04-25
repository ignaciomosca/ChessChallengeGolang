package algorithm

import (
	"testing"
	"chess/board"
	"github.com/stretchr/testify/assert"
)

func TestNQueen(t *testing.T) {
	testBoard := board.CreateBoard(9, 9)
	solutions := []board.Board{}
	testedConfigurations:= []board.Board{}
	Solution(testBoard, []rune{'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q'}, &solutions, &testedConfigurations)

	keys := []board.Board{}
	for _, k := range solutions {
		keys = append(keys, k)
	}

	for _, b := range keys {
		board.ShowBoard(b)
	}

	assert.Equal(t, len(solutions), 92, "Test N Queen Problem")
}

func Test3x3Board2Kings1Rook(t *testing.T){
	testBoard := board.CreateBoard(4, 4)
	var solutions []board.Board
	var testedConfigurations []board.Board
	Solution(testBoard, []rune{'K','K','R'}, &solutions, &testedConfigurations)

	keys := []board.Board{}
	for _,k := range solutions {
		keys = append(keys, k)
	}

	for _, b := range keys {
		board.ShowBoard(b)
	}

	assert.Equal(t, len(solutions), 4, "Test 3x3 Problem")
}

func Test4x4Board2Rooks4Knights(t *testing.T){
	testBoard := board.CreateBoard(5, 5)
	solutions := []board.Board{}
	testedConfigurations:= []board.Board{}
	Solution(testBoard, []rune{'N','N','N', 'N', 'R', 'R'}, &solutions, &testedConfigurations)

	keys := []board.Board{}
	for _,k := range solutions {
		keys = append(keys, k)
	}

	for _, b := range keys {
		board.ShowBoard(b)
	}

	assert.Equal(t, len(solutions), 8, "Test 4x4 Problem")
}