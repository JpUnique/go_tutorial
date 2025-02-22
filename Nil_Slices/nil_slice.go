package main

import "fmt"

func main() {
	var ebuka []int
	//ebuka = append(ebuka, 1, 2, 3, 4)

	fmt.Println(ebuka, len(ebuka), cap(ebuka))
	if ebuka == nil {  // tautological condition
		fmt.Println("nil!")
	}

	//fmt.Println(ebuka)
}
