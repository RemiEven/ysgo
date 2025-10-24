package container

import (
	"testing"

	"github.com/remieven/ysgo/internal/testutils"
)

func TestStackOneElement(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)

	if s.Peek() != 1 {
		t.Error("peek fail")
		return
	}

	if s.PeekFirst() != 1 {
		t.Error("peekFirst fail")
		return
	}

	if s.Pop() != 1 {
		t.Error("pop fail")
		return
	}
}

func TestStackManyElements(t *testing.T) {
	numberOfElements := 200 // arbitrary

	s := Stack[int]{}

	n := 0
	for n < numberOfElements {
		s.Push(n)
		n++
	}

	for n > 0 {
		if size := s.Size(); size != n {
			t.Errorf("unexpected size: wanted %d, got %d", n, size)
			return
		}

		if i := s.Peek(); i != n-1 {
			t.Errorf("unexpected peek: wanted %d, got %d", n-1, i)
			return
		}

		if i := s.PeekFirst(); i != 0 {
			t.Errorf("unexpected peekFirst: wanted %d, got %d", 0, i)
			return
		}

		if i := s.Pop(); i != n-1 {
			t.Errorf("unexpected pop: wanted %d, got %d", n-1, i)
			return
		}
		n--
	}
}

func TestEmptyStack(t *testing.T) {
	t.Run("Peek", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected a panic but got none")
				return
			}
			if expected := "cannot peek from empty stack"; r != expected {
				t.Errorf("unexpected panic value: wanted [%v], got [%v]", expected, r)
				return
			}
		}()

		s := Stack[int]{}
		s.Peek()
	})

	t.Run("Pop", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected a panic but got none")
				return
			}
			if expected := "cannot pop from empty stack"; r != expected {
				t.Errorf("unexpected panic value: wanted [%v], got [%v]", expected, r)
				return
			}
		}()

		s := Stack[int]{}
		s.Pop()
	})

	t.Run("PeekFirst", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected a panic but got none")
				return
			}
			if expected := "cannot peekFirst from empty stack"; r != expected {
				t.Errorf("unexpected panic value: wanted [%v], got [%v]", expected, r)
				return
			}
		}()

		s := Stack[int]{}
		s.PeekFirst()
	})
}

func TestStackPushAll(t *testing.T) {
	s := Stack[int]{}
	s.PushAll(8, 3)

	first := s.Pop()
	second := s.Pop()

	if first != 3 && second != 8 {
		t.Error("unexpected popped elements")
	}
}

func TestStackClear(t *testing.T) {
	s := Stack[int]{}
	s.PushAll(1, 2, 3)
	s.Clear()
	if s.Size() != 0 {
		t.Error("size should be zero after clearing stack")
	}
}

func TestNilStackCopyAsSlice(t *testing.T) {
	var s Stack[int]
	copiedSlice := s.CopyAsSlice()
	if copiedSlice == nil {
		t.Errorf("slice copy of nil stack should not be nil")
	} else if len(copiedSlice) != 0 {
		t.Errorf("slice copy of nil stack should be empty")
	}
}

func TestEmptyStackCopyAsSlice(t *testing.T) {
	s := Stack[int]{}
	copiedSlice := s.CopyAsSlice()
	if copiedSlice == nil {
		t.Errorf("slice copy of nil stack should not be nil")
	} else if len(copiedSlice) != 0 {
		t.Errorf("slice copy of nil stack should be empty")
	}
}

func TestStackCopyAsSlice(t *testing.T) {
	s := Stack[int]{1, 2, 3}
	copiedSlice := s.CopyAsSlice()
	expected := []int{1, 2, 3}
	if diff := testutils.DeepEqual(copiedSlice, expected); diff != "" {
		t.Error("unexpected value: " + diff)
		return
	}

	s.Pop()
	s.Push(4)
	if copiedSlice[2] == 4 {
		t.Error("slice copy of stack should not be affected by changes to the stack")
	}
}
