package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/adityjoshi/Swaasthya/Backend/database"
)

// SubscribeToPaymentUpdates listens for payment updates from Redis.
func SubscribeToPaymentUpdates() {
	redisClient := database.GetRedisClient()
	pubsub := redisClient.Subscribe(context.Background(), "patient_payment_updates")
	defer pubsub.Close()

	fmt.Println("Compounder subscribed to patient payment updates...")

	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Println("Error receiving message:", err)
			continue
		}

		// Process the received message
		fmt.Printf("Received payment update: %s\n", msg.Payload)
		// Example: Notify the compounder, update the UI, etc.
	}
}

// SubscribeToHospitalizationUpdates listens for Redis messages about patient hospitalization
func SubscribeToHospitalizationUpdates() {
	pubsub := database.GetRedisClient().Subscribe(context.Background(), "hospitalized-patients")

	// Infinite loop to listen for messages
	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Printf("Error receiving Redis message: %v", err)
			continue
		}

		// Log or process the received message
		fmt.Printf("Hospitalization Update: %s\n", msg.Payload)

		// You can trigger any action here, such as updating the frontend
	}
}
