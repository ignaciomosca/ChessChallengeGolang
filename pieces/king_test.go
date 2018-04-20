package pieces

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKingAttacks(t *testing.T) {
	kingA := CreateKing(5, 5)
	kingB := CreateKing(5, 6)
	result := kingA.Attacks(kingB)
	assert.Equal(t, result, true, "test king A attacks king B")
}

func TestKingAttack2(t *testing.T) {
	kingA := CreateKing(5, 5);
	kingB := CreateKing(6, 4);
	result := kingA.Attacks(kingB)
	assert.Equal(t, result, true, "test king A attacks king B")}

func TestKingAttack3(t *testing.T) {
	kingA := CreateKing(5, 5);
	kingB := CreateKing(5, 4);
	result := kingA.Attacks(kingB)
	assert.Equal(t, result, true, "test king A attacks king B")}

func TestKingAttack4(t *testing.T) {
	kingA := CreateKing(5, 5);
	kingB := CreateKing(4, 4);
	result := kingA.Attacks(kingB)
	assert.Equal(t, result, true, "test king A attacks king B")}

func TestKingAttack5(t *testing.T) {
	kingA := CreateKing(5, 5);
	kingB := CreateKing(4, 5);
	result := kingA.Attacks(kingB)
	assert.Equal(t, result, true, "test king A attacks king B")}

func TestKingAttack6(t *testing.T) {
	kingA := CreateKing(5, 5);
	kingB := CreateKing(4, 6);
	result := kingA.Attacks(kingB)
	assert.Equal(t, result, true, "test king A attacks king B")}

func TestKingAttack7(t *testing.T) {
	kingA := CreateKing(5, 5);
	kingB := CreateKing(5, 6);
	result := kingA.Attacks(kingB)
	assert.Equal(t, result, true, "test king A attacks king B")}

func TestKingAttack8(t *testing.T) {
	kingA := CreateKing(5, 5);
	kingB := CreateKing(6, 6);
	result := kingA.Attacks(kingB)
	assert.Equal(t, result, true, "test king A attacks king B")}

func TestKingDoesNotAttack(t *testing.T) {
	kingA := CreateKing(1, 1);
	kingB := CreateKing(8, 8);
	result := kingA.Attacks(kingB)
	assert.NotEqual(t, result, true, "test king A attacks king B")
}
