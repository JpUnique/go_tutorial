package main

import (
	"fmt"
	"net/http"
	"sync"
)

// implementing WaitGroup to manage concurrent HTTP requests
var wg sync.WaitGroup

func getStatusCode(endpoint string) {
	defer wg.Done() // decrement the WaitGroup counter when the HTTP request completes
	response, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", response.StatusCode, endpoint)
	}

}

func main() {

	// define a websitelist slice of strings representing the web addresses to check the status code for
	websitelist := []string{
		"https://www.google.com",
		"https://go.dev",
		"https://www.golang.org",
		"https://www.example.com",
		"https://github.com",
	}
	// using a range based loop to iterate over the slice
	for _, web := range websitelist {

		wg.Add(3)             // increment the WaitGroup counter
		go getStatusCode(web) // launch a new goroutine to handle the HTTP request
	}
	wg.Wait() // wait for all HTTP requests to complete
	fmt.Println("All HTTP requests completed")
}
