package problems

import "testing"

func TestBackspaceCompare(t *testing.T) {
	cases := []struct {
		S        string
		T        string
		expected bool
	}{
		{
			S:        "y#fo##f",
			T:        "y#f#o##f",
			expected: true,
		},
		{
			S:        "ab#c",
			T:        "ad#c",
			expected: true,
		},
		{
			S:        "ab##",
			T:        "c#d#",
			expected: true,
		},
		{
			S:        "a##c",
			T:        "#a#c",
			expected: true,
		},
		{
			S:        "a#c",
			T:        "b",
			expected: false,
		},
	}

	for i, c := range cases {
		if res := backspaceCompare(c.S, c.T); c.expected != res {
			t.Errorf("#%d unexpected result %t of %q vs %q -> %t", i, res, c.S, c.T, c.expected)
		}
	}
}
