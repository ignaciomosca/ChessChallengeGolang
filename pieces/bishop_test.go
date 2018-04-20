package pieces

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBishopAttacks(t *testing.T) {
	bishopA := CreateBishop(5, 5)
	bishopB := CreateBishop(7, 7)
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")
}

func TestBishopAttack2(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(6, 6);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestBishopAttack3(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(4, 6);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestBishopAttack4(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(3, 7);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestBishopAttack5(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(7, 3);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestBishopAttack6(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(6, 4);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestBishopAttack7(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(4, 4);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestBishopAttack8(t *testing.T) {
	bishopA := CreateBishop(5, 5);
	bishopB := CreateBishop(3, 3);
	result := bishopA.Attacks(bishopB)
	assert.Equal(t, result, true, "test bishop A attacks bishop B")}

func TestBishopDoesNotAttack(t *testing.T) {
	bishopA := CreateBishop(1, 1);
	bishopB := CreateBishop(8, 1);
	result := bishopA.Attacks(bishopB)
	assert.NotEqual(t, result, true, "test bishop A attacks bishop B")
}
