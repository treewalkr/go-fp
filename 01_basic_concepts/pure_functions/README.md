# Pure Functions in Go

A **pure function** is fundamental in functional programming, defined by two primary properties:
1. It always produces the same output for the same input.
2. It has no side effects, meaning it does not modify external state or rely on it.

Pure functions bring significant advantages in programming, especially when building applications that prioritize predictability and ease of testing.

## Benefits of Pure Functions

- **Predictability**: Pure functions always behave consistently, making them easier to understand and debug.
- **Testability**: Since their outputs are deterministic, pure functions are straightforward to test.
- **Reusability**: With no dependencies on external state, pure functions can be used across different parts of a program or in other projects.

## Example: Pure Functions in Go

In Go, here are some simple examples of pure functions, which don’t rely on any external state and always produce the same output given the same input.

```go
// Pure function example: adding two numbers
func Add(a, b int) int {
	return a + b
}

// Pure function example: calculating factorial recursively
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
```
**Add**: The `Add` function is pure because it only depends on the inputs `a` and `b`, producing the same result every time for identical inputs.

**Factorial**: The `Factorial` function uses recursion to calculate the factorial of a number, maintaining purity by not modifying any external data.

Recursion is commonly used in functional programming to achieve results without mutable state or loops, making it an ideal technique for pure functions in Go.

## Pitfalls with Purity in Go
While pure functions are ideal, certain Go libraries can introduce impurities that affect predictability and side-effect-free behavior. Here are some common examples.

### Randomness
Functions that rely on random number generation, such as `rand.Intn`, are impure. Each call to `rand.Intn` produces a different value based on the current state of the random seed, resulting in unpredictability.

```go
// Impure example: returns a different value each time due to global randomness
func ImpureRandomNumber(limit int) int {
	return rand.Intn(limit)
}
```

### Pure Alternative: Deterministic Randomness
To make a function’s randomness predictable, we can use a seeded random generator instance. This approach isn’t purely random but allows for repeatable, consistent results when the same seed is used.

```go
// Pure version: deterministic randomness using a seed
func PureRandomNumber(seed int64, limit int) int {
	r := rand.New(rand.NewSource(seed))
	return r.Intn(limit)
}
```
By using a seed-based approach, the result remains consistent across calls with the same seed, helping maintain purity.

### Time
Functions that rely on the current time, such as `time.Now()`, introduce impurities because they produce different results based on the ever-changing system time.
```go
// Impure example: returns the current Unix timestamp, varying with each call
func ImpureTimestamp() int64 {
	return time.Now().Unix()
}
```
Functions that depend on `time.Now()` or similar time-based calls should be avoided in pure functional contexts, as they rely on external, mutable states that change unpredictably.

## Summary
Pure functions are key to predictable, reusable, and testable code. However, Go’s standard libraries include functions that can introduce impurities, particularly through randomness and time. By understanding these pitfalls and implementing alternatives, such as deterministic random generators, you can ensure your functions are as pure as possible.

Explore `pure_functions.go` for more examples of pure functions, their benefits, and how to avoid common pitfalls.