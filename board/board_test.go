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

func TestSameBoardWithPieces(t *testing.T) {
	kingA := pieces.CreateKing(1,1)
	queenA := pieces.CreateQueen(2,2)
	piecesA := map[primitives.Attacks]bool{ kingA:true, queenA: true}
	piecesB := map[primitives.Attacks]bool{ queenA: true, kingA:true}
	boardA := CreateBoardWithPieces(8, 8, piecesA)
	boardB := CreateBoardWithPieces(8, 8, piecesB)
	result := reflect.DeepEqual(boardA, boardB)
	assert.Equal(t, result, true, "board A == board B")
}

func TestDifferentBoardWithPieces(t *testing.T) {
	kingA := pieces.CreateKing(1,1)
	queenA := pieces.CreateQueen(2,2)
	queenB := pieces.CreateQueen(2,3)
	piecesA := map[primitives.Attacks]bool{ kingA:true, queenA: true}
	piecesB := map[primitives.Attacks]bool{ queenB: true, kingA:true}
	boardA := CreateBoardWithPieces(8, 8, piecesA)
	boardB := CreateBoardWithPieces(8, 8, piecesB)
	result := reflect.DeepEqual(boardA, boardB)
	assert.NotEqual(t, result, true, "board A == board B")
}

