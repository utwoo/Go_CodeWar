package kata

import "testing"

func TestKata(t *testing.T) {
	if Kata(26) != [][2]uint64{{15, 21}, {21, 15}} {
		t.Error("Error")
	}
}
