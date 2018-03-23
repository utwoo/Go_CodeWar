package kata

// https://www.codewars.com/kata/piano-kata-part-1/go

// BlackOrWhiteKey ...
func BlackOrWhiteKey(keyPressCount int) string {
	blackKeys := []int{
		1,
		4, 6,
		9, 11, 13,
		16, 18,
		21, 23, 25,
		28, 30,
		33, 35, 37,
		40, 42,
		45, 47, 49,
		52, 54,
		57, 59, 61,
		64, 66,
		69, 71, 73,
		76, 78,
		81, 83, 85}

	for _, n := range blackKeys {
		if (keyPressCount-1)%88 == n {
			return "black"
		}
	}
	return "white"

	// keymap := [12]string{
	// 	"white",
	// 	"black",
	// 	"white",
	// 	"white",
	// 	"black",
	// 	"white",
	// 	"black",
	// 	"white",
	// 	"white",
	// 	"black",
	// 	"white",
	// 	"black",
	// }

	// return keymap[(keyPressCount-1)%88%12]
}
