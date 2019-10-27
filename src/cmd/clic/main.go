package main

import (
	"bufio"
	"fmt"
	"internal/calculator"
	"os"
)

func main() {
	fmt.Print("Welcome to clic! Enter a math equation (q quits): \n")
	var running bool = true
	for running {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		if input == "q" {
			running = false
			break
		}

		var output = calculator.Calculate(input)
		fmt.Printf("%s \n", output)
	}
}
