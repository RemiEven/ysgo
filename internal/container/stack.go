package container

// Stack represents a LIFO list.
type Stack[T any] []T

// Push adds a new element to the stack.
func (s *Stack[T]) Push(i T) {
	*s = append(*s, i)
}

// PushAll adds several new elements to the stack.
func (s *Stack[T]) PushAll(i ...T) {
	*s = append(*s, i...)
}

// Pop removes the last element of the stack and returns it (or panics if the stack is empty).
func (s *Stack[T]) Pop() T {
	l := len(*s)
	if l == 0 {
		panic("cannot pop from empty stack")
	}

	i := (*s)[l-1]
	*s = (*s)[:l-1]
	return i
}

// Peek returns the last element of the stack (or panics if the stack is empty).
func (s *Stack[T]) Peek() T {
	l := len(*s)
	if l == 0 {
		panic("cannot peek from empty stack")
	}

	return (*s)[l-1]
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(*s)
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	*s = (*s)[:0]
}
