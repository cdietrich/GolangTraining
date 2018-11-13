package main

import "fmt"

func main() {
	doit(42)
	doit(-23)
	doit(0)

}

func doit(a int) {
	switch {
	case a > 0:
		fmt.Println("Positive")
	case a < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
