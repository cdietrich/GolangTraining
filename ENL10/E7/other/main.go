package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := make(chan int, 20)
	for i := 0; i < 10; i++ {
		ix := i
		go func() {
			for j := 0; j < 10; j++ {
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				c <- 10*ix + j
			}
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}
	fmt.Println("done")
}
