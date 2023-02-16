package linkedlist

import "golang.org/x/exp/constraints"

// Bag is a linked list implementation of a bag.
type Bag[T constraints.Ordered] struct {
	first *Node[T] // beginning of bag
}

func (b *Bag[T]) Add(item T) {
	// same as Stack.Push
	b.first = NewNode(item)
	b.first.SetNext(b.first)
}

func NewBag[T constraints.Ordered]() *Bag[T] {
	return &Bag[T]{}
}
