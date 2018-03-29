package kata

import (
	"strings"
)

// https://www.codewars.com/kata/consecutive-strings/go

// LongestConsec ...
func LongestConsec(strarr []string, k int) string {
	var buffer string
	var largest string

	for i := 0; i <= len(strarr)-k; i++ {
		buffer = strings.Join(strarr[i : i+k][:], "")
		if len(buffer) > len(largest) {
			largest = buffer
		}
	}
	return largest
}
