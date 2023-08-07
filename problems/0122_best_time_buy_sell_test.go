package problems

import "testing"

func TestBestTimeToBuyAndSellStock122(t *testing.T) {
	cases := []struct {
		value    []int
		expected int
	}{
		// Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
		// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
		// Total profit is 4 + 3 = 7.
		{
			value:    []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		// Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
		// Total profit is 4.
		{
			value:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		// There is no way to make a positive profit,
		// so we never buy the stock to achieve the maximum profit of 0.
		{
			value:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, c := range cases {
		res := BestTimeToBuyAndSellStock122(c.value)
		if c.expected != res {
			t.Errorf("unexpected result (%d) of %v -> %d", res, c.value, c.expected)
		}
	}
}

func TestBestTimeToBuyAndSellStock122LocalMinMax(t *testing.T) {
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
		res := BestTimeToBuyAndSellStock122LocalMinMax(c.value)
		if c.expected != res {
			t.Errorf("unexpected result (%d) of %v -> %d", res, c.value, c.expected)
		}
	}
}
