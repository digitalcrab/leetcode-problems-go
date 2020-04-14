package problems

import (
	"testing"
)

func TestStringShift(t *testing.T) {
	cases := []struct {
		s        string
		shift    [][]int
		expected string
	}{
		{
			s:        "xqgwkiqpif",
			shift:    [][]int{{1, 4}, {0, 7}, {0, 8}, {0, 7}, {0, 6}, {1, 3}, {0, 1}, {1, 7}, {0, 5}, {0, 6}},
			expected: "qpifxqgwki",
		},
		{
			// {1, 4} - ecskm
			// {0, 5} - ecskm
			// {0, 4} - mecsk
			// {1, 1} - kmecs
			// {1, 5} - kmecs
			s:        "mecsk",
			shift:    [][]int{{1, 4}, {0, 5}, {0, 4}, {1, 1}, {1, 5}},
			expected: "kmecs",
		},
		{
			s:        "abc",
			shift:    [][]int{{0, 1}, {1, 2}},
			expected: "cab",
		},
		{
			s:        "abcdefg",
			shift:    [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}},
			expected: "efgabcd",
		},
		{
			s:        "fgabcde",
			shift:    [][]int{{0, 2}},
			expected: "abcdefg",
		},
	}

	for i, c := range cases {
		if res := stringShift(c.s, c.shift); res != c.expected {
			t.Errorf("#%d unexpected result %q with %+v\nout: %q\nwants: %q", i, c.s, c.shift, res, c.expected)
		}
	}
}
