package jrutil

import (
	"cmp"
	crand "crypto/rand"
	"math/rand/v2"
)

// MakePtr is useful for making pointers from intrinsic types like int
// for which the address-of operator does not work.
func MakePtr[T any](x T) *T {
	return &x
}

// IfElse is similar to C's ternary operator.
func IfElse[T any](b bool, consequent T, alternative T) T {
	if b {
		return consequent
	}
	return alternative
}

// MakeRandomBytes returns a slice of bytes initialzed from
// crypto.rand.Read().
func MakeRandomBytes(n uint64) ([]byte, error) {
	bs := make([]byte, 32)
	_, err := crand.Read(bs)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// NewRand returns a new math.rand.Rand (v2) instance that uses source
// that is seeded with random bytes from crypto.rand.Read().
func NewRand() (*rand.Rand, error) {
	bs, err := MakeRandomBytes(32)
	if err != nil {
		return nil, err
	}
	return rand.New(rand.NewChaCha8([32]byte(bs))), nil
}

// Merge merges the two sorted slices (xs and ys) and returns a new
// sorted slice that contains all of the elements from the two sorted
// input slices.  This function performs a shallow copy when copying
// values from xs and ys to the result.  Thus, if complex values are
// being sorted, xs and ys should hold pointers (for which shallow
// copying is correct).
func Merge[T cmp.Ordered](
	xs []T,
	ys []T,
	f func(x, y T) bool) []T {

	// Get the length of both slices.
	xLen := uint64(len(xs))
	yLen := uint64(len(ys))

	// Generate the slice to return.
	rLen := uint64(xLen + yLen)
	result := make([]T, rLen)

	// Merge the lists into the result.
	xIndex := uint64(0)
	yIndex := uint64(0)
	for i := uint64(0); i < rLen; i++ {

		// When we run out of xs, the next value must come from ys.
		if xIndex >= xLen {
			result[i] = ys[yIndex]
			yIndex++
			continue
		}

		// When we run out of ys, the next value must come from xs.
		if yIndex >= yLen {
			result[i] = xs[xIndex]
			xIndex++
			continue
		}

		// We still have values in both xs and ys.  Because xs and ys
		// are both sorted, the value at xIndex is the smallest value
		// in xs, and the value at yIndex is the smallest value in ys.
		// We just need to compare the two values at xIndex and yIndex
		// and copy the smaller to the result.
		if f(xs[xIndex], ys[yIndex]) {
			result[i] = xs[xIndex]
			xIndex++
		} else {
			result[i] = ys[yIndex]
			yIndex++
		}

	}

	return result
}

// MergeSort returns a new, sorted slice based on the comparison
// function f.
func MergeSort[T cmp.Ordered](
	xs []T,
	LessThanFn func(x1, x2 T) bool) []T {

	// Base case.
	if len(xs) <= 1 {
		return xs
	}

	// Divide and conquer.
	mid := uint64(len(xs) / 2)
	return Merge(
		MergeSort(xs[:mid], LessThanFn),
		MergeSort(xs[mid:], LessThanFn),
		LessThanFn)
}
