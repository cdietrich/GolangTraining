// Package mymath contains useful math functions

package mymath

import "testing"

func TestMySum(t *testing.T) {
	sum := MySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if sum != 55 {
		t.Error("Expected 55, got", sum)
	}
	sum = MySum()
	if sum != 0 {
		t.Error("Expected 0, got", 0)
	}

}
