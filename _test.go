package main

import (
	"reflect"
	"testing"
)

// TestFibonacci tests the fibonacci function.
func TestFibonacci(t *testing.T) {
	// Test cases
	tests := []struct {
		n        int   // Input value
		expected []int // Expected result
	}{
		{0, []int{}},          // n = 0
		{1, []int{0}},         // n = 1
		{2, []int{0, 1}},      // n = 2
		{5, []int{0, 1, 1, 2, 3}}, // n = 5
	}

	// Iterate over test cases
	for _, test := range tests {
		// Call the fibonacci function with the test input
		result := fibonacci(test.n)

		// Compare the result with the expected value
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("fibonacci(%d) returned %v, expected %v", test.n, result, test.expected)
		}
	}
}