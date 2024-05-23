# Immutability in Go

Immutability is a fundamental concept in functional programming (FP). In an immutable programming model, once a variable or data structure is created, it cannot be modified. Instead of altering data, operations create new versions of the data with the desired changes. Immutability promotes safe concurrency, predictable behavior, and easier reasoning about state.

This section explores immutability in Go, showing how to implement and leverage immutable patterns in your programs.

## 1. What is Immutability?

In Go, immutability means that once data is created, it cannot be modified. Instead of changing data in place, we create new data structures with updated values.

### Benefits of Immutability

- **Predictability**: Immutable data guarantees no unexpected changes, making it easier to reason about the state of a program.
- **Concurrency**: In concurrent programming, immutability reduces the risk of race conditions, as immutable objects cannot be changed by multiple goroutines.
- **Safety**: It helps avoid bugs caused by unintended side effects or shared state.

### Common Go Data Structures and Immutability

- **Slices**: Go slices are often passed by reference, so altering a slice affects the original. To ensure immutability, you must explicitly create a new slice.
- **Maps**: Like slices, Go maps are mutable by default. A copy of a map can be made to preserve the original.
- **Structs**: Go structs are typically passed by value. When working with structs, it's important to return modified copies rather than modifying the original instances.

---

## 2. Key Immutability Concepts and Examples

### Copying vs Modifying Data

A common pattern in Go for achieving immutability is **copying** data rather than **modifying** it. This can be done with slices, maps, and structs by creating new instances with updated values.

#### Example: Appending to a Slice Immutably

In Go, slices are reference types, meaning that appending to a slice changes its contents in place. To implement immutability, we create a new slice with the desired change:

```go
// Appending to a slice without modifying the original
func AppendImmutable(slice []int, element int) []int {
	newSlice := make([]int, len(slice)+1)
	copy(newSlice, slice)
	newSlice[len(slice)] = element
	return newSlice
}

```

This function creates a new slice and adds the new element, leaving the original slice unchanged.

#### Example: Updating a Struct Immutably
When working with structs, it's important to create and return a modified copy rather than directly changing the original.

```go
type Person struct {
	Name string
	Age  int
}

// Updating a person's age immutably
func UpdateAgeImmutable(p Person, newAge int) Person {
	return Person{Name: p.Name, Age: newAge}
}
```
## 3. Additional Immutability Patterns

### Avoiding Global Mutable State

Global variables can easily introduce mutable state in a program. To keep things immutable, we avoid exposing global variables directly. Instead, we return copies or provide controlled access.

#### Example: Immutable API Endpoints

```go
var apiEndpoints = map[string]string{
	"auth": "/api/auth",
	"user": "/api/user",
}

// Retrieve a copy of the API endpoints
func GetAPIEndpoints() map[string]string {
	newMap := make(map[string]string, len(apiEndpoints))
	for k, v := range apiEndpoints {
		newMap[k] = v
	}
	return newMap
}
```

In this example, the function `GetAPIEndpoints` returns a copy of the `apiEndpoints` map to avoid modifying the original global map.

## 4. Recursion and Immutability
In functional programming, recursion is a key technique for processing data without using mutable state. Go supports recursion, and it can be used to process data immutably.
#### Example: Recursive Factorial Function

```go
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
```

The `Factorial` function is a simple recursive function that avoids mutable state by using the recursive call stack to compute the result.

## 5. Conclusion
Immutability in Go allows for safer, more predictable code, especially when working with concurrent or complex systems. By avoiding direct modification of data, we ensure that data integrity is maintained, and the program's state is easier to reason about.

#### Key takeaways:

- Instead of modifying existing data structures, create new versions of them with the desired changes.
- Use immutability patterns to prevent side effects in your code.
- Be mindful of Go’s mutable types (slices, maps, structs) and manage them immutably by returning modified copies.

Explore the `immutability.go` file for practical implementations and experiment with these immutability patterns in your own code!
