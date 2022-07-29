package queue

import (
	"sync"
)

// uses the nodes to manage the linked list, but uses mutex to concurrency
type Queue struct {
	mu   *sync.Mutex
	Node *Node
}

// instantiate a new queue
func NewQueue() *Queue {
	return &Queue{
		mu:   &sync.Mutex{},
		Node: nil,
	}
}

// Append a node to the queue with the value passed by param
func (q *Queue) Append(n int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Node == nil {
		q.Node = NewNode(n)
	} else {
		q.Node.Add(NewNode(n))
	}
}

// Pop remove the last item from queue (not using the return)
func (q *Queue) Pop() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Len() > 1 {
		q.Node.Pop()
	} else {
		q.Node = nil
	}
}

// Len counts how many nodes are in the list
func (q *Queue) Len() int {
	counter := 0
	if q.Node != nil {
		curr := q.Node
		counter++
		for curr.Next != nil {
			counter++
			curr = curr.Next
		}
	}

	return counter
}
