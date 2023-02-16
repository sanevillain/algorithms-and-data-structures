package linkedlist

import (
	"golang.org/x/exp/constraints"
)

// Stack is a linked list implementation of a stack.
type Stack[T constraints.Ordered] struct {
	first *Node[T] // top of stack (most recently added node)
	n     int      // number of items
}

// Push adds an item on top of the stack.
func (s *Stack[T]) Push(item T) {
	// create a new node
	node := NewNode(item)

	// set the next field of the new node to the first node (previous first)
	node.SetNext(s.first)

	// set the first node to the new node
	s.first = node

	// increment the number of items
	s.n++
}

// Pop removes and returns the most recently added item (top of the stack)
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var item T
		// return the zero value of the type
		return item
	}

	// save the top item
	item := s.first.item

	// set the first node to the second node
	s.first = s.first.next

	// decrement the number of items
	s.n--

	return item
}

// Peek returns the most recently added item. (without removing the node)
func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var item T
		// return the zero value of the type
		return item
	}

	return s.first.Item()
}

// Copy returns a copy of the stack.
func (s *Stack[T]) Copy() *Stack[T] {
	// create a new stack
	stack := NewStack[T]()

	// we need to keep track of the last node in the copy
	var last *Node[T]

	// loop through the stack until you reach the end node
	for node := s.first; node != nil; node = node.next {
		if stack.IsEmpty() {
			// push the first item onto the stack if it's empty
			stack.Push(node.item)
			// set the last node to the first node
			last = stack.first
			continue
		}

		// create a new node with the item of the current node
		newNode := NewNode(node.item)

		// set the next field of the last node to the new node
		last.SetNext(newNode)

		// set the last node to the new node
		last = newNode

		// increment the number of items
		stack.n++
	}

	return stack
}

// Remove removes all of the items from the stack that have the input item
func (s *Stack[T]) Remove(item T) {
	if s.IsEmpty() {
		// return early because we can't remove an item from an empty stack
		return
	}

	// loop through the stack until we reach the end node
	for prev, next := s.first, s.first.next; prev != nil && next != nil; prev, next = next, next.next {
		// if we find a match we need to remove the reference to the node
		if next.item == item {
			// set next to the node after next because we want to skip the current node in the next iteration
			next = next.next

			// set the next field of the previous node to the node after next (remove match reference)
			prev.next = next

			// decrement the number of items
			s.n--

			if next == nil {
				// break if we reach the end of the stack
				break
			}
		}
	}

	// check if the first item matches because in the loop we're skipping the first node
	if s.first.item == item {
		// remove the first item if it matches
		s.first = s.first.next

		// decrement the number of items
		s.n--
	}
}

// Find returns true if the item exists in the stack
func (s *Stack[T]) Find(item T) bool {
	// loop through the stack until we reach the end node
	for node := s.first; node != nil; node = node.next {
		if node.item == item {
			// return true if we find the item
			return true
		}
	}

	return false
}

// Reverse returns a reversed copy of the stack
func (s *Stack[T]) Reverse() *Stack[T] {
	// initialize a new stack
	stack := NewStack[T]()

	// loop through the stack until we reach the end node
	for node := s.first; node != nil; node = node.next {
		// push each item onto the stack
		stack.Push(node.item)
	}

	// stack is now in reverse order because we started from the first item
	// and inserted each following item at the top of the stack
	return stack
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return s.first == nil
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	return s.n
}

// NewStack creates a new stack.
func NewStack[T constraints.Ordered]() *Stack[T] {
	return &Stack[T]{}
}
