package calculator

import (
	"internal/compiler"
	"internal/parser"
	"internal/tokenizer"
)

func Calculate(input string) string {
	defer func() {
		if err := recover(); err != nil {
			println("Something went wrong, please check your expression")
		}
	}()

	var tokens, err = tokenizer.Tokenize(input)
	if err != nil {
		return err.Error()
	}
	var operations = parser.Parse(tokens)
	return compiler.Compile(operations)
}
