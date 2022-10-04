package container

type Queue[T any] struct {
	base  []T
	first int
	next  int
}

func (q *Queue[T]) Enqueue(i T) {
	if len(q.base) == 0 {
		q.base = make([]T, 8) // 8 is arbitrary, any value > 2 would do
		q.first = -1
	}
	if q.next != q.first {
		if q.first == -1 {
			q.first = q.next
		}
		q.base[q.next] = i
		q.next = (q.next + 1) % cap(q.base)
		return
	}
	previousSize := cap(q.base)
	biggerBase := make([]T, previousSize*2)

	copy(biggerBase, q.base[q.first:previousSize])
	copy(biggerBase[previousSize-q.first:], q.base[:q.first])
	q.base = biggerBase
	q.first = 0
	q.base[previousSize] = i
	q.next = previousSize + 1
}

func (q *Queue[T]) Dequeue() T {
	size := q.Size()
	if size == 0 {
		panic("cannot dequeue from empty queue")
	}
	result := q.base[q.first]
	q.first = (q.first + 1) % cap(q.base)
	if q.first == q.next {
		q.first = -1
		q.next = 0
	}
	return result
}

func (q *Queue[T]) Peek() T {
	size := q.Size()
	if size == 0 {
		panic("cannot peek from empty queue")
	}
	return q.base[q.first]
}

func (q *Queue[T]) Size() int {
	if len(q.base) == 0 || q.first == -1 {
		return 0
	}
	if q.next == q.first {
		return cap(q.base)
	}
	return (q.next - q.first + cap(q.base)) % cap(q.base)
}
