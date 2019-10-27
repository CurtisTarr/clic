package parser

import (
	"internal/tokenizer"
)

func Parse(tokens []*tokenizer.Token) (out string) {
	var stack []*tokenizer.Token
	for _, token := range tokens {
		switch token.TokenType {
		case tokenizer.OpeningBrace:
			stack = append(stack, token)
		case tokenizer.ClosingBrace:
			var topToken *tokenizer.Token
			for {
				if len(stack) > 0 {
					topToken, stack = stack[len(stack)-1], popStack(stack)
					if topToken.TokenType == tokenizer.OpeningBrace {
						break
					}
					out += " " + topToken.Value
				} else {
					break
				}
			}
		case tokenizer.Operator:
			for len(stack) > 0 {
				topToken := stack[len(stack)-1]
				if topToken.TokenType != tokenizer.OpeningBrace &&
					(tokenizer.GetPrec(token) < tokenizer.GetPrec(topToken) ||
						(tokenizer.GetPrec(token) == tokenizer.GetPrec(topToken) &&
							tokenizer.GetAssoc(topToken) == "left")) {
					stack = popStack(stack)
					out += " " + topToken.Value
				} else {
					break
				}
			}
			stack = append(stack, token)
		default:
			if out != "" {
				out += " "
			}
			out += token.Value
		}
	}
	for len(stack) > 0 {
		out += " " + stack[len(stack)-1].Value
		stack = popStack(stack)
	}
	return
}

func popStack(stack []*tokenizer.Token) []*tokenizer.Token {
	if len(stack) > 0 {
		return stack[:len(stack)-1]
	}
	return stack
}
