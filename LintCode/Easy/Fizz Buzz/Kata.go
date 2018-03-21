package kata

import (
	"strconv"
)

// Kata http://www.lintcode.com/en/problem/fizz-buzz/

/**
 * @param n: An integer
 * @return: A list of strings.
 */
func fizzBuzz(n int) []string {
	// write your code here
	result := make([]string, n)

	mapping := map[int]string{
		3:  "fizz",
		6:  "fizz",
		9:  "fizz",
		12: "fizz",
		5:  "buzz",
		10: "buzz",
		0:  "fizz buzz",
	}

	for i := 1; i <= n; i++ {
		if v, ok := mapping[i%15]; ok {
			result[i-1] = v
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}
