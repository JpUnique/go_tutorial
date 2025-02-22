package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// struct field can be accessed using dot notation
	v.X = 4
	v.Y = 5
	fmt.Println(v.X, v.Y)
	y := Vertex{3, 5}
	p := &y // pointer to y
	fmt.Println(p)

}
