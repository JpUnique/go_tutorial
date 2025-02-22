package main

import "fmt"

var chioma []int

func main() {
	s := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := s[3:6]
	x := s[6:]
	a := s[:5]
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(a)

	//slices are like references to arrays
	names := [4]string{"john", "paul", "emma", "david"}
	fmt.Println(names)

	c := names[0:3]
	d := names[1:4]
	fmt.Println(c, d)

	d[0] = "ebuka"
	fmt.Println(names)
	fmt.Println(c, d)

	// slice default
	var ayo [10]int
	ayo[0] = 9
	ayo[1] = 8

	fmt.Println(ayo)

	//slice length and capacity

	chioma := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	fmt.Println(chioma)
	chioma = chioma[1:] // full capacity
	chioma = chioma[:0] // empty slice
	chioma = chioma[:5] // slice with 5 elements
	chioma = chioma[2:] // drop the first two values
	fmt.Println(chioma)
}
