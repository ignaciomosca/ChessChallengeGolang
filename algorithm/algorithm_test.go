package algorithm

import (
	"testing"
	"chess/board"
)

func TestNQueen(t *testing.T) {
	testBoard := board.CreateBoard(8, 8)
	solutions := Solution(testBoard, []rune{'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q'}, make(map[*board.Board]bool), make(map[*board.Board]bool))

	keys := make([]board.Board, len(solutions))
	for k := range solutions {
		keys = append(keys, *k)
	}

	for i:=0 ; i< len(keys); i++{
		board.ShowBoard(keys[i])
	}

	//assert.Equal(t, len(solutions), 92, "Test N Queen Problem")
}
