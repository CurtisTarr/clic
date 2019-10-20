package main

import (
	"fmt"
	"internal/compiler"
	"internal/parser"
	"internal/tokenizer"
	"strconv"
)

func main() {
	fmt.Print("Welcome to clic! Enter a math equation (q quits): \n")
	var running bool = true
	for running {
		var input string
		fmt.Scanln(&input)
		if input == "q" {
			running = false
			break
		}
		var tokens = tokenizer.Tokenize(input)
		var operations = parser.Parse(tokens)
		var output = compiler.Compile(operations)
		fmt.Println(strconv.ParseFloat(output) + "\n")
	}
}
