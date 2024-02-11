/*The code establishes a connection to RabbitMQ, creates a channel,
declares a queue, registers a consumer,and processes incoming messages,
logging the received patient data.
*/

package main

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"secureCommPlatformHealthcare/data"
)

func main() {
	// Establish a connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	data.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	data.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a RabbitMQ queue named "patient_lab_test_queue"
	q, err := ch.QueueDeclare(
		"patient_lab_test_queue", // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	data.FailOnError(err, "Failed to declare a queue")

	// Register a consumer to receive messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	data.FailOnError(err, "Failed to register a consumer")

	// Loop to process incoming messages
	for msg := range msgs {

		// Deserialize the received JSON message into the PatientLabData struct
		var patientData data.PatientLabData
		err := json.Unmarshal(msg.Body, &patientData)
		data.FailOnError(err, "Failed to decode JSON")
		log.Printf("Received patient data: %+v", patientData)
	}

	// Log a message indicating that all patient data has been received and processed
	log.Println("All patient data received and processed")
}
