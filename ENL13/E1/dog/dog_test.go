package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	v := Years(10)
	exp := 70
	if v != exp {
		t.Error("Expected:", exp, ". Got:", v, ".")
	}

	v = Years(1)
	exp = 7
	if v != exp {
		t.Error("Expected:", exp, ". Got:", v, ".")
	}
}
func TestYearsTwo(t *testing.T) {
	v := YearsTwo(10)
	exp := 70
	if v != exp {
		t.Error("Expected:", exp, ". Got:", v, ".")
	}

	v = YearsTwo(1)
	exp = 7
	if v != exp {
		t.Error("Expected:", exp, ". Got:", v, ".")
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
