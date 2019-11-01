package calculator

import (
	"internal/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	var expectedOutput = "7.000000"
	var output = Calculate("4 + 18 / (9 - 3)")

	assert.Equal(t, output, expectedOutput)
}

func TestCase2(t *testing.T) {
	var expectedOutput = "81.000000"
	var output = Calculate("4 + 20 - -3 * (4 ^ 2 + 3)")

	assert.Equal(t, output, expectedOutput)
}

func TestCase3(t *testing.T) {
	var expectedOutput = "2.000000"
	var output = Calculate("4 / 2)")

	assert.Equal(t, output, expectedOutput)
}

func TestCase4(t *testing.T) {
	var expectedOutput = "7.000000"
	var output = Calculate("4 - -3")

	assert.Equal(t, output, expectedOutput)
}

func TestCase5(t *testing.T) {
	var expectedOutput = "3.000122"
	var output = Calculate("3 + 4 * 2 / (1 - 5) ^ 2 ^ 3")

	assert.Equal(t, output, expectedOutput)
}