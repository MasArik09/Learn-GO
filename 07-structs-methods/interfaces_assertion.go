package main

import (
	"fmt"
)

// InspectAndReport demonstrates safe interface unpacking mechanisms.
func InspectAndReport(p NotificationProvider) {
	// 3. Safe Type Assertion using "comma ok" idiom
	// This avoids a runtime panic if the type doesn't match.
	email, ok := p.(EmailProvider)
	if ok {
		fmt.Printf("[INSPECTION] [SUCCESS] Detected EmailProvider. Target Address: %s\n", email.EmailAddress)
	}

	// 4. Flexible Type Switch
	// This is the idiomatic way to handle multiple potential concrete types elegantly.
	switch v := p.(type) {
	case EmailProvider:
		fmt.Println("[SWITCH] Acknowledged: This is an email-based notification system.")
	case SMSProvider:
		fmt.Printf("[SWITCH] Detected SMSProvider. Target Phone: %s\n", v.PhoneNumber)
	default:
		fmt.Println("[INSPECTION] Unknown provider type.")
	}
}

func RunInterfacesAssertionDemo() {
	// 5. Instantiate and wrap in interface
	var provider NotificationProvider = EmailProvider{EmailAddress: "arthur@example.edu"}

	fmt.Println("--- Starting Interface Inspection ---")
	InspectAndReport(provider)

	// Test with SMS
	fmt.Println("\n--- Testing with SMS Provider ---")
	sms := SMSProvider{PhoneNumber: "+1-555-0199"}
	InspectAndReport(sms)

	/*
	   6. ARCHITECTURAL DANGER ZONE:

	   The following line is commented out because it represents a critical stability risk:

	   badAssertion := provider.(SMSProvider)

	   EXPLANATION: A direct type assertion without the "comma ok" idiom or a type switch
	   is non-defensive. If the interface does NOT hold the asserted type (e.g., asserting
	   an EmailProvider as an SMSProvider), the Go runtime will trigger an immediate
	   RUNTIME PANIC, crashing the entire application. Always use the 'ok' boolean
	   or a 'switch' to ensure fault-tolerant type resolution.
	*/
}
