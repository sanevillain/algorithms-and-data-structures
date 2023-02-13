package linkedlist

import "golang.org/x/exp/constraints"

// Queue is a linked list implementation of a queue.
type Queue[T constraints.Ordered] struct {
	first *Node[T] // beginning of queue
	last  *Node[T] // end of queue
	n     int      // number of items
}

// EnqueuNewDoubleNodean item to the queue.
func (q *Queue[T]) Enqueue(item T) {
	node := NewNode(item, nil)

	if q.IsEmpty() {
		q.first = node
	} else if q.last == nil {
		q.first.next = node
		q.last = node
	} else {
		q.last.next = node
		q.last = node
	}

	q.n++
}

// Dequeue removes and returns the least recently added item.
func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		var item T
		return item
	}

	item := q.first.item
	q.first = q.first.next

	if q.first == nil {
		q.last = nil
	} else {
		q.n--
	}

	return item
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
