// my linked list implemented by myself and wikipedia
package queue

// defines the Node structure having just a value and the next item
type Node struct {
	Next  *Node
	Value *int
}

// NewNode creates a new node with determined value
func NewNode(value int) *Node {
	return &Node{
		Value: &value,
	}
}

// Add adds one item to the end of the list
func (n *Node) Add(newNode *Node) {
	curr := n
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

// Pop remove the last item from list and return it
func (n *Node) Pop() *Node {
	var prev *Node = nil
	curr := n
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	if prev != nil {
		prev.Next = nil
	}

	return curr
}
