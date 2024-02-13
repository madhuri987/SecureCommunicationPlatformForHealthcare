/*
The code demonstrates the process of serializing patient data, publishing it to a RabbitMQ queue,
and logging the status of the operation.
*/

package main

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go" //import library first
	"log"
	"secureCommPlatformHealthcare/data"
	"time"
)

func main() {
	//Connect to RabitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	data.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//Create a channel
	ch, err := conn.Channel()
	data.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	//Declare a queue to send the message
	q, err := ch.QueueDeclare(
		"patient_lab_test_queue", // name
		false,                    // durable
		false,                    // delete when unused
		false,                    // exclusive
		false,                    // no-wait
		nil,                      // arguments
	)
	data.FailOnError(err, "Failed to declare a queue")

	// patient data for 5 patients
	patients := []data.PatientLabData{
		{PatientID: "P001", TestType: "Blood Test", FirstName: "Madhuri", LastName: "Pawar", Age: 30, Gender: "Female", Timestamp: time.Now()},
		{PatientID: "P002", TestType: "Urine Test", FirstName: "John", LastName: "Doe", Age: 28, Gender: "Male", Timestamp: time.Now()},
		{PatientID: "P003", TestType: "MRI Scan", FirstName: "Akshay", LastName: "Parab", Age: 18, Gender: "Male", Timestamp: time.Now()},
		{PatientID: "P004", TestType: "X-Ray", FirstName: "Seema", LastName: "Rao", Age: 36, Gender: "Female", Timestamp: time.Now()},
		{PatientID: "P005", TestType: "CT Scan", FirstName: "Nikita", LastName: "Pagar", Age: 22, Gender: "Female", Timestamp: time.Now()},
	}

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure cancel is called to release resources when the function exits

	// Iterate over each patient
	for _, patient := range patients {

		// Serialize the patient object to JSON
		body, err := json.Marshal(patient)
		data.FailOnError(err, "Failed to serialize JSON")

		// Publish the JSON data to the RabbitMQ queue
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		data.FailOnError(err, "Failed to publish a message")

		// Log a message indicating that patient data has been sent
		log.Printf("Patient data sent: %+v", patient)
	}

	// Log a message indicating that patient data for lab tests has been sent successfully
	log.Println("Patient data for lab tests sent successfully")
}
