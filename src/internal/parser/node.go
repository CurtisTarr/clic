package parser

// Node object to represent node in operation tree
type Node struct {
	value     string
	leftNode  *Node
	rightNode *Node
}

// NewNode constructor for Node object
func NewNode(value string, rightNode *Node, leftNode *Node) *Node {
	node := new(Node)
	node.value = value
	node.leftNode = leftNode
	node.rightNode = rightNode
	return node
}
