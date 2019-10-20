package maths

import (
	"internal/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	assert.Equal(t, total, 10.)
}

func TestSubtract(t *testing.T) {
	total := Subtract(7, 3)
	assert.Equal(t, total, 4.)
}

func TestMultiply(t *testing.T) {
	total := Multiply(5, 5)
	assert.Equal(t, total, 25.)
}

func TestDivide(t *testing.T) {
	total := Divide(10, 2)
	assert.Equal(t, total, 5.)
}
