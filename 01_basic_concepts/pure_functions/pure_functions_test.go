package purefunctions

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{4, 24},
		{5, 120},
	}

	for _, test := range tests {
		result := Factorial(test.input)
		if result != test.expected {
			t.Errorf("Factorial(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestImpureRandomNumber(t *testing.T) {
	// Call the function twice and check if results differ, indicating impurity
	result1 := ImpureRandomNumber(100)
	result2 := ImpureRandomNumber(100)
	if result1 == result2 {
		t.Errorf("ImpureRandomNumber(100) = %d; expected different result on subsequent calls", result1)
	}
}

func TestImpureTimestamp(t *testing.T) {
	// Call the function twice and check if results differ, indicating impurity
	result1 := ImpureTimestamp()
	time.Sleep(1 * time.Second) // Delay to ensure different timestamps
	result2 := ImpureTimestamp()
	if result1 == result2 {
		t.Errorf("ImpureTimestamp() = %d; expected different result on subsequent calls", result1)
	}
}

func TestPureRandomNumber(t *testing.T) {
	seed := int64(42)
	limit := 100
	result1 := PureRandomNumber(seed, limit)
	result2 := PureRandomNumber(seed, limit)
	if result1 != result2 {
		t.Errorf("PureRandomNumber(%d, %d) = %d; expected same result for identical seeds", seed, limit, result1)
	}

	// Test that a different seed gives a different result
	anotherSeed := int64(43)
	result3 := PureRandomNumber(anotherSeed, limit)
	if result1 == result3 {
		t.Errorf("PureRandomNumber(%d, %d) = %d; expected different result for different seeds", anotherSeed, limit, result3)
	}
}
