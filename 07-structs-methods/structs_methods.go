package main

import (
	"fmt"
)

// Introduce uses a Value Receiver.
// It receives a COPY of the struct, so changes made here don't affect the original.
func (s Student) Introduce() {
	// Attempting to modify the copy
	s.Name = "Anonymous"
	fmt.Printf("[VALUE RECEIVER] Hello, my name is %s and my GPA is %.2f\n", s.Name, s.gpa)
}

// UpdateGPAPointer uses a Pointer Receiver.
// It receives a reference to the actual memory address, allowing for mutation.
func (s *Student) UpdateGPAPointer(newGPA float64) {
	fmt.Printf("[MUTATION] Updating GPA from %.2f to %.2f\n", s.gpa, newGPA)
	s.gpa = newGPA
}

func RunMethodsDemo() {
	// 1. Initial instantiation
	student := Student{Name: "Arthur", gpa: 3.50}
	fmt.Printf("--- Initial State in main: Name=%s, GPA=%.2f ---\n\n", student.Name, student.gpa)

	// 2. Call Introduce (Value Receiver)
	// Even though it sets Name to "Anonymous" internally, the original remains Arthur.
	student.Introduce()
	fmt.Printf("--- State in main after Introduce(): Name=%s (Memory Isolated) ---\n\n", student.Name)

	// 3. Call UpdateGPAPointer (Pointer Receiver)
	// This will permanently change the GPA at the memory address.
	student.UpdateGPAPointer(3.95)

	// 4. Final verification
	// We'll see the updated GPA but the original Name.
	student.Introduce()
	fmt.Printf("--- Final State in main: Name=%s, GPA=%.2f (Permanently Mutated) ---\n", student.Name, student.gpa)
}
