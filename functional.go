package jrutil

// Map use the function f to map each element of the input slice xs to
// the corresponding element in the output slice.
func Map[T1 any, T2 any](xs []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(xs))
	for i, x := range xs {
		result[i] = f(x)
	}
	return result
}

// Reduce uses the function f to reduce the input slice xs to a single
// output value.  The accumulator is initialized with init.  f will be
// called once for each element in xs with the accumulator as the
// first argument and the element of xs as the second argument.  After
// calling f for all element of xs, the accumulator is returned.
func Reduce[T1 any, T2 any](xs []T1, init T2, f func(T2, T1) T2) T2 {
	acc := init
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

// Filter uses the function f to filter the input slice xs.  The
// return slice will hold only the values of xs for which f returned
// true.
func Filter[T any](xs []T, f func(x T) bool) []T {
	var result []T
	for _, x := range xs {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}
