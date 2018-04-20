package pieces

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKnightAttacks(t *testing.T) {
	knightA := CreateKnight(5, 5)
	knightB := CreateKnight(7, 6)
	result := knightA.Attacks(knightB)
	assert.Equal(t, result, true, "test knight A attacks knight B")
}

func TestKnightAttack2(t *testing.T) {
	knightA := CreateKnight(5, 5);
	knightB := CreateKnight(6, 7);
	result := knightA.Attacks(knightB)
	assert.Equal(t, result, true, "test knight A attacks knight B")}

func TestKnightAttack3(t *testing.T) {
	knightA := CreateKnight(5, 5);
	knightB := CreateKnight(4, 7);
	result := knightA.Attacks(knightB)
	assert.Equal(t, result, true, "test knight A attacks knight B")}

func TestKnightAttack4(t *testing.T) {
	knightA := CreateKnight(5, 5);
	knightB := CreateKnight(3, 6);
	result := knightA.Attacks(knightB)
	assert.Equal(t, result, true, "test knight A attacks knight B")}

func TestKnightAttack5(t *testing.T) {
	knightA := CreateKnight(5, 5);
	knightB := CreateKnight(3, 4);
	result := knightA.Attacks(knightB)
	assert.Equal(t, result, true, "test knight A attacks knight B")}

func TestKnightAttack6(t *testing.T) {
	knightA := CreateKnight(5, 5);
	knightB := CreateKnight(4, 3);
	result := knightA.Attacks(knightB)
	assert.Equal(t, result, true, "test knight A attacks knight B")}

func TestKnightAttack7(t *testing.T) {
	knightA := CreateKnight(5, 5);
	knightB := CreateKnight(7, 4);
	result := knightA.Attacks(knightB)
	assert.Equal(t, result, true, "test knight A attacks knight B")}

func TestKnightAttack8(t *testing.T) {
	knightA := CreateKnight(5, 5);
	knightB := CreateKnight(6, 3);
	result := knightA.Attacks(knightB)
	assert.Equal(t, result, true, "test knight A attacks knight B")}

func TestKnightDoesNotAttack(t *testing.T) {
	knightA := CreateKnight(1, 1);
	knightB := CreateKnight(8, 8);
	result := knightA.Attacks(knightB)
	assert.NotEqual(t, result, true, "test knight A attacks knight B")
}
