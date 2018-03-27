package kata

import (
	"strings"
)

// https://www.codewars.com/kata/if-you-can-read-this-dot-dot-dot/go

// ToNato ...
func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}

	splitedWords := []rune(strings.ToLower(strings.Replace(words, " ", "", -1)))
	result := make([]string, len(splitedWords))
	for i, char := range splitedWords {
		if char >= 'a' && char <= 'z' {
			result[i] = nato[char-97]
		} else {
			result[i] = string(char)
		}
	}

	return strings.Join(result, " ")
}
