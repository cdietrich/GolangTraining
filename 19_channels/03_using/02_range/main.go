package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	sum := 0
	for v := range c {
		sum += v
	}
	fmt.Println(sum)
}
