package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// range in slice return two value  are the index and element of that index
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
