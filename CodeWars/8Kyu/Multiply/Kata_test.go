package multiply

import "testing"

func TestKata(t *testing.T) {
	if Kata(1, 1) != 1 {
		t.Error("Error")
	}
}
