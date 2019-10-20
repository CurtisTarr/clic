package tokenizer

import (
	"internal/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	var tokenType = Literal
	var value = "45"

	var token = NewToken(tokenType, value)

	assert.Equal(t, token.tokenType, tokenType)
	assert.Equal(t, token.value, value)
}
