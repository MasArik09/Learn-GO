package main

import (
	"fmt"
)

func main() {
	// --- SECTION 1: Shared Underlying Arrays ---
	fmt.Println("=== 1. SHARED UNDERLYING ARRAY (POINTER REFERENCE) ===")
	original := []int{10, 20, 30, 40}
	
	// 'sharedSubSlice' takes elements from index 1 to 3 (20, 30).
	// It shares the SAME underlying array as 'original'.
	sharedSubSlice := original[1:3]

	fmt.Printf("Original:       %v | Address: %p\n", original, original)
	fmt.Printf("Shared SubSlice: %v      | Address: %p\n", sharedSubSlice, sharedSubSlice)
	fmt.Println()

	// --- SECTION 2: Modification Leakage ---
	fmt.Println("=== 2. MODIFICATION LEAKAGE ===")
	fmt.Println("[ACTION] Changing index 0 of sharedSubSlice to 99...")
	sharedSubSlice[0] = 99

	fmt.Printf("Original:       %v (Changed!)\n", original)
	fmt.Printf("Shared SubSlice: %v\n", sharedSubSlice)
	fmt.Printf("[LOGIC] Note how original[1] became 99 because slices are reference types.\n\n")

	// --- SECTION 3: Deep Copying (Isolation) ---
	fmt.Println("=== 3. DEEP COPY (ISOLATED ARCHITECTURE) ===")
	
	// Initialization for isolation using make() and copy()
	isolatedCopy := make([]int, len(original))
	numCopied := copy(isolatedCopy, original)

	fmt.Printf("[INFO] Copied %d elements.\n", numCopied)
	fmt.Printf("Original:       %v | Address: %p\n", original, original)
	fmt.Printf("Isolated Copy:  %v | Address: %p\n", isolatedCopy, isolatedCopy)
	fmt.Println()

	// --- SECTION 4: Proof of Independence ---
	fmt.Println("=== 4. PROOF OF INDEPENDENCE ===")
	fmt.Println("[ACTION] Changing index 0 of isolatedCopy to 500...")
	isolatedCopy[0] = 500

	fmt.Printf("Original:       %v (Untouched)\n", original)
	fmt.Printf("Isolated Copy:  %v (Changed)\n", isolatedCopy)
	fmt.Printf("[LOGIC] Addresses are different (%p != %p), so no leak occurred.\n", original, isolatedCopy)
}
