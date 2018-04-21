package pieces

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreatePiece_Queen(t *testing.T) {
	createQueen := CreatePiece('Q',1,1)
	actual := createQueen.Name() == 'Q'
	assert.Equal(t, true, actual, "Test Create Queen")

	createKing := CreatePiece('K',1,1)
	actual = createKing.Name() == 'K'
	assert.Equal(t, true, actual, "Test Create King")

	createKnight := CreatePiece('N',1,1)
	actual = createKnight.Name() == 'N'
	assert.Equal(t, true, actual, "Test Create Knight")

	createBishop := CreatePiece('B',1,1)
	actual = createBishop.Name() == 'B'
	assert.Equal(t, true, actual, "Test Create Bishop")

	createRook := CreatePiece('R',1,1)
	actual = createRook.Name() == 'R'
	assert.Equal(t, true, actual, "Test Create Rook")
}
