package parser

import (
	"internal/assert"
	"internal/tokenizer"
	"testing"
)

func TestParser(t *testing.T) {
	var tokens []*tokenizer.Token
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "5"))

	var output = Parse(tokens)

	assert.Equal(t, output, 10.)
}
