package main

import "fmt"

func main() {
	doit("Apple")
	doit("Peanut")
	doit("Orange")

}

func doit(a string) {
	switch a {
	case "Apple", "Orange":
		fmt.Println("Juice")
	case "Peanut":
		fmt.Println("Butter")
	default:
		fmt.Println("Trash")
	}
}
