/*
Patient data retrieval client side code.
In order to retrieve patient data,
it sends a GET request to a server endpoint after encrypting a patient ID using a encryption module.
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"secureCommPlatformHealthcare/networking_client_server/cipher"
	"strings"
)

// Function to encrypt patient ID before sending to the server
func encryptPatientID(plainText string) string {
	return cipher.Rot13Encrypt(plainText)
}

func main() {
	// Ask user for patient ID and encrypt it
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter patient ID: ")
	patientID, _ := reader.ReadString('\n')
	patientID = strings.TrimSpace(patientID)

	encryptedID := encryptPatientID(patientID)
	fmt.Print("Encrypted patient ID: ", encryptedID)
	// Make a GET request with the encrypted ID
	resp, err := http.Get("http://localhost:8080/patient/" + encryptedID)

	if err != nil {
		fmt.Println("Client error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Decode and print patient data received from the server
		var patientData map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&patientData)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}
		fmt.Println("\nPatient Data:")
		fmt.Println("ID:", patientData["id"])
		fmt.Println("Name:", patientData["name"])
		fmt.Println("Age:", patientData["age"])
		fmt.Println("Diagnosis:", patientData["diagnosis"])
	} else {
		fmt.Println("Patient not found or server error:", resp.Status)
	}
}
