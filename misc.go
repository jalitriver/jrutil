package jrutil

// Useful for making pointers to intrinsic types like int for which
// the address-of operator does not work.
func MakePtr[T any](x T) *T {
	return &x
}

// Poor man's ternary operator.
func IfElse[T any](b bool, consequent T, alternative T) T {
	if b {
		return consequent
	}
	return alternative
}
