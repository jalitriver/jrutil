package jrutil

// Stack holds a stack of values implemented as a slice.
type Stack[T comparable] struct {
	items []T
}

// NewStack returns a new stack.
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Push pushes the item onto the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop return the item at the top of the stack.  If the stack is
// empty, false is returned.
func (s *Stack[T]) Pop() (T, bool) {

	// Handle the error case where the stack is empty.
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	// Pop the last value from the slice.
	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return result, true
}

// Contains returns true if the stack contains the item; otherwise, it
// returns false.
func (s *Stack[T]) Contains(item T) bool {
	for _, x := range s.items {
		if x == item {
			return true
		}
	}
	return false
}
