package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Engagespot/engagespot-go/engagespot"
)

func main() {
	// Get API key and secret from environment variables
	apiKey := os.Getenv("ENGAGESPOT_API_KEY")
	apiSecret := os.Getenv("ENGAGESPOT_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		log.Fatal("ENGAGESPOT_API_KEY and ENGAGESPOT_API_SECRET must be set")
	}

	// Create a new Engagespot client
	client := engagespot.NewClient(apiKey, apiSecret)

	// Create a notification
	notification := &engagespot.Notification{
		Workflow: engagespot.Workflow{
			Identifier: "plan-expiry-notification",
		},
		Data: map[string]interface{}{
			"message": "This is a test notification",
			"url":     "https://example.com",
		},
		SendTo: engagespot.SendTo{
			Recipients: []string{"user-001"},
		},
	}

	// Send the notification
	response, err := client.Notifications.Send(notification)
	if err != nil {
		log.Fatalf("Error sending notification: %v", err)
	}

	fmt.Printf("Notification sent successfully. Request ID: %s\n", response.RequestID)
}