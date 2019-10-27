package parser

type Node struct {
	value     string
	leftNode  *Node
	rightNode *Node
}

func NewNode(value string, rightNode *Node, leftNode *Node) *Node {
	node := new(Node)
	node.value = value
	node.leftNode = leftNode
	node.rightNode = rightNode
	return node
}
