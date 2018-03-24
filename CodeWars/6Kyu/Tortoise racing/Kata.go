package kata

// https://www.codewars.com/kata/tortoise-racing/go

// Race ...
func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	seconds := int((float32(g) / float32(v2-v1)) * 3600)
	return [3]int{seconds / 3600, (seconds % 3600) / 60, seconds % 60}
}
