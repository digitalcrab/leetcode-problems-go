package problems

import "testing"

func TestFindMaxLengthOfContiguous(t *testing.T) {
	cases := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{0, 1},
			expected: 2,
		},
		{
			arr:      []int{0, 1, 0},
			expected: 2,
		},
		{
			arr:      []int{1, 0, 1, 1, 1, 0, 0},
			expected: 6,
		},
		{
			arr:      []int{1, 1, 1, 1},
			expected: 0,
		},
	}

	for i, c := range cases {
		if res := findMaxLengthOfContiguous(c.arr); res != c.expected {
			t.Errorf("#%d, unexpected result %d of %v -> %d", i, res, c.arr, c.expected)
		}
	}
}
