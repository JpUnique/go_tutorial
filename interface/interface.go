package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{} // struct Dog implements Speaker

func (d Dog) Speak() string {
	return "Woof!"
}

type Animal interface {
	Speak() string
	Move() string
	eat() string
}
type Cat struct{}

// eat implements Animal.
func (c Cat) eat() string {
	return "food"
}

// struct Cat implements Speaker

func (c Cat) Move() string {
	return "Walks"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	var s Speaker
	s = Dog{} // Dog is assigned to the interface Speaker

	fmt.Println(s.Speak()) // Output: Woof!

	var a Animal = Cat{} // Cat is assigned to the interface Animal

	fmt.Println(" ", a.Speak(), "\n", a.Move(), "\n", a.eat()) // Output: Meow! Walks Eats
}
