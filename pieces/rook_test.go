package pieces

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRookAttacks(t *testing.T) {
	rookA := CreateRook(5, 5)
	rookB := CreateRook(6, 5)
	result := rookA.Attacks(rookB)
	assert.Equal(t, result, true, "test rook A attacks rook B")
}

func TestRookAttack2(t *testing.T) {
	rookA := CreateRook(5, 5);
	rookB := CreateRook(7, 5);
	result := rookA.Attacks(rookB)
	assert.Equal(t, result, true, "test rook A attacks rook B")
}

func TestRookAttack3(t *testing.T) {
	rookA := CreateRook(5, 5);
	rookB := CreateRook(4, 5);
	result := rookA.Attacks(rookB)
	assert.Equal(t, result, true, "test rook A attacks rook B")
}

func TestRookAttack4(t *testing.T) {
	rookA := CreateRook(5, 5);
	rookB := CreateRook(3, 5);
	result := rookA.Attacks(rookB)
	assert.Equal(t, result, true, "test rook A attacks rook B")
}

func TestRookAttack5(t *testing.T) {
	rookA := CreateRook(5, 5);
	rookB := CreateRook(5, 6);
	result := rookA.Attacks(rookB)
	assert.Equal(t, result, true, "test rook A attacks rook B")
}

func TestRookAttack6(t *testing.T) {
	rookA := CreateRook(5, 5);
	rookB := CreateRook(5, 7);
	result := rookA.Attacks(rookB)
	assert.Equal(t, result, true, "test rook A attacks rook B")
}

func TestRookAttack7(t *testing.T) {
	rookA := CreateRook(5, 5);
	rookB := CreateRook(5, 4);
	result := rookA.Attacks(rookB)
	assert.Equal(t, result, true, "test rook A attacks rook B")
}

func TestRookAttack8(t *testing.T) {
	rookA := CreateRook(5, 5);
	rookB := CreateRook(5, 3);
	result := rookA.Attacks(rookB)
	assert.Equal(t, result, true, "test rook A attacks rook B")
}

func TestRookDoesNotAttack(t *testing.T) {
	rookA := CreateRook(1, 1);
	rookB := CreateRook(7, 2);
	result := rookA.Attacks(rookB)
	assert.NotEqual(t, result, true, "test rook A attacks rook B")
}
