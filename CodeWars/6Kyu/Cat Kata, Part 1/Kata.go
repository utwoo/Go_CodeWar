package kata

import (
	"math"
	"strings"
)

// https://www.codewars.com/kata/cat-kata-part-1

// PeacefulYard ...
func PeacefulYard(yard []string, minDistance int) bool {
	var locats [][2]int
	for i, x := range yard {
		if strings.Contains(x, "L") {
			locats = append(locats, [2]int{i, strings.Index(x, "L")})
		}
		if strings.Contains(x, "M") {
			locats = append(locats, [2]int{i, strings.Index(x, "M")})
		}
		if strings.Contains(x, "R") {
			locats = append(locats, [2]int{i, strings.Index(x, "R")})
		}
	}

	if len(locats) < 2 {
		return true
	}

	for i := 0; i < len(locats); i++ {
		for j := i + 1; j < len(locats); j++ {
			if math.Sqrt(math.Pow(float64(locats[i][0]-locats[j][0]), 2)+math.Pow(float64(locats[i][1]-locats[j][1]), 2)) < float64(minDistance) {
				return false
			}
		}
	}

	return true

}
