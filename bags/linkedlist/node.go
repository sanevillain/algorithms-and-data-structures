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
func (n *Node[T]) Item() T { return n.item }

// Next returns the next node.
func (n *Node[T]) Next() *Node[T] { return n.next }
