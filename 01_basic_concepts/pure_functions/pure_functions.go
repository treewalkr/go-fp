package purefunctions

import (
	"math/rand"
	"time"
)

// Pure function example: add two numbers
func Add(a, b int) int {
	return a + b
}

// Pure function example: calculate the factorial of a number using recursion
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Pitfall: Functions with side effects (using randomness)
func ImpureRandomNumber(limit int) int {
	// Using rand.Intn() introduces impurity because it depends on a mutable global state (the seed).
	// Every time this function is called, it could return a different result.
	return rand.Intn(limit)
}

// Pitfall: Functions with side effects (using time)
func ImpureTimestamp() int64 {
	// Using time.Now().Unix() introduces impurity because it depends on the current time,
	// which is constantly changing.
	return time.Now().Unix()
}

// A pure alternative to randomness could be to use a seeded generator within the function,
// though it would still be impure if relying on changing seeds.
func PureRandomNumber(seed int64, limit int) int {
	// Create a new seeded generator instance, which ensures the result is predictable.
	r := rand.New(rand.NewSource(seed))
	return r.Intn(limit)
}
