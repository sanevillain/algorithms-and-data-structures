package linkedlist

import "testing"

func TestQueue_New(t *testing.T) {
	t.Run("it should create a new empty queue", func(t *testing.T) {
		d := NewQueue[int]()

		if d == nil {
			t.Errorf("NewQueue() = %v, want %v", d, 0)
		}

		if empty := d.IsEmpty(); !empty {
			t.Errorf("d.IsEmpty() = %v, want %v", empty, true)
		}

		if size := d.Size(); size != 0 {
			t.Errorf("d.Size() = %v, want %v", size, 0)
		}

		if first, last := d.First(), d.Last(); first != nil || last != nil {
			t.Errorf("d.First() = %v, d.Last() = %v, want %v, %v", first, last, nil, nil)
		}
	})
}

func TestQueue_Enqueue(t *testing.T) {
	t.Run("it should queue items", func(t *testing.T) {
		d := NewQueue[int]()

		start, end := 1, 10
		for i := start; i <= end; i++ {
			d.Enqueue(i)

			if size := d.Size(); size != i {
				t.Errorf("d.Size() = %v, want %v", size, i)
			}
		}

		if empty := d.IsEmpty(); empty {
			t.Errorf("d.IsEmpty() = %v, want %v", empty, false)
		}

		if first, last := d.First(), d.Last(); first.item != start || last.item != end {
			t.Errorf("d.First() = %v, d.Last() = %v, want %v, %v", first.item, last.item, start, end)
		}
	})
}

func TestQueue_Dequeue(t *testing.T) {
	t.Run("it should dequeue items in the same order of insertion", func(t *testing.T) {
		d := NewQueue[int]()

		start, end := 1, 10
		for i := start; i <= end; i++ {
			d.Enqueue(i)
		}

		for i := start; i <= end; i++ {
			if size := d.Size(); size != end-i+1 {
				t.Errorf("d.Size() = %v, want %v", size, end-i+1)
			}

			if deq := d.Dequeue(); deq != i {
				t.Errorf("d.Dequeue() = %v, want %v", deq, i)
			}
		}

		if empty := d.IsEmpty(); !empty {
			t.Errorf("d.IsEmpty() = %v, want %v", empty, true)
		}

		if first, last := d.First(), d.Last(); first != nil || last != nil {
			t.Errorf("d.First() = %v, d.Last() = %v, want %v, %v", first, last, nil, nil)
		}
	})
}

func TestQueue_Copy(t *testing.T) {
	t.Run("it should copy the queue", func(t *testing.T) {
		d := NewQueue[int]()

		size := 10
		for i := 1; i <= size; i++ {
			d.Enqueue(i)
		}

		c := d.Copy()

		if empty := d.IsEmpty(); empty {
			t.Errorf("d.IsEmpty() = %v, want %v", empty, false)
		}

		for !c.IsEmpty() {
			if deq, cDeq := d.Dequeue(), c.Dequeue(); deq != cDeq {
				t.Errorf("d.Dequeue() = %v, want %v", deq, cDeq)
			}
		}
	})
}
