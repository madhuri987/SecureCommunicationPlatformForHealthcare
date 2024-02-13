/*
Patient data retrieval server side code.
Custom encryption module from the "cipher" package is used to obtain patient data based on decrypted patient ID.
For the purpose of securely retrieving patient data, the server exposes an endpoint ("/patient/{id}") that
demonstrates HTTP handling, JSON parsing, and encryption/decryption.
*/
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"secureCommPlatformHealthcare/networking_client_server/cipher" // Import the cipher package
)

// Patient struct represents patient information
type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Diagnosis string `json:"diagnosis"`
}

var patients []Patient

func main() {
	// Read patient data from the JSON file
	file, err := ioutil.ReadFile("networking_client_server/patients.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON data into the patients slice
	err = json.Unmarshal(file, &patients)
	if err != nil {
		panic(err)
	}

	// Define a handler for patient data retrieval ("/patient/{id}")
	http.HandleFunc("/patient/", getPatientByID)

	// Start the server.
	port := ":8080"
	println("Server is listening on port", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

// Add a function to handle decryption of patient IDs
func decryptPatientID(cipherText string) string {
	return cipher.Rot13Decrypt(cipherText)
}

// getPatientByID retrieves patient data by ID (decrypted)
func getPatientByID(w http.ResponseWriter, r *http.Request) {
	// Extract and decrypt patient ID from URL
	id := r.URL.Path[len("/patient/"):]
	decryptedID := decryptPatientID(id)

	// Search for the patient with the decrypted ID
	for _, patient := range patients {
		if patient.ID == decryptedID {
			// Respond with JSON-encoded patient data
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(patient)
			return
		}
	}
	// Patient not found
	http.Error(w, "Patient not found", http.StatusNotFound)
}
