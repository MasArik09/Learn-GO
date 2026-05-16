package main

import (
	"fmt"
)

func main() {
	// --- SECTION 1: Standard If-Else ---
	fmt.Println("=== 1. STANDARD IF-ELSE ===")
	score := 75
	passingScore := 60

	if score >= passingScore {
		fmt.Printf("[PASSED] Score: %d (Minimum required: %d)\n", score, passingScore)
	} else {
		fmt.Printf("[FAILED] Score: %d (Minimum required: %d)\n", score, passingScore)
	}
	fmt.Println()

	// --- SECTION 2: If - Else If - Else with Logical Operators ---
	fmt.Println("=== 2. IF-ELSE IF-ELSE CHAIN ===")
	examGrade := 88
	attendancePercent := 95

	// Using logical AND (&&) to check two conditions
	if examGrade >= 90 && attendancePercent >= 90 {
		fmt.Println("[GRADE] A - Exceptional performance!")
	} else if examGrade >= 80 || attendancePercent >= 95 {
		// Using logical OR (||)
		fmt.Println("[GRADE] B - Solid effort!")
	} else if examGrade >= 70 {
		fmt.Println("[GRADE] C - Satisfactory results.")
	} else {
		fmt.Println("[GRADE] Fail - Improvement needed.")
	}
	fmt.Println()

	// --- SECTION 3: If with a Short Statement ---
	fmt.Println("=== 3. IF WITH SHORT STATEMENT ===")
	
	/*
	   Syntax: if statement; condition { }
	   The variable 'errCode' is only accessible within the scope of the 
	   if, else if, and else blocks. It is not available outside this structure.
	*/
	if errCode := 404; errCode == 200 {
		fmt.Printf("[INFO] Status: %d (Success)\n", errCode)
	} else if errCode == 404 {
		fmt.Printf("[INFO] Status: %d (Not Found)\n", errCode)
	} else {
		fmt.Printf("[INFO] Status: %d (Unknown Error)\n", errCode)
	}

	// fmt.Println(errCode) // Uncommenting this would cause a compilation error: "undefined: errCode"
}
