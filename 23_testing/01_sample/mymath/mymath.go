// Package mymath contains useful math functions
package mymath

// MySum adds the values of the given slice and returns the sum
func MySum(xi ...int) int {
	var sum = 0
	for _, i := range xi {
		sum += i
	}
	return sum
}
