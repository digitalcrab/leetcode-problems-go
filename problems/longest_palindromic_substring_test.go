package problems

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		input string
		oneOf map[string]struct{}
	}{
		{
			input: "babad",
			oneOf: map[string]struct{}{"bab": {}, "aba": {}},
		},
		{
			input: "cbbd",
			oneOf: map[string]struct{}{"bb": {}},
		},
		{
			input: "racecar",
			oneOf: map[string]struct{}{"racecar": {}},
		},
	}

	for _, c := range cases {
		res := longestPalindrome(c.input)
		if _, found := c.oneOf[res]; !found {
			t.Errorf("unexpected result %q of %q expects one of %#v", res, c.input, c.oneOf)
		}
	}
}
