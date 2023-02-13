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

			if d.Size() != i+1 {
				t.Errorf("s.Size() = %v, want %v", d.Size(), i+1)
			}
		}

		if d.First().item != 0 {
			t.Errorf("d.First().item = %v, want %v", d.First().item, 0)
		}

		if d.Last().item != 9 {
			t.Errorf("d.Last().item = %v, want %v", d.Last().item, 9)
		}
	})
}

func TestQueue_Dequeue(t *testing.T) {
	t.Run("it should dequeue items in the same order of insertion", func(t *testing.T) {
		d := NewQueue[int]()

		vals := []int{1, 2, 3, 4, 5}
		for _, v := range vals {
			d.Enqueue(v)

			if d.Size() != v {
				t.Errorf("s.Size() = %v, want %v", d.Size(), v)
			}
		}

		for _, v := range vals {
			if deq := d.Dequeue(); deq != v {
				t.Errorf("d.Dequeue() = %v, want %v", deq, v)
			}
		}

		if !d.IsEmpty() {
			t.Errorf("d.IsEmpty() = %v, want %v", d.IsEmpty(), true)
		}
	})
}

func TestQueue_Copy(t *testing.T) {
	t.Run("it should copy the queue", func(t *testing.T) {
		d := NewQueue[int]()

		vals := []int{1, 2, 3, 4, 5}
		for _, v := range vals {
			d.Enqueue(v)
		}

		c := d.Copy()
		for !c.IsEmpty() {
			if deq, cDeq := d.Dequeue(), c.Dequeue(); deq != cDeq {
				t.Errorf("d.Dequeue() = %v, want %v", deq, cDeq)
			}
		}
	})
}
