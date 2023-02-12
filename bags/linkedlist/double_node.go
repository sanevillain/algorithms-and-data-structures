package linkedlist

import "golang.org/x/exp/constraints"

// DoubleNode is a node in a linked list
type DoubleNode[T constraints.Ordered] struct {
	item T
	next *DoubleNode[T]
	prev *DoubleNode[T]
}

// Item returns the item of the node.
func (n *DoubleNode[T]) Item() T { return n.item }

// Next returns the next node.
func (n *DoubleNode[T]) Next() *DoubleNode[T] { return n.next }

// Prev returns the previous node.
func (n *DoubleNode[T]) Prev() *DoubleNode[T] { return n.prev }

// SetNext sets the next field to point to the input node
func (n *DoubleNode[T]) SetNext(next *DoubleNode[T]) {
	if n != nil {
		n.next = next
	}
}

// SetPrev sets the previous field to point to the input node
func (n *DoubleNode[T]) SetPrev(prev *DoubleNode[T]) {
	if n != nil {
		n.prev = prev
	}
}

// First returns the first node in the list.
func (n *DoubleNode[T]) First() *DoubleNode[T] {
	first := n
	for first.prev != nil {
		first = first.prev
	}

	return first
}

// Last returns the last node in the list.
func (n *DoubleNode[T]) Last() *DoubleNode[T] {
	last := n
	for last.next != nil {
		last = last.next
	}

	return last
}

// InsertBefore inserts a new node before the first node it points to
func (n *DoubleNode[T]) InsertAtBeginning(item T) {
	oldFirst := n.First()
	first := NewDoubleNode(item, oldFirst, nil)
	oldFirst.SetPrev(first)
}

// InsertAfter inserts a new node after the last node it points to
func (n *DoubleNode[T]) InsertAtEnd(item T) {
	oldLast := n.Last()
	last := NewDoubleNode(item, nil, oldLast)
	oldLast.SetNext(last)
}

// InsertBefore inserts a new node before the first match it finds
func (n *DoubleNode[T]) InsertBefore(match, item T) {
	if n.item == match {
		node := NewDoubleNode(item, n, n.prev)
		n.prev.SetNext(node)
		n.SetPrev(node)
		return
	}

	for node := n.First(); node != nil; node = node.next {
		if node.item == match {
			newNode := NewDoubleNode(item, node, node.prev)
			node.prev.SetNext(newNode)
			node.SetPrev(newNode)
			return
		}
	}
}

// InsertAfter inserts a new node after the first match it finds
func (n *DoubleNode[T]) InsertAfter(match, item T) {
	if n.item == match {
		node := NewDoubleNode(item, n.next, n)
		n.next.SetPrev(node)
		n.SetNext(node)
		return
	}

	for node := n.First(); node != nil; node = node.next {
		if node.item == match {
			newNode := NewDoubleNode(item, node.next, n)
			node.next.SetPrev(newNode)
			node.SetNext(newNode)
			return
		}
	}
}

// RemoveFirst removes the first node in the list
func (n *DoubleNode[T]) RemoveFirst() {
	if first := n.First(); first != nil {
		first.remove()
	}
}

// RemoveLast removes the last node in the list
func (n *DoubleNode[T]) RemoveLast() {
	if last := n.Last(); last != nil {
		last.remove()
	}
}

// Remove removes all the nodes that match the input item
func (n *DoubleNode[T]) Remove(item T) {
	for node := n.First(); node != nil; node = node.next {
		if node.item == item {
			node.remove()
		}
	}
}

// Remove removes the prev and next references
func (n *DoubleNode[T]) remove() {
	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	n.next = nil
	n.prev = nil
}

// NewDoubleNode creates a new node.
func NewDoubleNode[T constraints.Ordered](item T, next, prev *DoubleNode[T]) *DoubleNode[T] {
	return &DoubleNode[T]{item, next, prev}
}
