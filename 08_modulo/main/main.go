package main

import "fmt"

func main() {
	var a = 43
	var b = a % 3
	var c = a / 3
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(c*3 + b)

	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
