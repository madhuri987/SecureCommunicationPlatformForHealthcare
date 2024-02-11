/*
Sensitive patient information like lab test report is encoded for secure transmission or storage
and then decoded for interpretation or analysis
*/
package main

import (
	"fmt"
	"strings"
)

// Define your morse code
var customAlphabetMap = map[rune]string{
	'A': "@@", 'B': "$!", 'C': "$$", 'D': "()", 'E': ".",
	'F': "..-.", 'G': "--.", 'H': "$%", 'I': "!!", 'J': ".---",
	'K': "-.-", 'L': ".-..", 'M': "$$", 'N': "-.", 'O': "---",
	'P': ".--.", 'Q': "--.-", 'R': "--", 'S': "...", 'T': "-",
	'U': "..", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
	'Z': "--..",
	'1': ".----", '2': "..---", '3': "...--", '4': "....-",
	'5': ".....", '6': "-....", '7': "--...", '8': "---..",
	'9': "----.", '0': "-----",
}

// Custom integrity check function
func checkIntegrity(originalText, decodedText string) bool {

	// Convert both original and decoded texts to uppercase for case-insensitive comparison
	originalText = strings.ToUpper(strings.TrimSpace(originalText))
	decodedText = strings.ToUpper(strings.TrimSpace(decodedText))

	// Perform a simple integrity check by comparing the modified original text with the decoded text
	return originalText == decodedText
}

// textToMorse converts a given text to Morse code using the provided alphabet map.

func textToMorse(text string, alphabetMap map[rune]string) string {
	// Convert the text to uppercase to ensure case-insensitive mapping
	text = strings.ToUpper(text)

	// Initialize a string builder to store the Morse code
	var morseCode strings.Builder

	// Iterate over each character in the text
	for _, char := range text {
		// Check if the character is a space
		if char == ' ' {
			// Add spaces around the "/" delimiter
			morseCode.WriteString(" / ")
		} else if code, found := alphabetMap[char]; found {
			// Append the Morse code for the character, followed by a space
			morseCode.WriteString(code + " ")
		}
	}

	// Return the final Morse code as a string
	return morseCode.String()
}

// morseToText converts a given Morse code to text using the provided alphabet map.
func morseToText(morse string, alphabetMap map[rune]string) string {
	// Split the Morse code into words using "/" as the delimiter
	words := strings.Split(morse, "/")

	// Initialize a string builder to store the decoded text
	var decodedText strings.Builder

	// Iterate over each word in the Morse code
	for _, word := range words {
		// Split each word into Morse code letters
		letters := strings.Fields(strings.TrimSpace(word))

		// Iterate over each Morse code letter in the word
		for _, letter := range letters {
			// Iterate over the alphabet map to find the corresponding character
			for key, value := range alphabetMap {
				// Check if the Morse code matches the value in the alphabet map
				if value == letter {
					// Append the decoded character to the result
					decodedText.WriteRune(key)
					break
				}
			}
		}

		// Add a space between words in the decoded text
		decodedText.WriteRune(' ')
	}

	// Return the final decoded text as a string
	return decodedText.String()
}

func main() {
	// Simulating a patient's test result
	patientDiagnosis := "Patient P001 flu positive"

	// Encode patient report using the morse code map
	encodedDiagnosis := textToMorse(patientDiagnosis, customAlphabetMap)
	fmt.Println("Encoded Diagnosis:", encodedDiagnosis)

	// Simulating transmission or storage of encoded test result
	// For decoding, decode received encoded test result
	receivedEncodedDiagnosis := encodedDiagnosis

	// Decode the received data using the custom alphabet map
	decodedDiagnosis := morseToText(receivedEncodedDiagnosis, customAlphabetMap)
	fmt.Println("Decoded Diagnosis:", decodedDiagnosis)

	// Check data integrity
	fmt.Println("Data Integrity Check:", checkIntegrity(patientDiagnosis, decodedDiagnosis))
}
