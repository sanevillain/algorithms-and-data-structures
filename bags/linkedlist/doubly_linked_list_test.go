package linkedlist

import "testing"

func TestDoublyLinkedList_Queue(t *testing.T) {
	t.Run("it should queue items", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Queue(1)
		list.Queue(2)
		list.Queue(3)

		if list.Size() != 3 {
			t.Errorf("expected size to be 3, got %d", list.Size())
		}
	})
}

func TestDoublyLinkedList_Stack(t *testing.T) {
	t.Run("it should stack items", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Stack(1)
		list.Stack(2)
		list.Stack(3)

		if list.Size() != 3 {
			t.Errorf("expected size to be 3, got %d", list.Size())
		}
	})
}

func TestDoublyLinkedList_First_Last(t *testing.T) {
	t.Run("it should return the first and last node", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Queue(1)
		list.Queue(2)
		list.Queue(3)

		if item := list.First().Item(); item != 1 {
			t.Errorf("expected first item to be 1, got %d", item)
		}

		if item := list.Last().Item(); item != 3 {
			t.Errorf("expected last item to be 3, got %d", item)
		}

		list.Clear()
		list.Stack(1)
		list.Stack(2)
		list.Stack(3)

		if item := list.First().Item(); item != 3 {
			t.Errorf("expected first item to be 3, got %d", item)
		}

		if item := list.Last().Item(); item != 1 {
			t.Errorf("expected last item to be 1, got %d", item)
		}
	})
}

func TestDoublyLinkedList_AddBefore(t *testing.T) {
	t.Run("it should add a node before another node", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Queue(1)

		list.AddBefore(1, 2)
		if first, last := list.First().Item(), list.Last().Item(); first != 2 || last != 1 {
			t.Errorf("expected first item to be 2 and last item to be 1, got %d and %d", first, last)
		}

		list.AddBefore(2, 3)
		if first, last := list.First().Item(), list.Last().Item(); first != 3 || last != 1 {
			t.Errorf("expected first item to be 3 and last item to be 1, got %d and %d", first, last)
		}
	})
}

func TestDoublyLinkedList_AddAfter(t *testing.T) {
	t.Run("it should add a node after another node", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Queue(1)

		list.AddAfter(1, 2)
		if first, last := list.First().Item(), list.Last().Item(); first != 1 || last != 2 {
			t.Errorf("expected first item to be 1 and last item to be 2, got %d and %d", first, last)
		}

		list.AddAfter(2, 3)
		if first, last := list.First().Item(), list.Last().Item(); first != 1 || last != 3 {
			t.Errorf("expected first item to be 1 and last item to be 3, got %d and %d", first, last)
		}

		list.AddAfter(3, 4)
		if first, last := list.First().Item(), list.Last().Item(); first != 1 || last != 4 {
			t.Errorf("expected first item to be 1 and last item to be 4, got %d and %d", first, last)
		}
	})
}

func TestDoublyLinkedList_Peek(t *testing.T) {
	t.Run("it should peek at the first item", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Queue(1)
		list.Queue(2)

		for i := 0; i < 10; i++ {
			if item := list.Peek(); item != 1 {
				t.Errorf("expected first item to be 1, got %d", item)
			}
		}
	})
}

func TestDoublyLinkedList_PeekTail(t *testing.T) {
	t.Run("it should peek at the last item", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Queue(1)
		list.Queue(2)

		for i := 0; i < 10; i++ {
			if item := list.PeekTail(); item != 2 {
				t.Errorf("expected first item to be 2, got %d", item)
			}
		}
	})
}

func TestDoublyLinkedList_Pop(t *testing.T) {
	t.Run("it should pop the first item (queue)", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Queue(1)
		list.Queue(2)

		if item := list.Pop(); item != 1 {
			t.Errorf("expected first item to be 1, got %d", item)
		}

		if item := list.Pop(); item != 2 {
			t.Errorf("expected first item to be 2, got %d", item)
		}

		list.Stack(1)
		list.Stack(2)

		if item := list.Pop(); item != 2 {
			t.Errorf("expected first item to be 2, got %d", item)
		}

		if item := list.Pop(); item != 1 {
			t.Errorf("expected first item to be 1, got %d", item)
		}
	})
}

func TestDoublyLinkedList_PopTail(t *testing.T) {
	t.Run("it should pop the last item (queue)", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Queue(1)
		list.Queue(2)
		list.Queue(3)

		if item := list.PopTail(); item != 3 {
			t.Errorf("expected first item to be 3, got %d", item)
		}

		if item := list.PopTail(); item != 2 {
			t.Errorf("expected first item to be 2, got %d", item)
		}

		list.Clear()
		list.Stack(1)
		list.Stack(2)
		list.Stack(3)

		if item := list.PopTail(); item != 1 {
			t.Errorf("expected first item to be 1, got %d", item)
		}

		if item := list.PopTail(); item != 2 {
			t.Errorf("expected first item to be 2, got %d", item)
		}
	})
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	t.Run("it should remove a node", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.Queue(1)
		list.Queue(2)
		list.Queue(3)
		list.Queue(4)
		list.Queue(5)

		list.Remove(3)
		if first, last, size := list.First().Item(), list.Last().Item(), list.Size(); first != 1 && last != 5 && size != 4 {
			t.Errorf("expected first item to be 1, last item to be 3, and size to be 2, got %d, %d, %d", first, last, size)
		}

		list.Remove(5)
		if first, last, size := list.First().Item(), list.Last().Item(), list.Size(); first != 1 && last != 4 && size != 3 {
			t.Errorf("expected first item to be 1, last item to be 4, and size to be 3, got %d, %d, %d", first, last, size)
		}

		list.Remove(1)
		if first, last, size := list.First().Item(), list.Last().Item(), list.Size(); first != 2 && last != 4 && size != 2 {
			t.Errorf("expected first item to be 2, last item to be 4, and size to be 2, got %d, %d, %d", first, last, size)
		}
	})
}
