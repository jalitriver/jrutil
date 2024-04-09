package jrutil

func MapToSlice[T1 comparable, T2 any](m map[T1]T2) []Pair[T1, T2] {
	var result []Pair[T1, T2]
	for k, v := range m {
		result = append(result, MakePair(k, v))
	}
	return result
}
