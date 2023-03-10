package linkedlist

import "testing"

func TestStack_New(t *testing.T) {
	t.Run("it should create a new stack", func(t *testing.T) {
		s := NewStack[int]()

		if s == nil {
			t.Errorf("NewStack() = %v, want %v", s, 0)
		}

		if s.Size() != 0 {
			t.Errorf("NewStack() = %v, want %v", s.Size(), 0)
		}
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("it should push items to the stack", func(t *testing.T) {
		size := 10
		s, _ := newIntStack(size)

		if s.Size() != size {
			t.Errorf("s.Size() = %v, want %v", s.Size(), size)
		}
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("it should pop items from the stack in the reverse order of insertion", func(t *testing.T) {
		s, vals := newIntStack(5)

		for i := len(vals) - 1; i >= 0; i-- {
			if pop := s.Pop(); pop != vals[i] {
				t.Errorf("s.Pop() = %v, want %v", pop, vals[i])
			}
		}

		if !s.IsEmpty() {
			t.Errorf("s.IsEmpty() = %v, want %v", s.IsEmpty(), true)
		}
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("it should peek at the top item", func(t *testing.T) {
		s := NewStack[int]()

		item1 := 1
		s.Push(item1)

		if s.Peek() != item1 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), item1)
		}

		if s.Size() != 1 {
			t.Errorf("s.Size() = %v, want %v", s.Size(), 1)
		}

		s.Pop()
		item2 := 2
		s.Push(item2)

		if s.Peek() != item2 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), item2)
		}

		if s.Size() != 1 {
			t.Errorf("s.Size() = %v, want %v", s.Size(), 1)
		}
	})
}

func TestStack_Copy(t *testing.T) {
	t.Run("it should copy the stack and its items in the correct order without modifying the underlying stack", func(t *testing.T) {
		s, vals := newIntStack(5)
		s2 := s.Copy()

		if s.Size() != s2.Size() {
			t.Errorf("s.Size() = %v, s2.Size() = %v, want %v", s.Size(), s2.Size(), 5)
		}

		for i := len(vals) - 1; i >= 0; i-- {
			if pop, popCopy := s.Pop(), s2.Pop(); pop != vals[i] || popCopy != vals[i] {
				t.Errorf("s.Pop() = %v, s2.Pop() = %v, want %v", pop, popCopy, vals[i])
			}
		}

		if !s.IsEmpty() && !s2.IsEmpty() {
			t.Errorf("s.IsEmpty() = %v, s2.IsEmpty() = %v, want %v", s.IsEmpty(), s2.IsEmpty(), true)
		}
	})
}

func TestStack_Find(t *testing.T) {
	t.Run("it should return false if the item hasnt been pushed on to the stack", func(t *testing.T) {
		s := NewStack[int]()

		for i := 0; i < 10; i++ {
			if found := s.Find(i); found {
				t.Errorf("s.Find(%v) = %v, want %v", i, found, false)
			}
		}
	})

	t.Run("it should return true if the item has been pushed on to the stack", func(t *testing.T) {
		s, vals := newIntStack(5)

		for _, v := range vals {
			if found := s.Find(v); !found {
				t.Errorf("s.Find(%v) = %v, want %v", v, found, true)
			}
		}
	})
}

func TestStack_Remove(t *testing.T) {
	t.Run("it shouldnt impact stack if called with an item that hasnt been added to stack", func(t *testing.T) {
		s := NewStack[int]()
		s.Remove(1)

		if s.Size() != 0 {
			t.Errorf("s.Size() = %v, want %v", s.Size(), 0)
		}

		if s.Peek() != 0 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), 0)
		}
	})

	t.Run("it should remove all items added to the stack that match", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Push(1)
		s.Push(3)
		s.Push(1)
		s.Remove(1)

		if s.Size() != 2 {
			t.Errorf("s.Size() = %v, want %v", s.Size(), 2)
		}

		if s.Peek() != 3 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), 3)
		}
	})
}

func TestStack_Reverse(t *testing.T) {
	t.Run("it should reverse the stack and return a new copy without modifying the underlying stack", func(t *testing.T) {
		s, _ := newIntStack(10)

		rev := s.Reverse()

		if s.Size() != 10 && rev.Size() != 10 {
			t.Errorf("s.Size() = %v, rev.Size() = %v, want %v", s.Size(), rev.Size(), 10)
		}

		for i := 1; i <= 10; i++ {
			if pop := rev.Pop(); pop != i {
				t.Errorf("rev.Pop() = %v, want %v", pop, i)
			}
		}

		for i := 10; i >= 1; i-- {
			if pop := s.Pop(); pop != i {
				t.Errorf("s.Pop() = %v, want %v", pop, i)
			}
		}

		if !s.IsEmpty() && !rev.IsEmpty() {
			t.Errorf("s.IsEmpty() = %v, rev.IsEmpty() = %v, want %v", s.IsEmpty(), rev.IsEmpty(), true)
		}
	})
}

func newIntStack(n int) (*Stack[int], []int) {
	stack := NewStack[int]()
	vals := []int{}

	for i := 1; i <= n; i++ {
		stack.Push(i)
		vals = append(vals, i)
	}

	return stack, vals
}
