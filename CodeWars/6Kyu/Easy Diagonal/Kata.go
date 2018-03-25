package kata

// https://www.codewars.com/kata/easy-diagonal/go

// Diagonal ...
func Diagonal(n, p int) int {
	if n == 0 {
		return 1
	}

	triangle := []int{1, 1}
	for i := 0; i < n; i++ {
		triangle = pascalTriangle(triangle)
	}

	return triangle[p+1]

	/* Clever
	// var z big.Int
	// return int(z.Binomial(int64(n+1), int64(p+1)).Int64())
	*/

	/* Clever
	sum := 1
	num := 1
	for i := 1; i < n-p+1; i++ {
		num = num * (p + i) / (i)
		sum += num
	}
	return sum
	*/
}

func pascalTriangle(rows []int) []int {
	result := make([]int, 0, len(rows)+1)
	result = append(result, 1)
	for i := 0; i < len(rows)-1; i++ {
		result = append(result, rows[i]+rows[i+1])
	}
	result = append(result, 1)
	return result
}
