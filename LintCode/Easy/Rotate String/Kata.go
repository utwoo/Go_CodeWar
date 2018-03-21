package kata

// Kata http://www.lintcode.com/en/problem/rotate-string/
func rotateString(str *string, offset int) {
	// write your code here
	val := *str
	if n := len(val); n > 0 {
		*str = val[n-offset%n:] + val[0:n-offset%n]
	}
}
