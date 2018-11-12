package main

import (
	"fmt"
	"github.com/cdietrich/GolangTraining/02_fundamentals/util"
)

func main() {
	util.DoIt("now")
	fmt.Println(util.MyVar)
	var i = 1
	fmt.Printf("%t %d\n", i, i)
	b := i + 1
	b = b + 1
	fmt.Printf("%T %v\n", b, b)
	var c bool
	fmt.Printf("%T %v\n", c, c)
	{
		var d = 2
		fmt.Println(d)
	}
	//fmt.Println(d)
	fmt.Println(max)
	max := max(1, 2)
	fmt.Println(max)
}

func max(a int, b int) int {
	if b > a {
		return b
	} else {
		return a
	}
}
