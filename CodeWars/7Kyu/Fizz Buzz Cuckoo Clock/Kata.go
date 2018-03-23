package kata

import (
	"strconv"
	"strings"
)

// https://www.codewars.com/kata/fizz-buzz-cuckoo-clock/go

// FizzBuzzCuckooClock ...
func FizzBuzzCuckooClock(time string) string {
	times := strings.Split(time, ":")
	hour, _ := strconv.Atoi(times[0])
	minute, _ := strconv.Atoi(times[1])

	if minute == 0 {
		if hour%12 == 0 {
			return strings.TrimSpace(strings.Repeat("Cuckoo ", 12))
		}
		return strings.TrimSpace(strings.Repeat("Cuckoo ", hour%12))
	} else if minute == 30 {
		return "Cuckoo"
	} else if minute%15 == 0 {
		return "Fizz Buzz"
	} else if minute%5 == 0 {
		return "Buzz"
	} else if minute%3 == 0 {
		return "Fizz"
	}

	return "tick"
}
