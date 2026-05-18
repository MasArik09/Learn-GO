package main

import "fmt"

func main() {
	userRoles := make(map[string]string)

	userRoles["alice"] = "admin"
	userRoles["bob"] = "developer"
	userRoles["carol"] = "analyst"

	if role, ok := userRoles["alice"]; ok {
		fmt.Printf("[SUCCESS] user=%q role=%q\n", "alice", role)
	} else {
		fmt.Printf("[NOT FOUND] user=%q role=%q\n", "alice", role)
	}

	if role, ok := userRoles["dave"]; ok {
		fmt.Printf("[SUCCESS] user=%q role=%q\n", "dave", role)
	} else {
		fmt.Printf("[NOT FOUND] user=%q role=%q (safe default)\n", "dave", "")
	}

	// Centralized Execution for 06-maps
	fmt.Println("\n=== Executing Maps Advanced Demo ===")
	RunMapsAdvancedDemo()
}
