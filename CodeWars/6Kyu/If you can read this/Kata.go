package kata

import (
	"strings"
)

// https://www.codewars.com/kata/if-you-can-read-this-dot-dot-dot/go

// ToNato ...
func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}

	splitedWords := strings.Split(strings.ToLower(strings.Replace(words, " ", "", -1)), "")
	for i := 0; i < len(splitedWords); i++ {
		splitedWords[i]
	}

	return ""
}
