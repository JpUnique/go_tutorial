package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := Vertex{12, 13}
	fmt.Println(Abs(v))
}
