package stack

// ResizingArrayStack is a generic stack (LIFO) that resizes its internal array when it is full.
type ResizingArrayStack[T any] struct {
	n      int // number of items
	values []T // the stack values
}

// Push adds an item to the stack.
func (s *ResizingArrayStack[T]) Push(item T) {
	// double the length of the array if it is full
	if s.n == len(s.values) {
		s.resize(s.n * 2)
	}

	s.values[s.n] = item
	s.n++
}

// Pop removes and returns the most recently added item.
func (s *ResizingArrayStack[T]) Pop() T {
	if s.IsEmpty() {
		var empty T
		return empty
	}

	s.n--
	val := s.values[s.n]
	s.values = s.values[:s.n]

	// reduce the length of the array if it is 1/4 full
	if l := len(s.values); s.n > 0 && s.n == l/4 {
		s.resize(l / 2)
	}

	return val
}

// resize resizes the stack to the given length.
func (s *ResizingArrayStack[T]) resize(length int) {
	values := make([]T, length, length)
	copy(values, s.values)
	s.values = values
}

// IsEmpty returns true if the stack is empty.
func (s *ResizingArrayStack[T]) IsEmpty() bool { return s.Size() == 0 }

// Size returns the number of items in the stack.
func (s *ResizingArrayStack[T]) Size() int { return s.n }

// NewResizingArrayStack returns a new ResizingArrayStack
func NewResizingArrayStack[T any]() *ResizingArrayStack[T] {
	return &ResizingArrayStack[T]{
		n:      0,
		values: make([]T, 1, 1),
	}
}
