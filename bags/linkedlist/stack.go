package linkedlist

import (
	"golang.org/x/exp/constraints"
)

// Stack is a linked list implementation of a stack.
type Stack[T constraints.Ordered] struct {
	first *Node[T] // top of stack (most recently added node)
	n     int      // number of items
}

// Push addNewDoubleNodeem to the stack.
func (s *Stack[T]) Push(item T) {
	s.first = NewNode(item, s.first)
	s.n++
}

// Peek returns the most recently added item. (without removing the node)
func (s *Stack[T]) Peek() T {
	if s.first == nil {
		var item T
		return item
	}

	return s.first.item
}

// Pop removes and returns the most recently added item.
func (s *Stack[T]) Pop() T {
	if s.first == nil {
		var item T
		return item
	}

	item := s.first.item
	s.first = s.first.next
	s.n--

	return item
}

// Copy returns a copy of the stack.
func (s *Stack[T]) Copy() *Stack[T] {
	stack := NewStack[T]()

	items := []T{}
	for node := s.first; node != nil; node = node.next {
		items = append(items, node.item)
	}

	for i := len(items) - 1; i >= 0; i-- {
		stack.Push(items[i])
	}

	return stack
}

// Delete removes the kth element in the list if it exists
func (s *Stack[T]) Delete(k int) {
	if s.IsEmpty() || k < 0 || k >= s.n {
		return
	}

	if k == 0 {
		s.Pop()
		return
	}

	for i, node := 0, s.first; node != nil; i, node = i+1, node.next {
		if i == k-1 && node.next != nil {
			node.next = node.next.next
			s.n--
			return
		}
	}
}

// Find returns true if the item exists in the stack
func (s *Stack[T]) Find(item T) bool {
	for node := s.first; node != nil; node = node.next {
		if node.item == item {
			return true
		}
	}

	return false
}

// Remove removes all of the item from the stack that have the input item
func (s *Stack[T]) Remove(item T) {
	if s.IsEmpty() {
		return
	}

	for prev, next := s.first, s.first.next; prev != nil && next != nil; {
		if next.item != item {
			prev, next = next, next.next
			continue
		}

		prev.next = next.next
		prev, next = prev.next, next.next
		s.n--
	}

	if s.first.item == item {
		s.first = s.first.next
		s.n--
	}
}

// Max returns the maximum value in the stack
func (s *Stack[T]) Max() T {
	var max T

	if s.IsEmpty() {
		return max
	}

	for node := s.first; node != nil; node = node.next {
		if node.item > max {
			max = node.item
		}
	}

	return max
}

// Reverse returns a reversed copy of the stack
func (s *Stack[T]) Reverse() *Stack[T] {
	stack := NewStack[T]()
	stackCopy := s.Copy()

	for !stackCopy.IsEmpty() {
		stack.Push(stackCopy.Pop())
	}

	return stack
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool { return s.first == nil }

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int { return s.n }

// NewStack creates a new stack.
func NewStack[T constraints.Ordered]() *Stack[T] { return &Stack[T]{} }
