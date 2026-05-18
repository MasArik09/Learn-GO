package main

import (
	"fmt"
)

func RunMapsAdvancedDemo() {
	// 1. Initialize activeSessions map with map literal
	activeSessions := map[string]string{
		"token_abc123": "alice",
		"token_def456": "bob",
		"token_ghi789": "charlie",
		"token_jkl012": "dave",
	}

	// 2. Iterate and print every active session
	fmt.Println("--- Current Active Sessions ---")
	printSessions(activeSessions)

	// 3. Safe deletion of an existing key
	logoutToken := "token_def456"
	fmt.Printf("\nLogging out user with token: %s\n", logoutToken)
	delete(activeSessions, logoutToken)

	// 4. Delete a non-existent key to demonstrate idempotency
	nonExistentToken := "token_xyz999"
	fmt.Printf("Attempting to delete non-existent token: %s\n", nonExistentToken)
	delete(activeSessions, nonExistentToken)

	// 5. Iterate one final time to verify deletion
	fmt.Println("\n--- Post-Logout Sessions ---")
	printSessions(activeSessions)
}

// Helper function to iterate through the map and print aligned rows
func printSessions(sessions map[string]string) {
	if len(sessions) == 0 {
		fmt.Println("No active sessions.")
		return
	}
	for token, user := range sessions {
		fmt.Printf("Token: %-15s | User: %s\n", token, user)
	}
}
