package jrutil

type Vector[T any] []any

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
	for i := uint64(1) ; i < iMax ; i++ {
		v[vLen - i] = v[vLen - i - 1]
	}

	// Insert the new value.
	v[index] = value
	
	return v
}

func (v Vector[T]) PushBack(value T) Vector[T] {
	return append(v, value)
}

func (v Vector[T]) PushFront(value T) Vector[T] {
	return v.Insert(0, value)
}
