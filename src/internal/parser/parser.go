package parser

import (
	"internal/tokenizer"
)

func Parse(tokens []*tokenizer.Token) *Node {
	var queue []*Node
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
					queue = appendNodeToQueue(queue, topToken)
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
					queue = appendNodeToQueue(queue, topToken)
				} else {
					break
				}
			}
			stack = append(stack, token)
		default:
			queue = append(queue, NewNode(token.Value, nil, nil))
		}
	}
	for len(stack) > 0 {
		queue = appendNodeToQueue(queue, stack[len(stack)-1])
		stack = popStack(stack)
	}
	return queue[len(queue)-1]
}

func appendNodeToQueue(queue []*Node, token *tokenizer.Token) []*Node {
	var rightNode *Node
	var leftNode *Node
	rightNode, queue = queue[len(queue)-1], popQueue(queue)
	leftNode, queue = queue[len(queue)-1], popQueue(queue)

	var newNode = NewNode(token.Value, leftNode, rightNode)
	return append(queue, newNode)
}

func popQueue(queue []*Node) []*Node {
	if len(queue) > 0 {
		return queue[:len(queue)-1]
	}
	return queue
}

func popStack(stack []*tokenizer.Token) []*tokenizer.Token {
	if len(stack) > 0 {
		return stack[:len(stack)-1]
	}
	return stack
}
