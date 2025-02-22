package main

import (
    "fmt"
    
)

// ch <- v send v to  channel  ch for communication between goroutines
// v := <- ch receive value from channel ch and assign it to v
// ch := make(chan int) create a channel


func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	fmt.Println("Channels in golang are goroutine-safe")

	// creating channels in main function
	myCh := make(chan int)

	myCh <- 10 // send value to channel
	go func() { // create a new goroutine to receive value from channel
        fmt.Println(<-myCh)
    }()
}

