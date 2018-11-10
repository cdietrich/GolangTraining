package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	go send(eve, odd, quit)
	receive(eve, odd, quit)
}

func receive(eve, odd <-chan int, quit <-chan bool) {
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
		case i, ok := <-quit:
			{
				if !ok {
					fmt.Println("NOT OK", i)
					return
				} else {
					fmt.Println("OK", i)
				}

			}
		}
	}
}

func send(eve, odd chan<- int, quit chan<- bool) {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}
