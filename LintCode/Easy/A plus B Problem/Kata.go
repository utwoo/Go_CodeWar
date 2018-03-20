package kata

// Kata http://www.lintcode.com/en/problem/a-b-problem/
func aplusb(a int, b int) int {
	for b != 0 {
		carry := a ^ b
		b = (a & b) << 1
		a = carry
	}
	return a
}
