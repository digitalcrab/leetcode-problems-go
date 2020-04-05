package problems

import "testing"

func TestMaxProfit2(t *testing.T) {
	cases := []struct {
		value    []int
		expected int
	}{
		{
			value:    []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			value:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			value:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, c := range cases {
		res := maxProfit2(c.value)
		if c.expected != res {
			t.Errorf("unexpected result (%d) of %v -> %d", res, c.value, c.expected)
		}
	}
}
