package kata

import (
	"strings"
)

// https://www.codewars.com/kata/all-unique/go

// HasUniqueChar ...
func HasUniqueChar(str string) bool {
	for _, n := range str {
		if strings.Count(str, string(n)) >= 2 {
			return false
		}
	}
	return true
}
