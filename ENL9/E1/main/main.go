package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go a()
	go b()
	fmt.Println("main")
	wg.Wait()
}

func a() {
	fmt.Println("a")
	wg.Done()
}

func b() {
	fmt.Println("b")
	wg.Done()
}
