package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(eve, odd, quit)
	receive(eve, odd, quit)
}

func receive(eve, odd, quit <-chan int) {
	for {
		select {
		case v := <-eve:
			{
				fmt.Println("Even", v)
			}
		case v := <-odd:
			{
				fmt.Println("Odd", v)
			}
		case <-quit:
			{
				return
			}
		}
	}
}

func send(eve, odd, quit chan<- int) {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}
