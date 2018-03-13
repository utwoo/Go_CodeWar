package kata

// Kata https://www.codewars.com/kata/is-my-friend-cheating/go
func Kata(n uint64) [][2]uint64 {
	var result [][2]uint64
	var sum = n * (n + 1) / 2

	for i := uint64(1); i <= n; i++ {
		if (sum-i)%(i+1) == 0 {
			m := (sum - i) / (i + 1)
			if m < n {
				result = append(result, [2]uint64{i, m})
			}
		}
	}
	return result
}
