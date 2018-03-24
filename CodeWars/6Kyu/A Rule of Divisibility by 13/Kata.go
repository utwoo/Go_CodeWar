package kata

import (
	"strconv"
	"strings"
)

// https://www.codewars.com/kata/a-rule-of-divisibility-by-13/go
var pattern = [6]int{1, 10, 9, 12, 3, 4}

// Thirt ...
func Thirt(n int) int {
	strs := strings.Split(strconv.Itoa(n), "")
	total := 0
	for i := 0; i < len(strs); i++ {
		num, _ := strconv.Atoi(strs[len(strs)-i-1])
		total += num * pattern[i%6]
	}
	if total == n {
		return total
	}

	return Thirt(total)
}
