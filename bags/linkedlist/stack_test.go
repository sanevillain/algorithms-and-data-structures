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
		s := NewStack[int]()

		size := 10
		for i := 0; i < size; i++ {
			s.Push(i)
		}

		if s.Size() != size {
			t.Errorf("s.Size() = %v, want %v", s.Size(), size)
		}
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("it should pop items from the stack in the reverse order of insertion", func(t *testing.T) {
		s := NewStack[int]()

		vals := []int{1, 2, 3, 4, 5}
		for _, v := range vals {
			s.Push(v)
		}

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
	t.Run("it should copy the stack and its items in the correct order", func(t *testing.T) {
		s := NewStack[int]()

		vals := []int{1, 2, 3, 4, 5}
		for _, v := range vals {
			s.Push(v)
		}

		s2 := s.Copy()

		for i := len(vals) - 1; i >= 0; i-- {
			if pop := s2.Pop(); pop != vals[i] {
				t.Errorf("s2.Pop() = %v, want %v", pop, vals[i])
			}
		}
	})
}

func TestStack_Delete(t *testing.T) {
	t.Run("it shouldnt impact stack if called with an index beyond the stack items", func(t *testing.T) {
		s := NewStack[int]()

		vals := []int{1, 2, 3, 4, 5}
		for _, v := range vals {
			s.Push(v)
		}

		s.Delete(-1)
		s.Delete(10)

		if s.Size() != len(vals) {
			t.Errorf("s.Size() = %v, want %v", s.Size(), len(vals))
		}
	})

	t.Run("it should delete the first item and set the correct head", func(t *testing.T) {
		s := NewStack[int]()

		item1 := 1
		item2 := 2
		item3 := 3

		s.Push(item1)
		s.Push(item2)
		s.Push(item3)

		s.Delete(0)

		if s.Size() != 2 {
			t.Errorf("s.Size() = %v, want %v", s.Size(), 2)
		}

		if s.Pop() != item2 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), item2)
		}

		if s.Pop() != item1 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), item1)
		}
	})

	t.Run("it should delete the second item and set the correct head", func(t *testing.T) {
		s := NewStack[int]()

		item1 := 1
		item2 := 2
		item3 := 3

		s.Push(item1)
		s.Push(item2)
		s.Push(item3)

		s.Delete(1)

		if s.Size() != 2 {
			t.Errorf("s.Size() = %v, want %v", s.Size(), 2)
		}

		if s.Pop() != item3 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), item3)
		}

		if s.Pop() != item1 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), item1)
		}
	})

	t.Run("it should delete the third item and set the correct head", func(t *testing.T) {
		s := NewStack[int]()

		item1 := 1
		item2 := 2
		item3 := 3

		s.Push(item1)
		s.Push(item2)
		s.Push(item3)

		s.Delete(2)

		if s.Size() != 2 {
			t.Errorf("s.Size() = %v, want %v", s.Size(), 2)
		}

		if s.Pop() != item3 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), item3)
		}

		if s.Pop() != item2 {
			t.Errorf("s.Peek() = %v, want %v", s.Peek(), item2)
		}
	})
}

func TestStack_Find(t *testing.T) {
	t.Run("it should return false if the item hasnt been pushed on to the stack", func(t *testing.T) {
		s := NewStack[int]()

		vals := []int{1, 2, 3, 4, 5}
		for _, v := range vals {
			if s.Find(v) {
				t.Errorf("s.Find(%v) = %v, want %v", v, s.Find(v), false)
			}
		}
	})

	t.Run("it should return true if the item has been pushed on to the stack", func(t *testing.T) {
		s := NewStack[int]()

		vals := []int{1, 2, 3, 4, 5}
		for _, v := range vals {
			s.Push(v)
		}

		for _, v := range vals {
			if !s.Find(v) {
				t.Errorf("s.Find(%v) = %v, want %v", v, s.Find(v), true)
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

func TestStack_Max(t *testing.T) {
	t.Run("it should return 0 if the stack is empty", func(t *testing.T) {
		s := NewStack[int]()

		if max := s.Max(); max != 0 {
			t.Errorf("s.Max() = %v, want %v", max, 0)
		}
	})

	t.Run("it should return max element", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(10)
		s.Push(0)
		s.Push(3)
		s.Push(22)
		s.Push(5)
		s.Push(50)

		if max := s.Max(); max != 50 {
			t.Errorf("s.Max() = %v, want %v", max, 50)
		}
	})
}

func TestStack_Reverse(t *testing.T) {
	t.Run("it should reverse the stack and return a new copy without modifying the underlying stack", func(t *testing.T) {
		s := NewStack[int]()
		for i := 0; i <= 10; i++ {
			s.Push(i)
		}

		rev := s.Reverse()
		for i := 0; i <= 10; i++ {
			if pop := rev.Pop(); pop != i {
				t.Errorf("rev.Pop() = %v, want %v", pop, i)
			}
		}

		for i := 10; i >= 0; i-- {
			if pop := s.Pop(); pop != i {
				t.Errorf("s.Pop() = %v, want %v", pop, i)
			}
		}
	})
}