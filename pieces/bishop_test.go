package pieces

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAttacks(t *testing.T) {
	bishopA := CreateBishop(5, 5)
	bishopB := CreateBishop(7, 7)
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")
}

func TestAttack2(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(6, 6);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestAttack3(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(4, 6);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestAttack4(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(3, 7);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestAttack5(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(7, 3);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestAttack6(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(6, 4);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestAttack7(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(4, 4);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestAttack8(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(3, 3);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestDoesNotAttack(t *testing.T) {
	bishopA := CreateBishop(1, 1);
	bishopB := CreateBishop(8, 1);
	result := bishopA.Attacks(bishopB)
	assert.NotEqual(t, result, true, "test bishop A attacks bishop B")
}
