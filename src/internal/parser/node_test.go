package parser

import (
	"internal/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	var leftNode = NewNode("5", nil, nil)
	var rightNode = NewNode("10", nil, nil)
	var node = NewNode("*", leftNode, rightNode)

	assert.Equal(t, node.Value, "*")
	assert.Equal(t, node.LeftNode.Value, "5")
	assert.Equal(t, node.RightNode.Value, "10")
}