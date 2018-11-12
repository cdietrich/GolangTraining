package demo

import (
	"testing"
)

func TestDuplicate(t *testing.T) {
	v := Duplicate(2)
	if v != 4 {
		t.Error("Expected 4, got ", v)
	}
	v = Duplicate(9)
	if v != 18 {
		t.Error("Expected 18, got ", v)
	}
}
