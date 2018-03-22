package kata

import (
	"strings"
)

// https://www.codewars.com/kata/mumbling/go

// Accum ...
func Accum(s string) string {
	result := make([]string, len(s))
	for i, n := range s {
		result[i] = strings.Title(strings.Repeat(strings.ToLower(string(n)), i+1))
	}
	return strings.Join(result, "-")
}
