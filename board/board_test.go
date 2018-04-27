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
	assert.Equal(t, result, true, "Board A == Board B")
}

func TestBoard_Place(t *testing.T) {
	board := CreateBoard(9, 9)
	king := pieces.CreateKing(1,1)
	newBoard := Place(board, king)
	assert.Equal(t, 1, len(newBoard.Used), "add piece to Board")

	queen := pieces.CreateQueen(2,2)
	Place(newBoard, queen)
	assert.Equal(t, 2, len(newBoard.Used), "add another piece to Board")
}

func TestSameBoardWithPieces(t *testing.T) {
	kingA := pieces.CreateKing(1,1)
	queenA := pieces.CreateQueen(2,2)
	piecesA := []primitives.Attacks{ kingA, queenA}
	piecesB := []primitives.Attacks{ queenA, kingA}
	boardA := CreateBoardWithPieces(9, 9, piecesA)
	boardB := CreateBoardWithPieces(9, 9, piecesB)
	result := reflect.DeepEqual(boardA, boardB)
	assert.Equal(t, result, true, "Board A == Board B")
}

func TestDifferentBoardWithPieces(t *testing.T) {
	kingA := pieces.CreateKing(1,1)
	queenA := pieces.CreateQueen(2,2)
	queenB := pieces.CreateQueen(2,3)
	piecesA := []primitives.Attacks{ kingA, queenA}
	piecesB := []primitives.Attacks{ queenB, kingA}
	boardA := CreateBoardWithPieces(9, 9, piecesA)
	boardB := CreateBoardWithPieces(9, 9, piecesB)
	result := reflect.DeepEqual(boardA, boardB)
	assert.NotEqual(t, result, true, "Board A == Board B")
}

