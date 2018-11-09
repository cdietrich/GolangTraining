package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

}

func foo() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(1)
		fmt.Println("foo:", i)
	}
}

func bar() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(1)
		fmt.Println("bar:", i)
	}
}
