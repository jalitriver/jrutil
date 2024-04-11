package jrutil

// Pair holds a pair of values.
type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

// MakePair returns a new Pair initialized with the first and second
// values.
func MakePair[T1 any, T2 any](first T1, second T2) Pair[T1, T2] {
	return Pair[T1, T2]{
		First:  first,
		Second: second,
	}
}

// MapToPairs converts the key/value pairs in the unordered input map
// m to an unordered slice of key/value pairs.  For example, if you
// want to sort the key/value pairs in a map by value, you can do the
// following:
//
//	// Create the map.
//	counts := map[string]uint64{
//		"foo": 10,
//		"bar": 3,
//		"baz": 7,
//	}
//
//	// Convert the map into a list of key/value pairs.
//	kvPairs := jrutil.MapToPairs(counts)
//
//	// Sort the key/value pairs.
//	sort.Slice(kvPairs, func (i, j int) bool {
//		return kvPairs[i].Second < kvPairs[j].Second
//	})
//
//	// Print the results.
//	for _, kvPair := range kvPairs {
//		fmt.Printf("%v: %v\n", kvPair.First, kvPair.Second)
//	}
func MapToPairs[T1 comparable, T2 any](m map[T1]T2) []Pair[T1, T2] {
	var result []Pair[T1, T2]
	for k, v := range m {
		result = append(result, MakePair(k, v))
	}
	return result
}
