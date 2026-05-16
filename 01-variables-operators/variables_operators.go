package main

import (
	"fmt"
)

func main() {
	// --- SECTION 1: Explicit Variable Declaration ---
	fmt.Println("=== 1. EXPLICIT VARIABLE DECLARATION ===")
	var userName string = "Gopher"
	var userAge int = 25
	var userBalance float64 = 1500.75
	var isActive bool = true

	fmt.Printf("Name: %v (%T)\n", userName, userName)
	fmt.Printf("Age: %v (%T)\n", userAge, userAge)
	fmt.Printf("Balance: %v (%T)\n", userBalance, userBalance)
	fmt.Printf("Active: %v (%T)\n\n", isActive, isActive)

	// --- SECTION 2: Short Assignment Operator ---
	fmt.Println("=== 2. SHORT ASSIGNMENT OPERATOR (:=) ===")
	language := "Go"
	version := 1.21
	isFun := true

	fmt.Printf("Language: %v (%T)\n", language, language)
	fmt.Printf("Version: %v (%T)\n", version, version)
	fmt.Printf("Is Fun: %v (%T)\n\n", isFun, isFun)

	// --- SECTION 3: Zero Value Demonstration ---
	fmt.Println("=== 3. ZERO VALUES ===")
	var defaultString string
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool

	fmt.Printf("Default String: '%v'\n", defaultString)
	fmt.Printf("Default Int: %v\n", defaultInt)
	fmt.Printf("Default Float: %v\n", defaultFloat)
	fmt.Printf("Default Bool: %v\n\n", defaultBool)

	// --- SECTION 4: Arithmetic Operations & Type Casting ---
	fmt.Println("=== 4. ARITHMETIC & TYPE CASTING ===")
	a := 10
	b := 3

	sum := a + b
	diff := a - b
	prod := a * b
	rem := a % b

	// Explicit Type Casting for Division
	// Dividing two integers results in an integer. Casting to float64 preserves precision.
	quotient := float64(a) / float64(b)

	fmt.Printf("A: %d, B: %d\n", a, b)
	fmt.Printf("Addition (A + B): %d\n", sum)
	fmt.Printf("Subtraction (A - B): %d\n", diff)
	fmt.Printf("Multiplication (A * B): %d\n", prod)
	fmt.Printf("Modulus (A %% B): %d\n", rem)
	fmt.Printf("Division (float64(A) / float64(B)): %.4f (%T)\n", quotient, quotient, quotient)
}
