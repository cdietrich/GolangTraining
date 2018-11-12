package mymath

import (
	"fmt"
	//"github.com/cdietrich/GolangTraining/23_testing/03_example_tests/mymath"
	//"github.com/cdietrich/GolangTraining/23_testing/03_example_tests/mymath"
)

func ExampleMySum() {
	fmt.Println(MySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(MySum(1, 0, -1))
	// Output:
	// 55
	// 0
}
