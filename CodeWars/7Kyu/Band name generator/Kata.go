package kata

import (
	"strings"
)

// https://www.codewars.com/kata/band-name-generator/go

// bandNameGenerator ...
func bandNameGenerator(word string) string {
	if word[0] == word[len(word)-1] {
		return strings.Title(word + word[1:])
	}

	return "The " + strings.Title(word)
}
