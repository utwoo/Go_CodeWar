package kata

// Kata https://www.codewars.com/kata/string-repeat/go
func Kata(repititions int, value string) string {
	var result string

	for index := 1; index <= repititions; index++ {
		result = result + value
	}
	return result

	//return strings.Repeat(value, repititions)
}
