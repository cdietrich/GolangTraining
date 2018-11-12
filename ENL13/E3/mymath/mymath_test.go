package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type data struct {
		inputs []int
		result float64
	}
	tests := []data{
		{[]int{1, 10, 20, 30, 1000}, 20},
		{[]int{-1, 1, 2, 3, 1000}, 2},
		{[]int{-1, 5, 1000}, 5},
	}
	for i, tt := range tests {
		t.Run("Testcase "+string(i), func(t *testing.T) {
			if got := CenteredAvg(tt.inputs); got != tt.result {
				t.Errorf("CenteredAvg() = %v, want %v", got, tt.result)
			}
		})
	}

	v := CenteredAvg([]int{1, 10, 20, 30, 1000})
	if v != 20 {
		t.Error("Expected 20, got", v)
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 10, 20, 30, 1000}))
	// Output:
	// 20
}

func BenchmarkCenteredAvg(b *testing.B) {
	d := []int{1, 10, 20, 30, 1000}
	for i := 0; i < b.N; i++ {
		CenteredAvg(d)
	}
}
