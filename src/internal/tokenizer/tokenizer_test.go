package tokenizer

import (
	"internal/assert"
	"testing"
)

func TestSimpleLiteral(t *testing.T) {
	var input = "5"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Literal, "5"))

	var tokens, err = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
	assert.Equal(t, err, nil)
}

func TestOperators(t *testing.T) {
	var input = "+-*/^"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Operator, "+"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "-"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "*"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "/"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "^"))

	var tokens, err = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
	assert.Equal(t, err, nil)
}

func TestOpeningBrace(t *testing.T) {
	var input = "("
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(OpeningBrace, "("))

	var tokens, err = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
	assert.Equal(t, err, nil)
}

func TestClosingBrace(t *testing.T) {
	var input = ")"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(ClosingBrace, ")"))

	var tokens, err = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
	assert.Equal(t, err, nil)
}

func TestLongNumbers(t *testing.T) {
	var input = "3452 + 2"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Literal, "3452"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "+"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "2"))

	var tokens, err = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
	assert.Equal(t, err, nil)
}

func TestDecimal(t *testing.T) {
	var input = "34.52 + 2"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Literal, "34.52"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "+"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "2"))

	var tokens, err = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
	assert.Equal(t, err, nil)
}

func TestComplexExpression(t *testing.T) {
	var input = "3 + 5.3 * 2 / (3 - 1)"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Literal, "3"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "+"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "5.3"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "*"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "2"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "/"))
	expectedTokens = append(expectedTokens, NewToken(OpeningBrace, "("))
	expectedTokens = append(expectedTokens, NewToken(Literal, "3"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "-"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "1"))
	expectedTokens = append(expectedTokens, NewToken(ClosingBrace, ")"))

	var tokens, err = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
	assert.Equal(t, err, nil)
}

func TestUnknownValues(t *testing.T) {
	var input = "34.5a2 + 2b _"

	var _, err = Tokenize(input)

	assert.Equal(t, err.Error(), "unknown token detected [a] at location [4]")
}

func TestNegativeLiteral(t *testing.T) {
	var input = "-3"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Literal, "-3"))

	var tokens, err = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
	assert.Equal(t, err, nil)
}

func TestDeepNegativeLiteral(t *testing.T) {
	var input = "4 + 2 - -3"
	var expectedTokens []*Token
	expectedTokens = append(expectedTokens, NewToken(Literal, "4"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "+"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "2"))
	expectedTokens = append(expectedTokens, NewToken(Operator, "-"))
	expectedTokens = append(expectedTokens, NewToken(Literal, "-3"))

	var tokens, err = Tokenize(input)

	assert.DeepEqual(t, tokens, expectedTokens)
	assert.Equal(t, err, nil)
}
