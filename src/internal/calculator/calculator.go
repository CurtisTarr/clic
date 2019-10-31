package calculator

import (
	"internal/parser"
	"internal/tokenizer"
)

func Calculate(input string) string {
	var tokens, err = tokenizer.Tokenize(input)
	if err != nil {
		return err.Error()
	}
	var operations = parser.Parse(tokens)

	return operations
	//return compiler.Compile(operations)
}
