package main

import (
	"fmt"

	"github.com/cdietrich/GoLangTraining/23_testing/03_example_tests/mymath"
)

func main() {
	fmt.Println(mymath.MySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(mymath.MySum(1, 0, -1))
}
