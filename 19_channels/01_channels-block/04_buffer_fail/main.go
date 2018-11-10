package main

import (
	"fmt"
)

func main() {
	// buffer, too small
	c := make(chan int, 1)
	c <- 23
	c <- 42
	fmt.Println(<-c)
	fmt.Println(<-c)
}
