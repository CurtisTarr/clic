package tokenizer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Tokenize(input string) (tokens []*Token, err error) {
	var cleanedInput = cleanInput(input)
	var tokenValues = getTokenValues(cleanedInput)
	var literalStr = ""

	for i, value := range tokenValues {
		if isLiteral(value) || isDecimal(value) {
			literalStr += value
		} else {
			if literalStr != "" {
				tokens = append(tokens, NewToken(Literal, literalStr))
				literalStr = ""
			}

			if isOperator(value) {
				tokens = append(tokens, NewToken(Operator, value))
			} else if isOpeningBrace(value) {
				tokens = append(tokens, NewToken(OpeningBrace, value))
			} else if isClosingBrace(value) {
				tokens = append(tokens, NewToken(ClosingBrace, value))
			} else {
				return nil, fmt.Errorf("unknown token detected [%s] at location [%d]", value, i)
			}
		}
	}

	if literalStr != "" {
		tokens = append(tokens, NewToken(Literal, literalStr))
	}
	return tokens, nil
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

func isDecimal(token string) bool {
	return token == "."
}

func isOperator(token string) bool {
	var operator, _ = regexp.Compile(`[+\-*/^]`)
	return operator.MatchString(token)
}

func isOpeningBrace(token string) bool {
	return token == "("
}

func isClosingBrace(token string) bool {
	return token == ")"
}
