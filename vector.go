package jrutil

// Vector is a type alias for []any.
type Vector[T any] []any

// Insert inserts the value at the index.  If the index is greater
// than the length of the vector, the item is inserted at the end of
// the vector.  This method is O(N) and should be used sparingly.
func (v Vector[T]) Insert(index uint64, value T) Vector[T] {

	// Sanity check.  "index" should be no larger than one beyond the
	// end of the current slice.
	index = min(index, uint64(len(v)))

	// Grow the slice by one.
	var zero T
	v = append(v, zero)

	// Shift (which overwrites the "zero" that was just appended)
	vLen := uint64(len(v))
	iMax := vLen - index
	for i := uint64(1); i < iMax; i++ {
		v[vLen-i] = v[vLen-i-1]
	}

	// Insert the new value.
	v[index] = value

	return v
}

// PushBack appends the value to the end of the vector.
func (v Vector[T]) PushBack(value T) Vector[T] {
	return append(v, value)
}

// PushFront appends the value to the front of the vector.  This
// method is O(N) and should be used sparingly.
func (v Vector[T]) PushFront(value T) Vector[T] {
	return v.Insert(0, value)
}
