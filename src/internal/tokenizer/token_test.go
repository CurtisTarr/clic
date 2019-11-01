package tokenizer

import (
	"internal/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	var tokenType = Literal
	var value = "45"

	var token = NewToken(tokenType, value)

	assert.Equal(t, token.TokenType, tokenType)
	assert.Equal(t, token.Value, value)
}

func TestGetAssoc(t *testing.T) {
	var power = NewToken(Operator, "^")
	var multiply = NewToken(Operator, "*")
	var divide = NewToken(Operator, "/")
	var add = NewToken(Operator, "+")
	var subtract = NewToken(Operator, "-")

	assert.Equal(t, GetAssoc(power), "right")
	assert.Equal(t, GetAssoc(multiply), "left")
	assert.Equal(t, GetAssoc(divide), "left")
	assert.Equal(t, GetAssoc(add), "left")
	assert.Equal(t, GetAssoc(subtract), "left")
}

func TestGetPrec(t *testing.T) {
	var power = NewToken(Operator, "^")
	var multiply = NewToken(Operator, "*")
	var divide = NewToken(Operator, "/")
	var add = NewToken(Operator, "+")
	var subtract = NewToken(Operator, "-")

	assert.Equal(t, GetPrec(power), 4)
	assert.Equal(t, GetPrec(multiply), 3)
	assert.Equal(t, GetPrec(divide), 3)
	assert.Equal(t, GetPrec(add), 2)
	assert.Equal(t, GetPrec(subtract), 2)
}
