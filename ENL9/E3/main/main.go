package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var value = 0

func inc() {
	copy := value
	runtime.Gosched()
	copy++
	value = copy
	wg.Done()
}

// go run -race main.go
func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go inc()
	}

	wg.Wait()
	fmt.Println(value)
}
