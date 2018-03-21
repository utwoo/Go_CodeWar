package kata

// Kata http://www.lintcode.com/en/problem/strstr/
// http://blog.csdn.net/christ1750/article/details/51259425
func strStr(source string, target string) int {
	next := kmpnext(source)
	return kmp(source, target, next)
}

func kmp(source string, target string, next []int) int {
	for i, j := 0, 0; i < len(source); i++ {
		for j > 0 && source[i] != target[j] {
			j = next[j-1]
		}
		if source[i] == target[j] {
			j++
		}
		if j == len(target) {
			return i - j + 1
		}
	}
	return 0
}

func kmpnext(dest string) []int {
	next := make([]int, len(dest))
	next[0] = 0
	for i, j := 1, 0; i < len(dest); i++ {
		for j > 0 && dest[j] != dest[i] {
			j = next[j-1]
		}
		if dest[i] == dest[j] {
			j++
		}
		next[i] = j
	}
	return next
}
