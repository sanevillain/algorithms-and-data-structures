package linkedlist

import "testing"

func TestQueue_New(t *testing.T) {
	t.Run("it should create a new queue", func(t *testing.T) {
		d := NewQueue[int]()

		if d == nil {
			t.Errorf("NewQueue() = %v, want %v", d, 0)
		}

		if d.Size() != 0 {
			t.Errorf("NewQueue() = %v, want %v", d.Size(), 0)
		}
	})
}

func TestQueue_Enqueue(t *testing.T) {
	t.Run("it should queue items", func(t *testing.T) {
		d := NewQueue[int]()

		size := 10
		for i := 0; i < size; i++ {
			d.Enqueue(i)
		}

		if d.Size() != size {
			t.Errorf("s.Size() = %v, want %v", d.Size(), size)
		}
	})
}

func TestQueue_Dequeue(t *testing.T) {
	t.Run("it should dequeue items in the same order of insertion", func(t *testing.T) {
		d := NewQueue[int]()

		vals := []int{1, 2, 3, 4, 5}
		for _, v := range vals {
			d.Enqueue(v)
		}

		for _, v := range vals {
			if d.Dequeue() != v {
				t.Errorf("d.Dequeue() = %v, want %v", d.Dequeue(), v)
			}
		}

		if !d.IsEmpty() {
			t.Errorf("d.IsEmpty() = %v, want %v", d.IsEmpty(), true)
		}
	})
}
