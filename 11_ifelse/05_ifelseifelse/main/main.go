package main

import "fmt"

func main() {
	if false {
		fmt.Println("Never happen")
	} else if true {
		fmt.Println("Yeha")
	} else {
		fmt.Println("WTF")
	}
}
