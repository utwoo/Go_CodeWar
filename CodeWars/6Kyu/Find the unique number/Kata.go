package kata

// https://www.codewars.com/kata/find-the-unique-number-1/go

// FindUniq ...
func FindUniq(arr []float32) float32 {
	mapping := make(map[float32]int)

	for _, f := range arr {
		if _, ok := mapping[f]; ok {
			mapping[f]++
		} else {
			mapping[f] = 1
		}
	}

	for k, v := range mapping {
		if v == 1 {
			return k
		}
	}

	return 0

	/* Clever
	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	}
	for i, v := range arr[1:] {
		if v != arr[i] {
			return v
		}
	}
	return 0
	*/
}
