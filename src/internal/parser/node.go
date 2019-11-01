package parser

type Node struct {
	Value     string
	LeftNode  *Node
	RightNode *Node
}

func NewNode(value string, leftNode *Node, rightNode *Node) *Node {
	node := new(Node)
	node.Value = value
	node.LeftNode = leftNode
	node.RightNode = rightNode
	return node
}
