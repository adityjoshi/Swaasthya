package database

import (
	"encoding/json"
	"fmt"
)

// ListenForPatientUpdates - Compounder subscribes to patient updates
func ListenForPatientUpdates() {
	pubsub := GetRedisClient().Subscribe(Ctx, "patient_updates")
	defer pubsub.Close()

	for {
		// Receive messages from the "patient_updates" channel
		msg, err := pubsub.ReceiveMessage(Ctx)
		if err != nil {
			fmt.Println("Error receiving message:", err)
			continue
		}

		// Process the message (JSON payload)
		var patientUpdate map[string]interface{}
		if err := json.Unmarshal([]byte(msg.Payload), &patientUpdate); err != nil {
			fmt.Println("Error decoding message payload:", err)
			continue
		}

		fmt.Printf("Received Patient Update: %v\n", patientUpdate)

		// You can trigger any action here, such as updating the frontend or sending an email
	}
}
