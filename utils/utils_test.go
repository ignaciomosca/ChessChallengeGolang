package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAbsPositiveNumber(t *testing.T) {
	value := Abs(1)
	assert.Equal(t, 1, value, "test value of a positive number")
}

func TestAbsNegativeNumber(t *testing.T) {
	value := Abs(-1)
	assert.Equal(t, 1, value, "test value of a positive number")
}