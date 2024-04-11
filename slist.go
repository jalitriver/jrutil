//
// Immutable Single-Linked List (ala Common Lisp)
//
// - Length() is O(1) instead of O(N)
//
// - Can be faster than a slice when prepending elements or when
//   dealing with so many elements that slice re-allocation becomes a
//   bottleneck.
//

package jrutil

import (
	"fmt"
	"reflect"
	"strings"
)

// SList is a value and a recursive pointer to another list.  The
// elements in the list are constrained by type fmt.Stringer because
// we want to be able to convert the list to string which requires
// converting each element in the list to a string.
type SList[T fmt.Stringer] struct {
	value  T
	next   *SList[T]
	length uint64
}

// NewSList returns a new SList.
func NewSList[T fmt.Stringer]() *SList[T] {
	return nil
}

// Empty returns true if the list is empty; otherwise, it returns
// false.
func (l *SList[T]) Empty() bool {
	if l == nil {
		return true
	}
	return false
}

// PushFront appends the value to the front of the list.  This method
// is O(1) is generally used to build the list back-to-front.
func (l *SList[T]) PushFront(value T) *SList[T] {
	result := &SList[T]{
		value: value,
		next:  l,
	}
	if l == nil {
		result.length = 1
	} else {
		result.length = l.length + 1
	}
	return result
}

// Length returns the length of the list.  This method is O(1).
func (l *SList[T]) Length() uint64 {
	if l == nil {
		return 0
	}
	return l.length
}

// Head returns the value at the front of the list.
func (l *SList[T]) Head() (T, bool) {
	if l == nil {
		var zero T
		return zero, false
	}
	return l.value, true
}

// Tail returns the rest of the list which is the same as Drop(1).
func (l *SList[T]) Tail() *SList[T] {
	if l == nil {
		return nil
	}
	return l.next
}

// Reverse returns a new list that is the reverse of this list.  This
// method is often necessary, but it is O(N) and should be used
// sparingly.
func (l *SList[T]) Reverse() *SList[T] {

	result := NewSList[T]()

	for range l.Length() {
		result = result.PushFront(l.value)
		l = l.next
	}

	return result
}

// Drop removes the first n elements from the front of the list.
func (l *SList[T]) Drop(n uint64) *SList[T] {

	for range min(n, l.Length()) {
		l = l.next
	}

	return l
}

// DropUntil removes the elements from the head of the list until the
// predicate is true.
func (l *SList[T]) DropUntil(f func(x T) bool) *SList[T] {

	for range l.Length() {
		if f(l.value) {
			break
		}
		l = l.next
	}

	return l
}

// DropWhile removes elements from the head of the list while the
// predicate is true.
func (l *SList[T]) DropWhile(f func(x T) bool) *SList[T] {
	return l.DropUntil(func(x T) bool { return !f(x) })
}

// Take returns the first n elements.  This method is O(N) and should
// be used sparingly if n is large.  (Profiling shows this is
// implementation is faster (even for small lists) than a recursive
// implementation even though this implement requires a call to
// Reverse().)
func (l *SList[T]) Take(n uint64) *SList[T] {
	result := NewSList[T]()

	for range min(n, l.Length()) {
		result = result.PushFront(l.value)
		l = l.next
	}

	return result.Reverse()
}

// TakeWhile returns the elements from the head of the list while the
// predicate is true.  This method is O(N) and should be used
// sparingly if taking many elements.
func (l *SList[T]) TakeWhile(f func(x T) bool) *SList[T] {
	result := NewSList[T]()

	for range l.Length() {
		if !f(l.value) {
			break
		}
		result = result.PushFront(l.value)
		l = l.next
	}

	return result.Reverse()
}

// TakeUntil returns elements from the head of the list until the
// predicate is true.  This method is O(N) and should be used
// sparingly if taking many elements.
func (l *SList[T]) TakeUntil(f func(x T) bool) *SList[T] {
	return l.TakeWhile(func(x T) bool { return !f(x) })
}

// Contains returns true if the list contains an element that
// satisfies the predicate.
func (l *SList[T]) Contains(f func(x T) bool) bool {
	return l.DropUntil(f) != nil
}

// Nth returns the nth element (zero-based).
func (l *SList[T]) Nth(n uint64) (T, bool) {
	return l.Drop(n).Head()
}

// String returns the string representation of the list.
func (l *SList[T]) String() string {
	var b strings.Builder

	// Convert the value to display string.
	valueToString := func(value T) string {
		if reflect.TypeOf(value) == reflect.TypeOf("") {
			return fmt.Sprintf("%q", value)
		}
		return fmt.Sprintf("%v", value)
	}

	// Generate the string
	b.WriteString("[")
	for i := range l.Length() {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(valueToString(l.value))
		l = l.next
	}
	b.WriteString("]")

	return b.String()
}
