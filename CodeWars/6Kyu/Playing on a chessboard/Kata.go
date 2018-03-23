package kata

// https://www.codewars.com/kata/playing-on-a-chessboard/go

// Game ...
func Game(n int) []int {
	if n%2 == 1 {
		return []int{n * n, 2}
	}
	return []int{n * n / 2}
}
