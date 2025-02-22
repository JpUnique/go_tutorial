package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Reader in golang is called io.Reader interface, it is  standard way
// to read streams of data from various sources like files, network connections, buffer etc.

func main() {
	// Reading
	data := "Hello, World!"
	reader := strings.NewReader(data) // create an io.Reader from a string
	// Read the data from the reader
	buffer := make([]byte, 8) // read in chunks of 8 bytes

	for {
		n, err := reader.Read(buffer)
		// EOF means end of file an error returned by functions that read from a file
		//stream or buffer when there is no more data left to read
		if err == io.EOF {
			break
		}
		// print the chunk of data
		fmt.Print(string(buffer[:n]))
	}

	file, err := os.Open("example.txt") // Open a file
	if err != nil {
		panic(err)
	}
	defer file.Close() // Ensure file closes after function ends

	buffer = make([]byte, 16) // Read in chunks of 16 bytes
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Print(string(buffer[:n])) // Convert bytes to string
	}

}
