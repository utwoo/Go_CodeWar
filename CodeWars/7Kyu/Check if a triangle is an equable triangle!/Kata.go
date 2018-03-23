package kata

import (
	"math"
)

// https://www.codewars.com/kata/check-if-a-triangle-is-an-equable-triangle/go

// EquableTriangle ...
func EquableTriangle(a, b, c int) bool {
	p := a + b + c
	s := p / 2
	return float64(p) == math.Sqrt(float64(s*(s-a)*(s-b)*(s-c)))

	// p := a+b+c
	// return 16*p == (p-2*a)*(p-2*b)*(p-2*c)
}
