package word

import "strings"

// UseCount counts words by frequency
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count counts the number of words in given string
func Count(s string) int {
	return len(strings.Fields(s))
}
