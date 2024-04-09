package jrutil

func Map[T1 any, T2 any](xs []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(xs))
	for i, x := range xs {
		result[i] = f(x)
	}
	return result
}

func Reduce[T1 any, T2 any](xs []T1, init T2, f func(T2, T1) T2) T2 {
	result := init
	for _, x := range xs {
		result = f(result, x)
	}
	return result
}

func Filter[T any](xs []T, f func(x T) bool) []T {
	var result []T
	for _, x := range xs {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}
