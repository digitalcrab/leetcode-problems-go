package problems

import "testing"

func TestTwoSumLessThanK(t *testing.T) {
	cases := []struct {
		A        []int
		K        int
		expected int
	}{
		{
			A:        []int{34, 23, 1, 24, 75, 33, 54, 8},
			K:        60,
			expected: 58,
		},
		{
			A:        []int{10, 20, 30},
			K:        15,
			expected: -1,
		},
	}

	for i, c := range cases {
		if res := twoSumLessThanK(c.A, c.K); c.expected != res {
			t.Errorf("#%d unexpected result (%v) of %v and %d -> %v", i, res, c.A, c.K, c.expected)
		}
	}
}
