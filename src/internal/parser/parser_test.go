package parser

import (
	"internal/assert"
	"internal/tokenizer"
	"testing"
)

func TestCase1(t *testing.T) {
	var tokens []*tokenizer.Token
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "3"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Operator, "+"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "4"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Operator, "*"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "2"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Operator, "/"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.OpeningBrace, "("))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "1"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Operator, "-"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "5"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.ClosingBrace, ")"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Operator, "^"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "2"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Operator, "^"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "3"))

	var output = Parse(tokens)

	assert.Equal(t, output, "3 4 2 * 1 5 - 2 3 ^ ^ / +")
}

func TestCase2(t *testing.T) {
	var tokens []*tokenizer.Token
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "4"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Operator, "+"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "18"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Operator, "/"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.OpeningBrace, "("))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "9"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Operator, "-"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.Literal, "3"))
	tokens = append(tokens, tokenizer.NewToken(tokenizer.ClosingBrace, ")"))

	var output = Parse(tokens)

	assert.Equal(t, output, "4 18 9 3 - / +")
}