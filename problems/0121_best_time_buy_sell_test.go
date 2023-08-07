package problems

import "testing"

func TestBestTimeToBuyAndSellStock121(t *testing.T) {
	cases := []struct {
		value    []int
		expected int
	}{
		// Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
		// Not 7-1 = 6, as selling price needs to be larger than buying price.
		{
			value:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		// In this case, no transaction is done, i.e. max profit = 0.
		{
			value:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, c := range cases {
		res := BestTimeToBuyAndSellStock121(c.value)
		if c.expected != res {
			t.Errorf("unexpected result (%d) of %v -> %d", res, c.value, c.expected)
		}
	}
}
