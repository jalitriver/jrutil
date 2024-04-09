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

type SList[T any] struct {
	value  T
	next   *SList[T]
	length uint64
}

func NewSList[T any]() *SList[T] {
	return nil
}

func (l *SList[T]) Empty() bool {
	if l == nil {
		return true
	}
	return false
}

// You have to build the list back-to-front.
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

func (l *SList[T]) Length() uint64 {
	if l == nil {
		return 0
	}
	return l.length
}

func (l *SList[T]) Head() (T, bool) {
	if l == nil {
		var zero T
		return zero, false
	}
	return l.value, true
}

func (l *SList[T]) Tail() *SList[T] {
	if l == nil {
		return nil
	}
	return l.next
}

// Reverse the list.
func (l *SList[T]) Reverse() *SList[T] {

	result := NewSList[T]()

	for range l.Length() {
		result = result.PushFront(l.value)
		l = l.next
	}

	return result
}

// Remove the first n elements from the head of the list.
func (l *SList[T]) Drop(n uint64) *SList[T] {

	for range min(n, l.Length()) {
		l = l.next
	}

	return l
}

// Remove the first elements from the head of the list until the
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

// Remove the first elements from the head of the list until the
// predicate is true.
func (l *SList[T]) DropWhile(f func(x T) bool) *SList[T] {
	return l.DropUntil(func(x T) bool {return !f(x)})
}

// Return the first n elements.  Profiling shows this is
// implementation is faster than a recursive implementation (which
// does not need to call Reverse()) even for small lists.
func (l *SList[T]) Take(n uint64) *SList[T] {
	result := NewSList[T]()

	for range min(n, l.Length()) {
		result = result.PushFront(l.value)
		l = l.next
	}

	return result.Reverse()
}

// Return the first elements from the head of the list while the
// predicate is true.
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

// Return the first elements from the head of the list until the
// predicate is true.
func (l *SList[T]) TakeUntil(f func(x T) bool) *SList[T] {
	return l.TakeWhile(func (x T) bool {return !f(x)})
}

// Return true if the list contains an element that satisfies the
// predicate.
func (l *SList[T]) Contains(f func(x T) bool) bool {
	return l.DropUntil(f) != nil
}

// Return the nth element (zero-based).
func (l *SList[T]) Nth(n uint64) (T, bool) {
	return l.Drop(n).Head()
}

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
