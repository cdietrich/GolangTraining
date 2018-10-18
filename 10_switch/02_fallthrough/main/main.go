package main

import "fmt"

func main() {
	doit("Apple")
	doit("Peanut")
	doit("Nut")

}

func doit(a string) {
	switch a {
	case "Apple":
		fmt.Println("Pie")
		fallthrough
	case "Orange":
		fmt.Println("Juice")
	case "Peanut":
		fmt.Println("Butter")
	default:
		fmt.Println("Trash")
	}
	fmt.Println("==========")
}