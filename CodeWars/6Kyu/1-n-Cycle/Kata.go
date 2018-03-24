package kata

// https://www.codewars.com/kata/1-slash-n-cycle/go
func cycle(n int) int {
	result := 1

	for count := 1; count < n; count++ {
		result = (result * 10) % n
		if result == 0 {
			return -1
		} else if result == 1 {
			return count
		}
	}
	return -1
}
