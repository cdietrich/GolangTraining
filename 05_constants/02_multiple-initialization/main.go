package main

import "fmt"

const (
	_ = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
)

const (
	Z0 = iota
	Z1 = iota
	Z2 = iota
)

func main() {
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\n", KB)
	fmt.Printf("%x\n", KB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\n", MB)
	fmt.Printf("%x\n", MB)

	fmt.Println(Z0)
	fmt.Println(Z1)
	fmt.Println(Z2)
}

