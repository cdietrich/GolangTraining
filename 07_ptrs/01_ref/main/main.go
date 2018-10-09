package main

import "fmt"

func main() {
	var a = 42
	fmt.Println(a)
	fmt.Println(&a)
	var b *int = &a
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Printf("type of b %t %q", b, b)
}
