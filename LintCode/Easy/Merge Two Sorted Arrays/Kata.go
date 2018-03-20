package kata

// Kata http://www.lintcode.com/en/problem/merge-two-sorted-arrays/
func mergeSortedArray(A []int, B []int) []int {
	var (
		result = make([]int, 0, len(A)+len(B))
		a      int
		b      int
	)
	for {
		if A[a] < B[b] {
			result = append(result, A[a])
			a++
			if a == len(A) {
				result = append(result, B[b:]...)
				break
			}
		} else {
			result = append(result, B[b])
			b++
			if b == len(B) {
				result = append(result, A[a:]...)
				break
			}
		}
	}

	return result
}
