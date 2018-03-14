package kata

import "strconv"

// Kata https://www.codewars.com/kata/century-from-year/go
func century(year int) int {
	strYear := ("0000" + strconv.Itoa(year))
	strYear = strYear[len(strYear)-4 : len(strYear)]
	result, _ := strconv.Atoi(strYear[0:2])

	if strYear[2:4] != "00" {
		result++
	}

	return result

	//return (year + 99) / 100
}
