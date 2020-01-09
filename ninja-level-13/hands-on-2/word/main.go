// Package word contains functions to count words on large texts.
package word

import "strings"

// UseCount returns a map containing the number of times each word appears within the text contained in string s.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words within the text contained withing the string s.
func Count(s string) int {
	// write the code for this func
	return len(strings.Fields(s))
}
