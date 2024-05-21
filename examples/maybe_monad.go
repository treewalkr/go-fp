package examples

// Maybe represents an optional value; it could be Some (a value) or None (no value)
type Maybe[T any] struct {
	value *T
}

// Some creates a Maybe with a value
func Some[T any](v T) Maybe[T] {
	return Maybe[T]{value: &v}
}

// None creates a Maybe with no value
func None[T any]() Maybe[T] {
	return Maybe[T]{value: nil}
}

// Bind applies a function to the value inside Maybe, if it exists
func (m Maybe[T]) Bind(f func(T) Maybe[T]) Maybe[T] {
	if m.value == nil {
		return None[T]()
	}
	return f(*m.value)
}

// Return wraps a value in a Maybe (similar to Haskell's `return`)
func Return[T any](v T) Maybe[T] {
	return Some(v)
}

// IsSome checks if the Maybe contains a value
func (m Maybe[T]) IsSome() bool {
	return m.value != nil
}

// IsNone checks if the Maybe does not contain a value
func (m Maybe[T]) IsNone() bool {
	return m.value == nil
}

// Get retrieves the value from the Maybe, panicking if it is None
func (m Maybe[T]) Get() T {
	if m.value == nil {
		panic("Attempted to get value from None")
	}
	return *m.value
}

// GetOrElse retrieves the value if present, or returns a default if None
func (m Maybe[T]) GetOrElse(defaultValue T) T {
	if m.value == nil {
		return defaultValue
	}
	return *m.value
}
