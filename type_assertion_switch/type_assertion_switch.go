package main

import (
	"fmt"
)

// type switch which permit several cases for the same value
//
//	the
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*3)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!", v)
	}
}

func main() {
	// Type assertion example
	//A type assertion provides access to an interface value's underlying concrete value.
	var i interface{} = "2"

	s := i.(string)
	fmt.Println(s)

	// the i which is an interface hold a string value, then ok give a  boolean value indicating whether the type assertion was successful (true or false).
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	do(21)
	do("hello")
	do(true)
}

// type switch example
