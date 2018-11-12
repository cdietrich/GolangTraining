package greet

import (
	"fmt"
)

// Greet is a really nice function
func Greet(name string) string {
	return fmt.Sprint("Welcome Mr. ", name)
}
