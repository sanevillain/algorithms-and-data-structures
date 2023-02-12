package linkedlist

import "golang.org/x/exp/constraints"

// Node is a node in a linked list
type Node[T constraints.Ordered] struct {
	item T
	next *Node[T]
}

// NewNode creates a new node.
func NewNode[T constraints.Ordered](item T, next *Node[T]) *Node[T] {
	return &Node[T]{item, next}
}

// Item returns the item of the node.
func (n *Node[T]) Item() T {
	return n.item
}

// Next returns the next node.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// DoubleNode is a node in a linked list
type DoubleNode[T constraints.Ordered] struct {
	item T
	next *DoubleNode[T]
	prev *DoubleNode[T]
}

// Item returns the item of the node.
func (n *DoubleNode[T]) Item() T {
	if n != nil {
		return n.item
	}

	var empty T
	return empty
}

// Next returns the next node.
func (n *DoubleNode[T]) Next() *DoubleNode[T] {
	return n.next
}

// Prev returns the previous node.
func (n *DoubleNode[T]) Prev() *DoubleNode[T] {
	return n.prev
}

// SetNext sets the next field to point to the input node
func (n *DoubleNode[T]) SetNext(next *DoubleNode[T]) {
	n.next = next
}

// SetPrev sets the previous field to point to the input node
func (n *DoubleNode[T]) SetPrev(prev *DoubleNode[T]) {
	n.prev = prev
}

// Remove removes the prev and next references and returns the item
func (n *DoubleNode[T]) Remove() T {
	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	n.next = nil
	n.prev = nil

	return n.item
}

// NewDoubleNode creates a new node.
func NewDoubleNode[T constraints.Ordered](item T, next, prev *DoubleNode[T]) *DoubleNode[T] {
	return &DoubleNode[T]{item, next, prev}
}
