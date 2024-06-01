package examples

import "fmt"

// Either represents a value of one of two types: Left or Right.
type Either[L, R any] struct {
	left  *L
	right *R
}

// Left creates an Either with a Left value.
func Left[L, R any](value L) Either[L, R] {
	return Either[L, R]{left: &value, right: nil}
}

// Right creates an Either with a Right value.
func Right[L, R any](value R) Either[L, R] {
	return Either[L, R]{left: nil, right: &value}
}

// IsLeft checks if the Either contains a Left value.
func (e Either[L, R]) IsLeft() bool {
	return e.left != nil
}

// IsRight checks if the Either contains a Right value.
func (e Either[L, R]) IsRight() bool {
	return e.right != nil
}

// GetLeft retrieves the Left value, if present.
func (e Either[L, R]) GetLeft() (L, bool) {
	if e.IsLeft() {
		return *e.left, true
	}
	var zero L
	return zero, false
}

// GetRight retrieves the Right value, if present.
func (e Either[L, R]) GetRight() (R, bool) {
	if e.IsRight() {
		return *e.right, true
	}
	var zero R
	return zero, false
}

// String implements a string representation for Either.
func (e Either[L, R]) String() string {
	if e.IsRight() {
		return fmt.Sprintf("Right{%v}", *e.right)
	}
	return fmt.Sprintf("Left{%v}", *e.left)
}
