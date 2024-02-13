/*
Patient lab data and error handling
*/
package data

import (
	"log"
	"time"
)

// PatientLabData represents the patient's lab test information
type PatientLabData struct {
	PatientID string    `json:"patient_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	TestType  string    `json:"test_type"`
	Age       int       `json:"age"`
	Gender    string    `json:"gender"`
	Timestamp time.Time `json:"timestamp"`
}

// FailOnError is an exported function to handle errors
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
