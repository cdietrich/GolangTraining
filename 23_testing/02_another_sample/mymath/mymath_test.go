// Package mymath contains useful math functions

package mymath

import (
	"testing"
)

// this test is a table test
func TestMySum(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "Sum of 1 to 10",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   55,
		},
		{
			name:   "Empty Sum",
			values: []int{},
			want:   0,
		},
		{
			name:   "Empty Sum",
			values: []int{1, 2},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySum(tt.values...); got != tt.want {
				t.Errorf("MySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
