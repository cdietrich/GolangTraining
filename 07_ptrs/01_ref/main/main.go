package main

import "fmt"

func main() {
	var a = 42
	fmt.Println(a)
	fmt.Println(&a)
	var b *int = &a
	fmt.Println(b)
	fmt.Println(&b)
	fmt.Printf("type of b %t %q\n", b, b)

	fmt.Printf("%T\n", &b)
}
