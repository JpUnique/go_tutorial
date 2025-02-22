package main

import (
	"encoding/json" // for encoding and decoding JSON data
	"fmt"
	"os" // for printing to console
)

type User struct {
	ID          int      `json:"user_id"`        // int field with json tag
	Name        string   `json:"name,omitempty"` // omitempty means field will be omitted if its value is empty string or zero value
	Age         int      `json:"age"`            // int field
	Password    string   `json:"-"`              // hidden field
	Permissions []string `json:"roles"`          // mapping to an array of strings
}

type Person struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"is_active"`
}

// Main function to demonstrate encoding and decoding JSON data in Go
func main() {
	// Create a new User called (u) instance with the given values
	u := User{
		ID:          1,
		Name:        "John Doe",
		Age:         30,
		Password:    "password123",
		Permissions: []string{"admin", "user"},
	}

	// Marshal the User instance to JSON
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		panic(err)
	}
	fmt.Println(string(jsonBytes))

	// Unmarshal JSON back to User instance
	var u2 User // create a new User instance to store the unmarshaled data having the same structure as the original User instance
	err = json.Unmarshal(jsonBytes, &u2)
	if err != nil { // handle error if unmarshaling fails
		fmt.Println("Error unmarshaling JSON:", err)
		panic(err)
	}
	// Print the unmarshaled User instance
	fmt.Println(u2)

	// ENCODING AND DECODING JSON WITH ANOTHER STRUCTURE
	p := Person{
		Name:   "John Doe",
		Age:    30,
		Active: true,
	}
	f, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		panic(err)
	}
	defer f.Close() // close the file when we're done

	// Marshal the Person instance to JSON and write it to the file
	encoder := json.NewEncoder(f)
	err = encoder.Encode(p)
	if err != nil {
		fmt.Println("Error encoding person:", err)
		panic(err)
	}

	//DECODING JSON WITH ANOTHER STRUCTURE FROM A FILE
	f, err = os.Open("output.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&p)
	if err != nil {
		fmt.Println("Error decoding person:", err)
		panic(err)
	}
	fmt.Println(p)
}
