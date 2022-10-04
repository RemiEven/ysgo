package container

import "testing"

func TestStackOneElement(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)

	if s.Peek() != 1 {
		t.Error("peek fail")
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
}
