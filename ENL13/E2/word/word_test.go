package word

import (
	"fmt"
	"sort"
	"testing"

	"github.com/cdietrich/GoLangTraining/ENL13/E2/quote"
)

func TestCount(t *testing.T) {
	v := Count("This is a test")
	exp := 4
	if v != exp {
		t.Error("Expected:", exp, "Got:", v)
	}

	v = Count("There is no business like show business")
	exp = 7
	if v != exp {
		t.Error("Expected:", exp, "Got:", v)
	}
}

func TestUseCount(t *testing.T) {
	r := UseCount("Test Demo Demo Demo Test Hello")
	v := r["Test"]
	if v != 2 {
		t.Error("Expected 2 for 'Test', but got", v)
	}
	v = r["Demo"]
	if v != 3 {
		t.Error("Expected 3 for 'Demo', but got", v)
	}
	v = r["Hello"]
	if v != 1 {
		t.Error("Expected 1 for 'Hello', but got", v)
	}
	if v, ok := r["DoesNotExist"]; ok {
		t.Error("Did not expect any data for 'DoesNotExist' but got", v)
	}
}

func ExampleUseCount() {
	keys := []string{}
	r := UseCount("Test Demo Demo Demo Test Hello")
	for k := range r {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, r[k])
	}

	// Output:
	// Demo 3
	// Hello 1
	// Test 2
}

func ExampleCount() {
	fmt.Println(Count("There is no business like show business"))
	fmt.Println(Count(quote.SunAlso))
	// Output:
	// 7
	// 1349
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
