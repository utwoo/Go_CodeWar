package kata

// https://www.codewars.com/kata/piano-kata-part-2/go

// WhichNote ...
func WhichNote(keyPressCount int) string {
	keymap := [12]string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	return keymap[(keyPressCount-1)%88%12]
}
