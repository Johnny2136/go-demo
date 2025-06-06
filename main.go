package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// API interface that gets the value that will be passed to the Func fibonacci
func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the 'n' query parameter
	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		http.Error(w, "Invalid value for 'n'", http.StatusBadRequest)
		return
	}

	// Calculate the Fibonacci sequence up to the nth number
	result := fibonacci(n)

	// Set Content-Type without charset=utf-8
	w.Header().Set("Content-Type", "text/plain")

	// Write the Fibonacci sequence as a response
	fmt.Fprintf(w, "Fibonacci sequence up to %d: %v", n, result)
}

// Func fibonacci generates a fibonacci sequence based on "fibonacci(n)" passed from the handler
func fibonacci(n int) []int {
	sequence := make([]int, n)
	if n > 0 {
		sequence[0] = 0
	}
	if n > 1 {
		sequence[1] = 1
	}
	for i := 2; i < n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}
	return sequence
}

func main() {
	http.HandleFunc("/fibonacci", fibonacciHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
