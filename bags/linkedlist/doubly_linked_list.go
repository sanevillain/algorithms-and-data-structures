package linkedlist

import (
	"golang.org/x/exp/constraints"
)

// DoublyLinkedList is a doubly linked list
type DoublyLinkedList[T constraints.Ordered] struct {
	first *DoubleNode[T]
	last  *DoubleNode[T]
	size  int
}

// NewDoublyLinkedList returns a new doubly linked list
func NewDoublyLinkedList[T constraints.Ordered]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// First returns the first node in the list.
func (l *DoublyLinkedList[T]) First() *DoubleNode[T] {
	return l.first
}

// Last returns the last node in the list.
func (l *DoublyLinkedList[T]) Last() *DoubleNode[T] {
	return l.last
}

// Size returns the size of the list.
func (l *DoublyLinkedList[T]) Size() int {
	return l.size
}

// IsEmpty returns true if the list is empty.
func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Stack add a new node to the beginning of the list
func (l *DoublyLinkedList[T]) Stack(item T) {
	oldFirst := l.First()
	l.first = NewDoubleNode(item, oldFirst, nil)

	if oldFirst != nil {
		oldFirst.SetPrev(l.first)

		if l.last == nil {
			l.last = oldFirst
		}
	}

	l.size++
}

// Queue adds a new node to the end of the list
func (l *DoublyLinkedList[T]) Queue(item T) {
	if l.IsEmpty() {
		l.first = NewDoubleNode(item, nil, nil)
		l.size++
		return
	}

	oldLast := l.Last()
	l.last = NewDoubleNode(item, nil, oldLast)

	if oldLast != nil {
		oldLast.SetNext(l.last)
	} else {
		l.first.SetNext(l.last)
	}

	l.size++
}

// AddBefore adds a new node before the first match it finds
func (l *DoublyLinkedList[T]) AddBefore(match, item T) {
	for node := l.First(); node != nil; node = node.next {
		if node.item != match {
			continue
		}

		newNode := NewDoubleNode(item, node, node.prev)
		node.SetPrev(newNode)

		if node.prev != nil {
			node.prev.SetNext(newNode)
		}

		if l.first == node {
			l.first = newNode
		}

		if l.last == nil {
			l.last = node
		}

		l.size++
		return
	}
}

// AddAfter inserts a new node after the first match it finds
func (l *DoublyLinkedList[T]) AddAfter(match, item T) {
	for node := l.First(); node != nil; node = node.next {
		if node.item != match {
			continue
		}

		newNode := NewDoubleNode(item, node.next, node)
		node.SetNext(newNode)

		if node.next != nil {
			node.next.SetPrev(newNode)
		}

		if l.last == node {
			l.last = newNode
		}

		l.size++
		return
	}
}

// Peek returns the first item in the list without removing it
func (l *DoublyLinkedList[T]) Peek() T {
	return l.First().Item()
}

// PeekTail returns the old item in the list without removing it
func (l *DoublyLinkedList[T]) PeekTail() T {
	return l.Last().Item()
}

// Pop removes the first node in the list
func (l *DoublyLinkedList[T]) Pop() T {
	if first := l.First(); first != nil {
		l.first = first.next
		l.size--
		return first.Remove()
	}

	var empty T
	return empty
}

// PopTail removes the last node in the list
func (l *DoublyLinkedList[T]) PopTail() T {
	if last := l.Last(); last != nil {
		l.last = last.prev
		l.size--
		return last.Remove()
	}

	var empty T
	return empty
}

// Remove removes the first node it finds as a match
func (l *DoublyLinkedList[T]) Remove(item T) T {
	for node := l.First(); node != nil; node = node.next {
		if node.item != item {
			continue
		}

		l.size--

		if node == l.first {
			return l.Pop()
		}

		if node == l.last {
			return l.PopTail()
		}

		return node.Remove()
	}

	var empty T
	return empty
}

// Clear removes all nodes from the list.
func (l *DoublyLinkedList[T]) Clear() {
	for node := l.First(); node != nil; node = node.next {
		node.Remove()
	}

	l.first = nil
	l.last = nil
	l.size = 0
}
