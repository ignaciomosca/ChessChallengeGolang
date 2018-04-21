package algorithm

import (
	"testing"
	"chess/board"
	"github.com/stretchr/testify/assert"
)

func TestNQueen(t *testing.T) {
	testBoard := board.CreateBoard(8, 8)
	solutions := Solution(testBoard, []rune{'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q', 'Q'}, make(map[*board.Board]bool), make(map[*board.Board]bool))

	keys := []board.Board{}
	for k := range solutions {
		keys = append(keys, *k)
	}

	for _, b := range keys {
		board.ShowBoard(b)
	}

	assert.Equal(t, len(solutions), 92, "Test N Queen Problem")
}
