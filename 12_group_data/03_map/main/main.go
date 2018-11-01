package main

import (
	"fmt"
)

func main() {
	fmt.Println("BASIC")
	x := map[string]int{
		"x": 1,
		"y": 2,
		"z": 3,
	}
	fmt.Println(x)
	fmt.Println(x["x"])
	x["x"] = 10
	fmt.Println(x["x"])

	fmt.Println(x["does not exist"])

	fmt.Println("V,OK")
	var v int
	var ok bool
	v, ok = x["x"]
	fmt.Println(v, ok)
	v2, ok2 := x["does not exist"]
	fmt.Println(v2, ok2)
	if v, ok = x["y"]; ok {
		fmt.Println(v)
	}
	fmt.Println("NEW")
	fmt.Println(x["new"])
	x["new"] = 10
	fmt.Println(x["new"])
	fmt.Println("RANGE")
	for key, value := range x {
		fmt.Println(key, "->", value)
	}

	fmt.Println("DELETE")
	delete(x, "x")
	delete(x, "y")
	delete(x, "z")

	for key, value := range x {
		fmt.Println(key, "->", value)
	}
}
