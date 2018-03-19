package kata

import (
	"math"
)

// Kata https://www.codewars.com/kata/return-negative/train/go
func Kata(x int) int {
	return 0 - int(math.Abs(float64(x)))
}
