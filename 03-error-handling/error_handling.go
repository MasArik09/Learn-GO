package main

import (
	"errors"
	"fmt"
)

// divideNumbers performs division and returns the result and an error if applicable.
// In Go, it's idiomatic to return the result first and the error second.
func divideNumbers(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		// Returning a sanitized error message using errors.New
		return 0, errors.New("division by zero is not allowed")
	}

	// Operation success: return the result and nil for the error
	return numerator / denominator, nil
}

func main() {
	// --- Path A: Success Scenario ---
	fmt.Println("=== SCENARIO A: VALID INPUT ===")
	num1, den1 := 10.0, 2.0
	result1, err1 := divideNumbers(num1, den1)

	// Idiomatic 'if err != nil' pattern
	if err1 != nil {
		fmt.Printf("[ERROR] Failed to divide %.2f by %.2f: %v\n", num1, den1, err1)
	} else {
		fmt.Printf("[SUCCESS] %.2f / %.2f = %.2f\n", num1, den1, result1)
	}
	fmt.Println()

	// --- Path B: Failure Scenario ---
	fmt.Println("=== SCENARIO B: DIVISION BY ZERO ===")
	num2, den2 := 10.0, 0.0
	result2, err2 := divideNumbers(num2, den2)

	// Graceful error handling
	if err2 != nil {
		fmt.Printf("[ERROR] %v\n", err2)
	} else {
		fmt.Printf("[SUCCESS] %.2f / %.2f = %.2f\n", num2, den2, result2)
	}

	fmt.Println("\n[INFO] Program execution completed gracefully despite the error.")

	// Centralized Execution for 03-error-handling
	fmt.Println("\n=== Executing Custom Errors Demo ===")
	RunCustomErrorsDemo()
}
