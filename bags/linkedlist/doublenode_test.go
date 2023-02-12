package linkedlist

import "testing"

func TestNode_New(t *testing.T) {
	t.Run("it should create a new node", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)

		if d == nil {
			t.Errorf("NewNode() = %v, want %v", d, 0)
		}
	})
}

func TestNode_SetNext(t *testing.T) {
	t.Run("it should set the next node", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d.SetNext(d1)

		if next := d.Next(); next != d1 {
			t.Errorf("d.Next() = %v, want %v", next, d1)
		}
	})
}

func TestNode_SetPrev(t *testing.T) {
	t.Run("it should set the previous node", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d1.SetPrev(d)

		if prev := d1.Prev(); prev != d {
			t.Errorf("d.Prev() = %v, want %v", prev, d)
		}
	})
}

func TestNode_First(t *testing.T) {
	t.Run("it should return the first node in the list", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d2 := NewDoubleNode(2, nil, nil)

		d.SetNext(d1)
		d1.SetPrev(d)
		d1.SetNext(d2)
		d2.SetPrev(d1)

		if first := d.First(); first != d {
			t.Errorf("d.First() = %v, want %v", first, d)
		}

		if first := d1.First(); first != d {
			t.Errorf("d1.First() = %v, want %v", first, d)
		}

		if first := d2.First(); first != d {
			t.Errorf("d2.First() = %v, want %v", first, d)
		}
	})
}

func TestNode_Last(t *testing.T) {
	t.Run("it should return the last node in the list", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d2 := NewDoubleNode(2, nil, nil)

		d.SetNext(d1)
		d1.SetPrev(d)
		d1.SetNext(d2)
		d2.SetPrev(d1)

		if last := d.Last(); last != d2 {
			t.Errorf("d.Last() = %v, want %v", last, d2)
		}

		if last := d1.Last(); last != d2 {
			t.Errorf("d1.Last() = %v, want %v", last, d2)
		}

		if last := d2.Last(); last != d2 {
			t.Errorf("d2.Last() = %v, want %v", last, d2)
		}
	})
}

func TestNode_InsertAtBeginning(t *testing.T) {
	t.Run("it should insert a new node at the beginning of the list", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d2 := NewDoubleNode(2, nil, nil)

		d.SetNext(d1)
		d1.SetPrev(d)
		d1.SetNext(d2)
		d2.SetPrev(d1)

		d.InsertAtBeginning(3)

		if item := d.First().Item(); item != 3 {
			t.Errorf("d.First().Item() = %v, want %v", item, 3)
		}
	})
}

func TestNode_InsertAtEnd(t *testing.T) {
	t.Run("it should insert a new node at the end of the list", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d2 := NewDoubleNode(2, nil, nil)

		d.SetNext(d1)
		d1.SetPrev(d)
		d1.SetNext(d2)
		d2.SetPrev(d1)

		d.InsertAtEnd(3)

		if item := d.Last().Item(); item != 3 {
			t.Errorf("d.Last().Item() = %v, want %v", item, 3)
		}
	})
}

func TestNode_InsertBefore(t *testing.T) {
	t.Run("it should insert a new node before the current node", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d2 := NewDoubleNode(2, nil, nil)

		d.SetNext(d1)
		d1.SetPrev(d)
		d1.SetNext(d2)
		d2.SetPrev(d1)

		d2.InsertBefore(0, 5)

		if item := d.First().Item(); item != 5 {
			t.Errorf("d.First().Item() = %v, want %v", item, 5)
		}
	})
}

func TestNode_InsertAfter(t *testing.T) {
	t.Run("it should insert a new node after the current node", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d2 := NewDoubleNode(2, nil, nil)

		d.SetNext(d1)
		d1.SetPrev(d)
		d1.SetNext(d2)
		d2.SetPrev(d1)

		d.InsertAfter(2, 5)

		if item := d.Last().Item(); item != 5 {
			t.Errorf("d.First().Item() = %v, want %v", item, 5)
		}
	})
}

func TestNode_RemoveFirst(t *testing.T) {
	t.Run("it should remove the first node in the list", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d2 := NewDoubleNode(2, nil, nil)

		d.SetNext(d1)
		d1.SetPrev(d)
		d1.SetNext(d2)
		d2.SetPrev(d1)

		d2.RemoveFirst()

		if first := d2.First(); first != d1 {
			t.Errorf("d.First() = %v, want %v", first, d1)
		}
	})
}

func TestNode_RemoveLast(t *testing.T) {
	t.Run("it should remove the last node in the list", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d2 := NewDoubleNode(2, nil, nil)

		d.SetNext(d1)
		d1.SetPrev(d)
		d1.SetNext(d2)
		d2.SetPrev(d1)

		d.RemoveLast()

		if last := d.Last(); last != d1 {
			t.Errorf("d.Last() = %v, want %v", last, d1)
		}
	})
}

func TestNode_Remove(t *testing.T) {
	t.Run("it should remove all the matching nodes from the list", func(t *testing.T) {
		d := NewDoubleNode(0, nil, nil)
		d1 := NewDoubleNode(1, nil, nil)
		d2 := NewDoubleNode(2, nil, nil)
		d3 := NewDoubleNode(1, nil, nil)
		d4 := NewDoubleNode(1, nil, nil)

		d.SetNext(d1)
		d1.SetPrev(d)
		d1.SetNext(d2)
		d2.SetPrev(d1)
		d3.SetPrev(d2)
		d3.SetNext(d4)
		d4.SetPrev(d3)

		d.Remove(1)
		d.Remove(2)

		if first, last := d.First(), d.Last(); first != last && last != d {
			t.Errorf("d.First() = %v, d.Last() = %v, want %v", first, last, d)
		}
	})
}
