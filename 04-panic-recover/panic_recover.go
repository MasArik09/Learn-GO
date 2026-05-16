package main

import (
	"fmt"
)

// accessSecureAsset simulates a scenario that triggers a fatal runtime error.
func accessSecureAsset() {
	fmt.Println("[INFO] Attempting to access secure asset...")
	
	// Intentionally triggering a manual panic to simulate a critical failure.
	panic("Unauthorized system breach simulated!")
}

func main() {
	// 1. Deferred Recovery Wrapper
	// This function is scheduled to run when main() returns, even if a panic occurs.
	defer func() {
		if r := recover(); r != nil {
			// 2. Recovery Logic
			// If recover() returns a value, it means a panic was caught.
			fmt.Printf("[SYSTEM RECOVERY] Caught a critical panic: %v\n", r)
			fmt.Println("[SYSTEM RECOVERY] Process state remains stable. Exiting gracefully...")
		}
	}()

	fmt.Println("[START] Application initiated.")

	// 3. Triggering the failure
	accessSecureAsset()

	// 4. Observation Point
	// This line will NEVER be reached because the panic stops the normal 
	// execution flow of the current goroutine.
	fmt.Println("[END] Application finished successfully.")
}
