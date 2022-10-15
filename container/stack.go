package container

type Stack[T any] []T

func (s *Stack[T]) Push(i T) {
	*s = append(*s, i)
}

func (s *Stack[T]) PushAll(i ...T) {
	*s = append(*s, i...)
}

func (s *Stack[T]) Pop() T {
	l := len(*s)
	if l == 0 {
		panic("cannot pop from empty stack")
	}

	i := (*s)[l-1]
	*s = (*s)[:l-1]
	return i
}

func (s *Stack[T]) Peek() T {
	l := len(*s)
	if l == 0 {
		panic("cannot peek from empty stack")
	}

	return (*s)[l-1]
}

func (s *Stack[T]) Size() int {
	return len(*s)
}
