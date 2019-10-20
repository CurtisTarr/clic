package tokenizer

import (
	"internal/assert"
	"testing"
)

func TestSimpleLiteral(t *testing.T) {
	var input = "5"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Literal, "5"))

	var tokens = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
}

func TestOperators(t *testing.T) {
	var input = "+-*/"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Operator, "+"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "-"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "*"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "/"))

	var tokens = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
}

func TestOpeningBrace(t *testing.T) {
	var input = "("
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(OpeningBrace, "("))

	var tokens = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
}

func TestClosingBrace(t *testing.T) {
	var input = ")"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(ClosingBrace, ")"))

	var tokens = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
}

func TestComplexExpression(t *testing.T) {
	var input = "3 + 5 * 2 / (3 - 1)"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Literal, "3"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "+"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "5"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "*"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "2"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "/"))
	expectedTokens = append(expectedTokens, NewToken(OpeningBrace, "("))
	expectedTokens = append(expectedTokens, NewToken(Literal, "3"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "-"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "1"))
	expectedTokens = append(expectedTokens, NewToken(ClosingBrace, ")"))

	var tokens = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
}
