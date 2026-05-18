package main

import (
	"fmt"
)

// NotificationEngine represents a decoupled orchestration layer.
// It uses Dependency Injection by holding a slice of interface types.
type NotificationEngine struct {
	providers []NotificationProvider
}

// RegisterProvider injects a new concrete implementation into the engine.
func (ne *NotificationEngine) RegisterProvider(p NotificationProvider) {
	ne.providers = append(ne.providers, p)
}

// Broadcast orchestrates the polymorphic execution across all registered providers.
func (ne *NotificationEngine) Broadcast(message string) {
	fmt.Println("[ENGINE] Initiating global broadcast...")

	for _, provider := range ne.providers {
		err := provider.Send(message)
		if err != nil {
			// Architecture: Failures in one provider do not halt the entire system.
			fmt.Printf("[ENGINE ALERT] Failed to send via a provider: %v\n", err)
		}
	}
}

func RunInterfacesEngineDemo() {
	// 5. Instantiate the Orchestrator
	engine := NotificationEngine{}

	// Register concrete implementations (Dependency Injection)
	engine.RegisterProvider(EmailProvider{EmailAddress: "admin@educentral.edu"})
	engine.RegisterProvider(SMSProvider{PhoneNumber: "+1-800-SAFE-EDU"})

	// 6. Execute centralized logic
	engine.Broadcast("System Update: Security patch deployed successfully!")
}
