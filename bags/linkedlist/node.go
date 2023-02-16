package linkedlist

import "golang.org/x/exp/constraints"

// Node is a node in a linked list that keeps a reference to the node after it
type Node[T constraints.Ordered] struct {
	item T        // item of the node
	next *Node[T] // next node in the list
}

// NewNode creates a new node.
func NewNode[T constraints.Ordered](item T) *Node[T] {
	return &Node[T]{item, nil}
}

// Item returns the item of the node.
func (n *Node[T]) Item() T {
	return n.item
}

// SetItem sets the item of the node.
func (n *Node[T]) SetItem(item T) {
	n.item = item
}

// SetNext sets the next field to point to the input node
func (n *Node[T]) SetNext(next *Node[T]) {
	n.next = next
}

// Next returns the next node.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// DoubleNode is a node in a linked list that keeps a references to both the node before and after it
type DoubleNode[T constraints.Ordered] struct {
	item T              // item of the node
	next *DoubleNode[T] // next node in the list
	prev *DoubleNode[T] // previous node in the list
}

// Item returns the item of the node.
func (n *DoubleNode[T]) Item() T {
	return n.item
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
	// removing a double node from a list requires linking the previous and next nodes to one another (skipping the node that we want to remove)

	if n.prev != nil {
		// always make sure to check for nil references before setting
		// the previous node's next field to point to the one after next
		n.prev.next = n.next
	}

	if n.next != nil {
		// always make sure to check for nil references before setting
		// the next node's previous field to point to the one before prev
		n.next.prev = n.prev
	}

	// remove the references to the previous and next nodes (optional)
	n.next = nil
	n.prev = nil

	// return the item it holds
	return n.item
}

// NewDoubleNode creates a new node.
func NewDoubleNode[T constraints.Ordered](item T) *DoubleNode[T] {
	return &DoubleNode[T]{item, nil, nil}
}
