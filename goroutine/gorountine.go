package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// ASYNCHRONOUS FUNCTIONALITY: GOROUTINES in unpredicted order
func task1() {
	fmt.Println("Task 1 has started")
	time.Sleep(1 * time.Second)
	fmt.Println("Task 1 has finished")

}

func task2() {
	fmt.Println("Task 2 started")
	time.Sleep(1 * time.Second)
	fmt.Println("Task 2 has finished")
}

func main() {
	go say("world")
	say("hello")

	go task1()
	go task2()

	time.Sleep(3 * time.Second)
}

// time.Sleep() ensures that the main function says active long enough for tbe goroutines to finish.
