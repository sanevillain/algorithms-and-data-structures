package linkedlist

import "golang.org/x/exp/constraints"

// Queue is a linked list implementation of a queue.
type Queue[T constraints.Ordered] struct {
	first *Node[T] // beginning of queue
	last  *Node[T] // end of queue
	n     int      // number of items
}

// Enqueue adds an item to the end of the queue.
func (q *Queue[T]) Enqueue(item T) {
	// start by creating a new node
	node := NewNode(item)

	// increment the size of the queue
	q.n++

	if q.IsEmpty() {
		// queue is empty so the new node is first
		q.first = node
		// we don't have anything else to do for now
		return
	}

	if q.last == nil {
		// queue only has one node (the first) so set the next field of first node to the new node
		q.first.SetNext(node)
		// set the last node to the new node (our new last node)
		q.last = node
		// we don't have anything else to do for now
		return
	}

	// if we reach this part of the code we know that the queue has at least two nodes set (first and last)
	// so set the next field of the last node to the new node
	q.last.SetNext(node)
	// set the last node to the new node (our new last node)
	q.last = node
}

// Dequeue removes and returns the least recently added item.
func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		var item T
		// return the zero value of the type if queue is empty
		return item
	}

	// start by saving the item that the first node holds
	item := q.first.item
	// set the first field to the second node (first's next node)
	q.first = q.first.next

	// queue isn't empty so decrement the size of the queue
	q.n--

	if q.IsEmpty() {
		// we're done here
		return item
	}

	if q.first.next == nil {
		// if the new first node is the old last node then set the last node to nil
		q.last = nil
	}

	return item
}

func (q *Queue[T]) Copy() *Queue[T] {
	// start by creating a new queue
	queue := NewQueue[T]()

	// loop through the nodes in the queue and enqueue each them to the new queue until you reach the end node
	for node := q.first; node != nil; node = node.next {
		queue.Enqueue(node.item)
	}

	return queue
}

// First returns the least recently added node.
func (q *Queue[T]) First() *Node[T] {
	return q.first
}

// Last returns the most recently added node.
func (q *Queue[T]) Last() *Node[T] {
	return q.last
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.first == nil
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return q.n
}

// NewQueue creates a new queue.
func NewQueue[T constraints.Ordered]() *Queue[T] {
	return &Queue[T]{}
}
