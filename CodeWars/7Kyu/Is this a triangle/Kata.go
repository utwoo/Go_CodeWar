package kata

import (
	"sort"
)

// https://www.codewars.com/kata/is-this-a-triangle/go

// IsTriangle ...
func IsTriangle(a, b, c int) bool {
	list := []int{a, b, c}
	sort.Ints(list)
	return list[0]+list[1] > list[2]

	// return a+b > c && b+c > a && a+c > b
}
