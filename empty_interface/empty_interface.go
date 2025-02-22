package main

import "fmt"

// printAnything accepts any type of value as input and prints it out.
func printAnything(value interface{}) {
	fmt.Println(value)
}

func main() {
	printAnything(42)
	printAnything("Hello, World!")
	printAnything(true)
	printAnything([]int{1, 2, 3, 4, 5})
}
