package stack

import "testing"

func TestResizingArrayStack_New(t *testing.T) {
	t.Run("should create a new stack", func(t *testing.T) {
		if r := NewResizingArrayStack[int](); r == nil {
			t.Errorf("NewResizingArrayStack() = %v, want %v", r, 0)
		}
	})
}

func TestResizingArrayStack_Push(t *testing.T) {
	t.Run("should add items to the stack", func(t *testing.T) {
		r := NewResizingArrayStack[int]()

		size := 10
		for i := 0; i < size; i++ {
			r.Push(i)
		}

		if s := r.Size(); s != size {
			t.Errorf("error: expected size to be 1, but got %d", s)
		}
	})
}

func TestResizingArrayStack_Pop(t *testing.T) {
	t.Run("should remove items from the stack in the reverse order of insertion", func(t *testing.T) {
		r := NewResizingArrayStack[int]()

		items := []int{1, 2, 3, 4, 5}
		for _, i := range items {
			r.Push(i)
		}

		for i := len(items) - 1; i >= 0; i-- {
			if p := r.Pop(); p != items[i] {
				t.Errorf("error: expected to pop %d, but got %d", items[i], p)
			}
		}
	})

	t.Run("popping an empty stack should return the default value for the stored type", func(t *testing.T) {
		r := NewResizingArrayStack[int]()

		for i := 0; i < 10; i++ {
			if p := r.Pop(); p != 0 {
				t.Errorf("error: expected to pop 0, but got %d", p)
			}
		}
	})

}
