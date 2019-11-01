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

	var expectedOutput = NewNode("+",
		NewNode("/",
			NewNode("^",
				NewNode("^",
					NewNode("3", nil, nil),
					NewNode("2", nil, nil)),
				NewNode("-",
					NewNode("5", nil, nil),
					NewNode("1", nil, nil))),
			NewNode("*",
				NewNode("2", nil, nil),
				NewNode("4", nil, nil))),
		NewNode("3", nil, nil))
	var output = Parse(tokens)

	assert.DeepEqual(t, output, expectedOutput)
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

	var expectedOutput = NewNode("+",
		NewNode("/",
			NewNode("-",
				NewNode("3", nil, nil),
				NewNode("9", nil, nil)),
			NewNode("18", nil, nil)),
		NewNode("4", nil, nil))
	var output = Parse(tokens)

	assert.DeepEqual(t, output, expectedOutput)
}