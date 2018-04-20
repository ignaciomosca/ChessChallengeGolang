package pieces

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQueenAttacks(t *testing.T) {
	queenA := CreateQueen(5, 5)
	queenB := CreateQueen(7, 7)
	result := queenA.Attacks(queenB)
	assert.Equal(t, result, true, "test queen A attacks queen B")
}

func TestQueenAttack2(t *testing.T) {
	queenA := CreateQueen(5, 5);
	queenB := CreateQueen(6, 6);
	result := queenA.Attacks(queenB)
	assert.Equal(t, result, true, "test queen A attacks queen B")}

func TestQueenAttack3(t *testing.T) {
	queenA := CreateQueen(5, 5);
	queenB := CreateQueen(4, 6);
	result := queenA.Attacks(queenB)
	assert.Equal(t, result, true, "test queen A attacks queen B")}

func TestQueenAttack4(t *testing.T) {
	queenA := CreateQueen(5, 5);
	queenB := CreateQueen(3, 7);
	result := queenA.Attacks(queenB)
	assert.Equal(t, result, true, "test queen A attacks queen B")}

func TestQueenAttack5(t *testing.T) {
	queenA := CreateQueen(5, 5);
	queenB := CreateQueen(7, 3);
	result := queenA.Attacks(queenB)
	assert.Equal(t, result, true, "test queen A attacks queen B")}

func TestQueenAttack6(t *testing.T) {
	queenA := CreateQueen(5, 5);
	queenB := CreateQueen(6, 4);
	result := queenA.Attacks(queenB)
	assert.Equal(t, result, true, "test queen A attacks queen B")}

func TestQueenAttack7(t *testing.T) {
	queenA := CreateQueen(5, 5);
	queenB := CreateQueen(4, 4);
	result := queenA.Attacks(queenB)
	assert.Equal(t, result, true, "test queen A attacks queen B")}

func TestQueenAttack8(t *testing.T) {
	queenA := CreateQueen(5, 5);
	queenB := CreateQueen(3, 3);
	result := queenA.Attacks(queenB)
	assert.Equal(t, result, true, "test queen A attacks queen B")}

func TestQueenDoesNotAttack(t *testing.T) {
	queenA := CreateQueen(1, 1);
	queenB := CreateQueen(3, 4);
	result := queenA.Attacks(queenB)
	assert.NotEqual(t, result, true, "test queen A attacks queen B")
}
