package board

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"chess/primitives"
	"chess/pieces"
)

func TestBoardEquals(t *testing.T) {
	boardA := CreateBoard(1, 1)
	boardB := CreateBoard(1, 1)
	result := reflect.DeepEqual(boardA, boardB)
	assert.Equal(t, result, true, "board A == board B")
}

func TestBoard_Place(t *testing.T) {
	board := CreateBoard(9, 9)
	king := pieces.CreateKing(1,1)
	Place(board, king)
	assert.Equal(t, 1, len(board.used), "add piece to board")

	queen := pieces.CreateQueen(2,2)
	Place(board, queen)
	assert.Equal(t, 2, len(board.used), "add another piece to board")
}

func TestSameBoardWithPieces(t *testing.T) {
	kingA := pieces.CreateKing(1,1)
	queenA := pieces.CreateQueen(2,2)
	piecesA := map[primitives.Attacks]bool{ kingA:true, queenA: true}
	piecesB := map[primitives.Attacks]bool{ queenA: true, kingA:true}
	boardA := CreateBoardWithPieces(9, 9, piecesA)
	boardB := CreateBoardWithPieces(9, 9, piecesB)
	result := reflect.DeepEqual(boardA, boardB)
	assert.Equal(t, result, true, "board A == board B")
}

func TestDifferentBoardWithPieces(t *testing.T) {
	kingA := pieces.CreateKing(1,1)
	queenA := pieces.CreateQueen(2,2)
	queenB := pieces.CreateQueen(2,3)
	piecesA := map[primitives.Attacks]bool{ kingA:true, queenA: true}
	piecesB := map[primitives.Attacks]bool{ queenB: true, kingA:true}
	boardA := CreateBoardWithPieces(9, 9, piecesA)
	boardB := CreateBoardWithPieces(9, 9, piecesB)
	result := reflect.DeepEqual(boardA, boardB)
	assert.NotEqual(t, result, true, "board A == board B")
}

