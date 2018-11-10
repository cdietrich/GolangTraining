package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 23
	}()
	fmt.Println(<-c)
}
