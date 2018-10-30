package problems

import "testing"

func TestMostCommonWord(t *testing.T) {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	expected := "ball"

	if val := MostCommonWord(paragraph, banned); val != expected {
		t.Errorf("Unexpected value %q, should be %q", val, expected)
	}
}
