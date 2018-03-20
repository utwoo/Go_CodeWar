package kata

// Kata http://www.lintcode.com/en/problem/trailing-zeros/
func trailingZeros(n int64) int64 {
	var (
		result int64
		i      int64 = 5
	)
	for ; i <= n; i *= 5 {
		result += n / i
	}
	return result
}
