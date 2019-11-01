package compiler

import (
	"fmt"
	"internal/parser"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Compile(operations *parser.Node) string {
	return calculateNode(operations).Value
}

func calculateNode(node *parser.Node) *parser.Node {
	node.RightNode = handleSubNode(node.RightNode)
	node.LeftNode = handleSubNode(node.LeftNode)

	node.Value = performFunction(node.Value, node.LeftNode.Value, node.RightNode.Value)
	return node
}

func handleSubNode(subNode *parser.Node) *parser.Node {
	if subNode != nil && isOperator(subNode.Value) {
		subNode = calculateNode(subNode)
	}
	return subNode
}

func performFunction(operator string, x string, y string) string {
	xF, _ := strconv.ParseFloat(strings.TrimSpace(x), 64)
	yF, _ := strconv.ParseFloat(strings.TrimSpace(y), 64)

	var result float64
	switch operator {
	case "^":
		result = math.Pow(xF, yF)
	case "*":
		result = xF * yF
	case "/":
		result = xF / yF
	case "+":
		result = xF + yF
	case "-":
		result = xF - yF
	}
	return fmt.Sprintf("%f", result)
}

func isOperator(token string) bool {
	var operator, _ = regexp.Compile(`[+\-*/^]`)
	var literal, _ = regexp.Compile(`[0-9]`)
	return operator.MatchString(token) && !literal.MatchString(token)
}