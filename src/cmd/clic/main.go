package main

import (
	"bufio"
	"fmt"
	"internal/calculator"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to clic! Enter a math equation (q quits): \n")

	var running bool = true
	for running {
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if input == "q" {
			running = false
			break
		}

		var output = calculator.Calculate(input)
		fmt.Printf("%s \n", output)
	}
}
