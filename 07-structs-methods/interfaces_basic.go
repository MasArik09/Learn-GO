package main

import (
	"fmt"
)

// NotificationProvider defines the architectural contract for sending alerts.
// Any struct that implements the Send(string) error method satisfies this interface implicitly.
type NotificationProvider interface {
	Send(message string) error
}

// EmailProvider is a concrete implementation of the NotificationProvider interface.
type EmailProvider struct {
	EmailAddress string
}

// Send implements the NotificationProvider interface for EmailProvider.
func (e EmailProvider) Send(message string) error {
	fmt.Printf("[EMAIL] Sending email to %s: %s\n", e.EmailAddress, message)
	return nil
}

// SMSProvider is a second concrete implementation of the NotificationProvider interface.
type SMSProvider struct {
	PhoneNumber string
}

// Send implements the NotificationProvider interface for SMSProvider.
func (s SMSProvider) Send(message string) error {
	fmt.Printf("[SMS] Sending SMS to %s: %s\n", s.PhoneNumber, message)
	return nil
}

func RunInterfacesBasicDemo() {
	// 4. Interface Architecture: Polymorphic Slice
	// We create a slice of the interface type, not concrete types.
	providers := []NotificationProvider{
		EmailProvider{EmailAddress: "arthur@example.edu"},
		SMSProvider{PhoneNumber: "+1-555-0199"},
	}

	fmt.Println("--- System Dispatching Notifications ---")

	// 5. Dynamic Dispatch
	// The Go runtime determines which Send method to call based on the concrete type stored.
	message := "Welcome to EduCentral-Hub!"
	for _, provider := range providers {
		err := provider.Send(message)
		if err != nil {
			fmt.Printf("Error sending notification: %v\n", err)
		}
	}
}
