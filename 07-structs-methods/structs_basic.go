package main

import (
	"fmt"
)

// Student represents an academic entity
type Student struct {
	ID    int     // Exported field
	Name  string  // Exported field
	Major string  // Exported field
	gpa   float64 // Unexported field: starts with lowercase, encapsulated within package main
}

func main() {
	// 1. Initialization using Named Struct Literal
	student1 := Student{
		ID:    1,
		Name:  "Arthur",
		Major: "Information Systems",
		gpa:   3.95,
	}

	// 2. Initialization using Positional Struct Literal
	// Note: Fields must be in the exact order defined in the struct.
	student2 := Student{2, "Morgan", "Computer Science", 3.82}

	// 3. Visualization
	fmt.Println("--- Student Registry ---")
	fmt.Printf("Instance 1 (Named): %+v\n", student1)
	fmt.Printf("Instance 2 (Positional): %+v\n\n", student2)

	// Centralized Execution Flow (Package-Level Cooperation)
	fmt.Println("=== Executing Methods Demo ===")
	RunMethodsDemo()
	fmt.Println("\n=== Executing Constructor Demo ===")
	RunConstructorDemo()
	fmt.Println("\n=== Executing Interfaces Basic Demo ===")
	RunInterfacesBasicDemo()
	fmt.Println("\n=== Executing Interfaces Engine Demo ===")
	RunInterfacesEngineDemo()
	fmt.Println("\n=== Executing Interfaces Assertion Demo ===")
	RunInterfacesAssertionDemo()
}
