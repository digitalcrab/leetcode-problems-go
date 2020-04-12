package problems

import (
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	cases := []struct {
		stones   []int
		expected int
	}{
		{
			stones:   []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
	}

	for i, c := range cases {
		if res := lastStoneWeight(c.stones); res != c.expected {
			t.Errorf("#%d unexpected result %d of %+v -> %d", i, res, c.stones, c.expected)
		}
	}
}
