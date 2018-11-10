package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(ctx, cancel)
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				{
					fmt.Println(n)
					n++
					time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				}
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())
	time.Sleep(time.Second * 2)
	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())

}
