package jrutil

import (
)

type Stack[T comparable] struct {
	items []T
}

func New[T comparable]() *Stack[T] {
	return &Stack[T] {}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {

	// Handle the error case where the stack is empty.
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	// Pop the last value from the slice.
	result := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	return result, true
}

func (s *Stack[T]) Contains(item T) bool {
	for _, x := range s.items {
		if x == item {
			return true
		}
	}
	return false
}
