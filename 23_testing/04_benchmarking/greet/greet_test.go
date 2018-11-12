package greet

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	ex := "Welcome Mr. James Bond"
	s := Greet("James Bond")
	if s != ex {
		t.Error("Expected:", ex, "Got:", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James Bond"))
	// Output:
	// Welcome Mr. James Bond
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James Bond")
	}
}
