package problems

import "testing"

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
func TestSingleNumber(t *testing.T) {
	cases := []struct {
		in       []int
		expected int
	}{
		{
			in:       []int{2, 2, 1},
			expected: 1,
		},
		{
			in:       []int{4, 1, 2, 1, 2},
			expected: 4,
		},
	}

	for _, c := range cases {
		if res := singleNumber(c.in); c.expected != res {
			t.Errorf("unexpected result (%d) of %v -> %d", res, c.in, c.expected)
		}
	}
}
