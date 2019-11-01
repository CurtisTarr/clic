package compiler

import (
	"internal/assert"
	"internal/parser"
	"testing"
)

func TestCompile1(t *testing.T) {
	var operations = parser.NewNode("+",
		parser.NewNode("3", nil, nil),
		parser.NewNode("/",
			parser.NewNode("*",
				parser.NewNode("4", nil, nil),
				parser.NewNode("2", nil, nil)),
			parser.NewNode("^",
				parser.NewNode("-",
					parser.NewNode("1", nil, nil),
					parser.NewNode("5", nil, nil)),
				parser.NewNode("^",
					parser.NewNode("2", nil, nil),
					parser.NewNode("3", nil, nil)))))

	var expectedOutput = "3.000122"
	var output = Compile(operations)

	assert.Equal(t, output, expectedOutput)
}

func TestCompile2(t *testing.T) {
	var operations = parser.NewNode("+",
		parser.NewNode("4", nil, nil),
		parser.NewNode("/",
			parser.NewNode("18", nil, nil),
			parser.NewNode("-",
				parser.NewNode("9", nil, nil),
				parser.NewNode("3", nil, nil))))

	var expectedOutput = "7.000000"
	var output = Compile(operations)

	assert.Equal(t, output, expectedOutput)
}

func TestAdd(t *testing.T) {
	total := performFunction("+", "5", "5")
	assert.Equal(t, total, "10.000000")
}

func TestSubtract(t *testing.T) {
	total := performFunction("-", "7", "3")
	assert.Equal(t, total, "4.000000")
}

func TestMultiply(t *testing.T) {
	total := performFunction("*", "5", "5")
	assert.Equal(t, total, "25.000000")
}

func TestDivide(t *testing.T) {
	total := performFunction("/", "10", "2")
	assert.Equal(t, total, "5.000000")
}

func TestPower(t *testing.T) {
	total := performFunction("^", "10", "2")
	assert.Equal(t, total, "100.000000")
}

func TestNegativeAdd(t *testing.T) {
	total := performFunction("+", "10", "-2")
	assert.Equal(t, total, "8.000000")
}