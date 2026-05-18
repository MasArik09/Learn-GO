package main

import (
	"fmt"
)

func main() {
	// --- SECTION 1: Fixed-Size Arrays ---
	fmt.Println("=== 1. FIXED-SIZE ARRAYS (RIGID ARCHITECTURE) ===")
	// Arrays have a fixed length defined at compile time.
	var languages [3]string = [3]string{"Go", "Python", "Java"}

	/*
	   COMPILATION ERROR DEMONSTRATION:
	   Uncommenting the line below would cause: "invalid argument: index 3 out of bounds [0:3]"
	   languages[3] = "C++"
	*/

	fmt.Printf("Fixed Array: %v\n", languages)
	fmt.Printf("Array Length: %d\n\n", len(languages))

	// --- SECTION 2: Dynamic Slices & Memory Allocation ---
	fmt.Println("=== 2. DYNAMIC SLICES (AUTO-GROWING ARCHITECTURE) ===")

	// 'make' initializes a slice with a length of 0 and a capacity of 2.
	// Capacity is the number of elements the underlying array can hold before re-allocation.
	frameworks := make([]string, 0, 2)
	fmt.Printf("Initial: len=%d cap=%d %v\n", len(frameworks), cap(frameworks), frameworks)

	itemsToAdd := []string{"Gin", "Echo", "Fiber", "Revel", "Buffalo"}

	for i, item := range itemsToAdd {
		frameworks = append(frameworks, item)
		// Logic & Architecture: Observe how 'cap' doubles or grows when 'len' exceeds it.
		fmt.Printf("Append %d [%s]: len=%d cap=%d %v\n", i+1, item, len(frameworks), cap(frameworks), frameworks)
	}

	fmt.Println()

	// --- SECTION 3: Defensive Boundary Validation ---
	fmt.Println("=== 3. DEFENSIVE VALIDATION GATE ===")
	targetIndex := 4

	// Validating length before access to prevent runtime index-out-of-bounds panic
	if len(frameworks) > targetIndex {
		fmt.Printf("[SAFE ACCESS] Element at index %d: %s\n", targetIndex, frameworks[targetIndex])
	} else {
		fmt.Printf("[CRITICAL] Index %d is out of bounds for slice length %d\n", targetIndex, len(frameworks))
	}

	// Centralized Execution for 05-arrays-slices
	fmt.Println("\n=== Executing Advanced Slices Demo ===")
	RunAdvancedSlicesDemo()
}
