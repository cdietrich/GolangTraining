package main

import (
	"fmt"
	"log"
)

const m2y = 1.09361

func main() {
	var meters float64
	fmt.Print("Please enter number of meters: ")
	_, err := fmt.Scan(&meters)
	if err != nil {
		log.Fatal(err)
	}
	var yrds = meters * m2y
	fmt.Println(meters, "is", yrds, "yards")
}
