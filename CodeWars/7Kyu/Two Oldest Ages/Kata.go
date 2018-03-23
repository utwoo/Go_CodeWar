package kata

import (
	"sort"
)

// https://www.codewars.com/kata/two-oldest-ages-1/go

// TwoOldestAges ...
func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}
