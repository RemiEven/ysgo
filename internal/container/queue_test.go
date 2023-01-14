package container

import "testing"

func TestQueueOneElement(t *testing.T) {
	q := Queue[int]{}
	if actual := q.Size(); actual != 0 {
		t.Errorf("unexpected size before queueing element: wanted [%d], got [%d]", 0, actual)
		return
	}
	q.Enqueue(1)
	if actual := q.Size(); actual != 1 {
		t.Errorf("unexpected size after queueing element: wanted [%d], got [%d]", 1, actual)
		return
	}
	if actual := q.Peek(); actual != 1 {
		t.Errorf("unexpected peek result: wanted [%d], got [%d]", 1, actual)
		return
	}
	if actual := q.Size(); actual != 1 {
		t.Errorf("unexpected size after peeking element: wanted [%d], got [%d]", 1, actual)
		return
	}
	if actual := q.Dequeue(); actual != 1 {
		t.Errorf("unexpected dequeue result: wanted [%d], got [%d]", 1, actual)
		return
	}
	if actual := q.Size(); actual != 0 {
		t.Errorf("unexpected size after dequeueing element: wanted [%d], got [%d]", 0, actual)
		return
	}
}

func TestQueueManyElementsEnqueueThenDequeue(t *testing.T) {
	numberOfElements := 200 // arbitrary

	q := Queue[int]{}

	n := 0

	for n < numberOfElements {
		q.Enqueue(n)
		n++
	}

	for n > 0 {
		n--
		if size := q.Size(); size != n+1 {
			t.Errorf("unexpected size: wanted %d, got %d", n, size)
			return
		}

		expectedPeek := numberOfElements - n - 1
		if i := q.Peek(); i != expectedPeek {
			t.Errorf("unexpected peek: wanted %d, got %d", expectedPeek, i)
			return
		}

		if actual := q.Dequeue(); actual != expectedPeek {
			t.Errorf("unexpected dequeue: wanted %d, got %d", expectedPeek, actual)
			return
		}
	}
}

func TestQueueManyElementsMixedOperations(t *testing.T) {
	q := Queue[int]{}

	q.Enqueue(0)
	n := 1
	initialBaseCap := cap(q.base)
	for n < cap(q.base) {
		q.Enqueue(n)
		n++
	}
	i := 0
	const shift = 2
	for i < shift {
		q.Dequeue()
		i++
	}
	for i > 0 {
		q.Enqueue(n)
		n++
		i--
	}
	if baseCap := cap(q.base); baseCap != initialBaseCap {
		t.Errorf("unexpected base reallocation")
		return
	}

	q.Enqueue(n)

	if baseCap := cap(q.base); baseCap == initialBaseCap {
		t.Errorf("expected base reallocation but seems none happened")
		return
	}

	i = shift
	for i <= n {
		if actual := q.Dequeue(); actual != i {
			t.Errorf("unexpected dequeue: wanted %d, got %d", i, actual)
			return
		}
		i++
	}
}

func TestEmptyQueue(t *testing.T) {
	t.Run("Peek", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected a panic but got none")
				return
			}
			if expected := "cannot peek from empty queue"; r != expected {
				t.Errorf("unexpected panic value: wanted [%v], got [%v]", expected, r)
				return
			}
		}()

		q := Queue[string]{}
		q.Peek()
	})

	t.Run("Pop", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("expected a panic but got none")
				return
			}
			if expected := "cannot dequeue from empty queue"; r != expected {
				t.Errorf("unexpected panic value: wanted [%v], got [%v]", expected, r)
				return
			}
		}()

		q := Queue[string]{}
		q.Dequeue()
	})
}

func TestQueueFullCircle(t *testing.T) {
	q := Queue[bool]{}
	q.Enqueue(true)
	q.Dequeue()

	n := 0
	for n < cap(q.base) {
		q.Enqueue(true)
		n++
	}
	for n > 0 {
		q.Dequeue()
		n--
	}
	q.Enqueue(true)
	q.Dequeue()
}
