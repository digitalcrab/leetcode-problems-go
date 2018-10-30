package problems

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		value    []int
		expected int
	}{
		{
			value:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6, // [4,-1,2,1] has the largest sum = 6
		},
		{
			value:    []int{-2, -1},
			expected: -1,
		},
		{
			value:    []int{-2},
			expected: -2,
		},
		{
			value:    []int{4, -1, 2, 1},
			expected: 6,
		},
	}

	for _, c := range cases {
		res := MaxSubArray(c.value)
		if c.expected != res {
			t.Errorf("unexpected result (%d) of %v -> %d", res, c.value, c.expected)
		}
	}
}
