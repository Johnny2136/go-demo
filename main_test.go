package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFibonacciHandlerInvalidInput(t *testing.T) {
	req, err := http.NewRequest("GET", "/fibonacci?n=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fibonacciHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestFibonacciHandlerValidInput(t *testing.T) {
	req, err := http.NewRequest("GET", "/fibonacci?n=5", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fibonacciHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "Fibonacci sequence up to 5: [0 1 1 2 3]"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestFibonacciFunction(t *testing.T) {
	// Test cases for fibonacci function
	tests := []struct {
		n        int
		expected []int
	}{
		{0, []int{}},
		{1, []int{0}},
		{2, []int{0, 1}},
		{5, []int{0, 1, 1, 2, 3}},
		{10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
	}

	for _, test := range tests {
		result := fibonacci(test.n)
		if len(result) != len(test.expected) {
			t.Errorf("fibonacci(%d) returned unexpected length: got %v want %v",
				test.n, len(result), len(test.expected))
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("fibonacci(%d) returned unexpected result at index %d: got %v want %v",
					test.n, i, result[i], test.expected[i])
			}
		}
	}
}

func TestServer(t *testing.T) {
	// Create a new HTTP request to test the server
	req, err := http.NewRequest("GET", "/fibonacci?n=5", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a new HTTP handler using the fibonacciHandler function
	handler := http.HandlerFunc(fibonacciHandler)

	// Serve the HTTP request with the handler and record the response
	handler.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the body of the response
	expected := "Fibonacci sequence up to 5: [0 1 1 2 3]"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
