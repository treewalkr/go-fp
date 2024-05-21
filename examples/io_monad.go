package examples

import (
	"bufio"
	"fmt"
	"os"
)

// IO represents an IO action that can be performed
type IO[T any] struct {
	run func() T
}

// New creates a new IO monad instance
func New[T any](run func() T) IO[T] {
	return IO[T]{run: run}
}

// Run executes the IO action and returns the result
func (io IO[T]) Run() T {
	return io.run()
}

// Bind allows us to chain IO operations (similar to `flatMap`)
func (io IO[T]) Bind(f func(T) IO[any]) IO[any] {
	return New(func() any {
		return f(io.Run()).Run()
	})
}

// Bind chains two IO actions, transforming T to another IO[R]
func Bind[T, R any](io IO[T], f func(T) IO[R]) IO[R] {
	return New(func() R {
		return f(io.Run()).Run()
	})
}

// Map transforms the result of an IO action from T to R
func Map[T, R any](io IO[T], f func(T) R) IO[R] {
	return New(func() R {
		return f(io.Run())
	})
}

// PrintIO is an IO action that prints to the console
func PrintIO(message string) IO[struct{}] {
	return New(func() struct{} {
		fmt.Println(message)
		return struct{}{}
	})
}

// ReadLineIO is an IO action that reads a line of input from the user
func ReadLineIO(prompt string) IO[string] {
	return New(func() string {
		fmt.Print(prompt)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		return scanner.Text()
	})
}
