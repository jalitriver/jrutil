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
	"strings"
)

// SList is a value and a recursive pointer to another list.
type SList[T any] struct {
	value  T
	next   *SList[T]
	length uint64
}

// NewSList returns a new SList.
func NewSList[T any]() *SList[T] {
	return nil
}

// NewSListFromSlice returns a new SList from the slice xs by
// performing a shallow copy on each element in xs.
func NewSListFromSlice[T any](xs []T) *SList[T] {
	result := NewSList[T]()
	for i := len(xs); i > 0; i-- {
		result = result.PushFront(xs[i-1])
	}
	return result
}

// ToSlice returns a new slice having the same elements as this list.
// The slice is created using shallow copies of the elements of this
// list.
func (l *SList[T]) ToSlice() []T {
	rLen := l.Length()
	result := make([]T, rLen)
	for i := uint64(0); (i < rLen) && (l != nil); i++ {
		result[i] = l.value
		l = l.next
	}
	return result
}

// Empty returns true if the list is empty; otherwise, it returns
// false.
func (l *SList[T]) Empty() bool {
	if l == nil {
		return true
	}
	return false
}

// Equal returns true if the two lists have the same length and if
// every element in this xs list is equal to the corresponding element
// in the ys list as determined by the isEqual function.
func (xs *SList[T]) Equal(
	ys *SList[T],
	isEqual func(x, y T) bool) bool {

	// Compare the lengths.
	if xs.Length() != ys.Length() {
		return false
	}

	// Compare each pair of elements.
	for (xs != nil) && (ys != nil) {
		if !isEqual(xs.value, ys.value) {
			return false
		}
		xs = xs.next
		ys = ys.next
	}

	return true
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

	for l != nil {
		result = result.PushFront(l.value)
		l = l.next
	}

	return result
}

// Drop removes the first n elements from the front of the list.  This
// method is relatively efficient because it does not have to allocate
// or deallocate any memory.
func (l *SList[T]) Drop(n uint64) *SList[T] {

	for i := uint64(0); (i < n) && (l != nil); i++ {
		l = l.next
	}

	return l
}

// DropUntil removes the elements from the head of the list until the
// predicate is true.
func (l *SList[T]) DropUntil(f func(x T) bool) *SList[T] {

	for l != nil {
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

// Take returns the first n elements.  This method is relatively
// inefficient because it has to build the new list that is returned.
// (Profiling shows this is implementation is faster (even for small
// lists) than a recursive implementation even though this implement
// requires a call to Reverse().)
func (l *SList[T]) Take(n uint64) *SList[T] {
	result := NewSList[T]()

	for i := uint64(0); (i < n) && (l != nil); i++ {
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

	for l != nil {
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
// satisfies the predicate.  Note that we cannot just pass in an
// element of type T to compare directly because T is constrained by
// "any" not "comparable".  This is done so SList works with more
// types.
func (l *SList[T]) Contains(f func(x T) bool) bool {
	return l.DropUntil(f) != nil
}

// Nth returns the nth element (zero-based).
func (l *SList[T]) Nth(n uint64) (T, bool) {
	return l.Drop(n).Head()
}

// Merge the sorted lists xs and ys into a new sorted list and return
// the result.
func (xs *SList[T]) Merge(
	ys *SList[T],
	isLessThan func(x, y T) bool) *SList[T] {

	// Generate the slice to return.
	result := NewSList[T]()

	// Merge the lists into the result.
	for (xs != nil) || (ys != nil) {

		// When we run out of xs, the next value must come from ys.
		if xs == nil {
			result = result.PushFront(ys.value)
			ys = ys.next
			continue
		}

		// When we run out of ys, the next value must come from xs.
		if ys == nil {
			result = result.PushFront(xs.value)
			xs = xs.next
			continue
		}

		// We still have values in both xs and ys.  Because xs and ys
		// are both sorted, the value at xsIndex is the smallest value
		// in xs, and the value at ysIndex is the smallest value in
		// ys.  We just need to compare the two values at xsIndex and
		// ysIndex and copy the smaller to the result.
		if isLessThan(xs.value, ys.value) {
			result = result.PushFront(xs.value)
			xs = xs.next
		} else {
			result = result.PushFront(ys.value)
			ys = ys.next
		}

	}

	return result.Reverse()
}

func (l *SList[T]) MergeSort(isLessThan func(l1, l2 T) bool) *SList[T] {

	// Base case.
	if l.Length() <= 1 {
		return l
	}

	// Divide and conquer.
	mid := l.Length() / 2
	xs := l.Take(mid).MergeSort(isLessThan)
	ys := l.Drop(mid).MergeSort(isLessThan)
	return xs.Merge(ys, isLessThan)
}

// String returns the string representation of the list.
func (l *SList[T]) String() string {
	var b strings.Builder

	// Generate the string
	b.WriteString("[")
	for curr := l; curr != nil; curr = curr.next {
		if curr != l {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%v", curr.value))
	}
	b.WriteString("]")

	return b.String()
}
