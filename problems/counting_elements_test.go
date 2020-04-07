package problems

import (
	"testing"
)

func TestCountElements(t *testing.T) {
	cases := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{1, 2, 3},
			expected: 2,
		},
		{
			arr:      []int{1, 1, 3, 3, 5, 5, 7, 7},
			expected: 0,
		},
		{
			arr:      []int{1, 3, 2, 3, 5, 0},
			expected: 3,
		},
	}

	for i, c := range cases {
		if res := countElements(c.arr); res != c.expected {
			t.Errorf("#%d, unexpected result %d of %v -> %d", i, res, c.arr, c.expected)
		}
	}
}
