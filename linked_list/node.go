package linkedlist

// Node represents a node in linked list
type Node struct {
	Value int
	Next  *Node
}

// NewNode initiates Node
func NewNode(value int, next *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}

// HasNext checkes if a node has a pointer of the next node
func (node *Node) HasNext() bool {
	return node.Next != nil
}
