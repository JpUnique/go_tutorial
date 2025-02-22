package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// PARSE JSON RESPONSE
type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// Send a GET request to the given URL
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil { // check for errors
		panic(err)
	}
	defer response.Body.Close() // close the response body when we're done

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body)) // print the response body as a string

	// PARSE JSON RESPONSE
	var post Post
	decoder := json.NewDecoder(response.Body) // use io.Reader directly
	err = decoder.Decode(&post)               // Decode JSON into struct
	if err != nil {
		panic(err)
	}
	fmt.Printf("Post Title: %s\n", post.Title)

}
