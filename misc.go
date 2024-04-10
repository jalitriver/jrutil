package jrutil

import (
	crand "crypto/rand"
	"math/rand/v2"
)

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

// Return a slice of bytes initialzed from crypto.rand.Read().
func MakeRandomBytes(n uint64) ([]byte, error) {
	bs := make([]byte, 32)
	_, err := crand.Read(bs)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// Return a new math.rand.Rand (v2) instance that uses source that is
// seeded with random bytes from crypto.rand.Read().
func NewRand() (*rand.Rand, error) {
	bs, err := MakeRandomBytes(32)
	if err != nil {
		return nil, err
	}
	return rand.New(rand.NewChaCha8([32]byte(bs))), nil
}
