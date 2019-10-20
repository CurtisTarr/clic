package tokenizer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ToDo - error checking and handling long numbers and decimals

// Tokenize extracts all of the Tokens from the input string
func Tokenize(input string) []*Token {
	var cleanedInput = cleanInput(input)
	var tokenValues = getTokenValues(cleanedInput)

	var tokens []*Token
	for i, value := range tokenValues {
		if isLiteral(value) {
			tokens = append(tokens, NewToken(Literal, value))
		} else if isOperator(value) {
			tokens = append(tokens, NewToken(Operator, value))
		} else if isOpeningBrace(value) {
			tokens = append(tokens, NewToken(OpeningBrace, value))
		} else if isClosingBrace(value) {
			tokens = append(tokens, NewToken(ClosingBrace, value))
		} else {
			fmt.Printf("Unkown token detected [%s] at location [%d]", value, i)
		}
	}
	return tokens
}

func cleanInput(input string) string {
	return strings.Replace(input, " ", "", -1)
}

func getTokenValues(input string) []string {
	return strings.Split(input, "")
}

func isLiteral(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOperator(token string) bool {
	var operator, _ = regexp.Compile(`\+|-|\*|/`)
	return operator.MatchString(token)
}

func isOpeningBrace(token string) bool {
	return token == "("
}

func isClosingBrace(token string) bool {
	return token == ")"
}
