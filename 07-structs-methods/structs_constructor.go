package main

import (
	"fmt"
)

// NewStudent acts as a Factory Constructor with architectural guard clauses.
// It returns a pointer to a Student and an error, enforcing safe object initialization.
func NewStudent(id int, name string, major string, initialGPA float64) (*Student, error) {
	// 3. Defensive validation gate (Guard Clause)
	if initialGPA < 0.0 || initialGPA > 4.0 {
		return nil, fmt.Errorf("validation error: GPA %.2f is out of bounds (0.0 - 4.0)", initialGPA)
	}

	// 4. Successful allocation and pointer return
	return &Student{
		ID:    id,
		Name:  name,
		Major: major,
		gpa:   initialGPA,
	}, nil
}

func RunConstructorDemo() {
	// Path A: Invalid Input Simulation
	fmt.Println("--- Path A: Simulating Invalid Input ---")
	s1, err := NewStudent(101, "Invalid User", "Undeclared", 4.5)
	if err != nil {
		fmt.Printf("[VALIDATION FAILED] %v\n", err)
	} else {
		fmt.Printf("[OBJECT CREATED] %+v\n", s1)
	}

	fmt.Println("\n--- Path B: Simulating Valid Input ---")
	// Path B: Valid Input Simulation
	s2, err := NewStudent(102, "Arthur", "Information Systems", 3.85)
	if err != nil {
		fmt.Printf("[VALIDATION FAILED] %v\n", err)
	} else {
		// %+v prints field names + values
		fmt.Printf("[OBJECT CREATED] %+v\n", s2)
	}
}
