package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// ctx is the background context for Redis operations
var ctx = context.Background()

func main() {
	// Initialize a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Update with your Redis server address
		Password: "",               // No password set by default
		DB:       0,                // Use default DB
	})

	// Defer the client's close to ensure it's closed after main function execution.
	defer func() {
		if err := rdb.Close(); err != nil {
			panic(err)
		}
	}()

	for {
		var choice int
		fmt.Println("Select an option:")
		fmt.Println("1. Create")
		fmt.Println("2. Read")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. Display All")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			//Create
			// Get user input for patient details
			var id, fname, lname, mhistory, gender string
			var age int

			fmt.Print("Enter patient ID: ")
			fmt.Scanln(&id)

			fmt.Print("Enter patient first name: ")
			fmt.Scanln(&fname)

			fmt.Print("Enter patient last name: ")
			fmt.Scanln(&lname)

			fmt.Print("Enter patient age: ")
			fmt.Scanln(&age)

			fmt.Print("Enter patient medical history: ")
			fmt.Scanln(&mhistory)

			fmt.Print("Enter patient gender: ")
			fmt.Scanln(&gender)

			// Create a map with user-input patient details
			patient := map[string]interface{}{
				"id":       id,
				"fname":    fname,
				"lname":    lname,
				"age":      age,
				"mhistory": mhistory,
				"gender":   gender,
			}

			// Perform creation operation in Redis
			err := rdb.HMSet(ctx, fmt.Sprintf("patient:%s", id), patient).Err()
			if err != nil {
				panic(err)
			}
			fmt.Println("Patient created successfully!")
		case 2:
			// Read
			var idToRead string

			fmt.Print("Enter patient ID to read details: ")
			fmt.Scanln(&idToRead)

			// Read patient details from Redis based on the user-provided ID
			val, err := rdb.HGetAll(ctx, fmt.Sprintf("patient:%s", idToRead)).Result()
			if err != nil {
				panic(err)
			}

			if len(val) == 0 {
				fmt.Println("Patient with given ID not found.")
			} else {
				fmt.Println("Patient details:", val)
			}
		case 3:
			// Update
			var patientID, lname string

			fmt.Print("Enter patient ID: ")
			fmt.Scanln(&patientID)

			fmt.Print("Enter last name: ")
			fmt.Scanln(&lname)

			updateData := map[string]interface{}{
				"lname": lname,
				// Add more fields to update if needed
			}

			err := rdb.HMSet(ctx, "patient:"+patientID, updateData).Err()
			if err != nil {
				panic(err)
			}
			fmt.Println("Patient updated successfully!")
		case 4:
			// Delete
			var patientID string

			fmt.Print("Enter patient ID to delete: ")
			fmt.Scanln(&patientID)

			err := rdb.Del(ctx, "patient:"+patientID).Err()
			if err != nil {
				panic(err)
			}
			fmt.Println("Patient Record deleted successfully!")
		case 5:
			// Display all patient data
			keys, err := rdb.Keys(ctx, "patient:*").Result()
			if err != nil {
				panic(err)
			}

			if len(keys) == 0 {
				fmt.Println("No patient records found.")
			} else {
				fmt.Println("Patient records:")
				for _, key := range keys {
					val, err := rdb.HGetAll(ctx, key).Result()
					if err != nil {
						panic(err)
					}
					fmt.Printf("Patient ID: %s - Details: %v\n", key, val)
				}
			}
		case 6:
			// Exit
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
