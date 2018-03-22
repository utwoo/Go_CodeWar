package kata

import (
	"math"
)

// https://www.codewars.com/kata/to-square-root-or-not-to-square-root/go

// SquareOrSquareRoot ...
func SquareOrSquareRoot(arr []int) []int {
	result := make([]int, 0, len(arr))
	for _, n := range arr {
		num := math.Sqrt(float64(n))
		if num == math.Floor(num) {
			result = append(result, int(num))
		} else {
			result = append(result, n*n)
		}
	}
	return result
}
