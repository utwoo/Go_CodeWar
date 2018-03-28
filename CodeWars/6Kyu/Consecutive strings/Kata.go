package kata

import (
	"sort"
	"strings"
)

// https://www.codewars.com/kata/consecutive-strings/go

// StringSort ...
type StringSort []string

// Len ...
func (s StringSort) Len() int {
	return len(s)
}

// Swap ...
func (s StringSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less ...
func (s StringSort) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

// LongestConsec ...
func LongestConsec(strarr []string, k int) string {
	if len(strarr) == 0 || k > len(strarr) || k <= 0 {
		return ""
	}
	trans := StringSort(strarr)
	sort.Sort(trans)
	return strings.Join(trans[:k], "")
}
