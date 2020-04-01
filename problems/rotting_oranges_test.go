package problems

import (
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	cases := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid:     [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			expected: 4,
		},
		{
			grid:     [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			expected: -1,
		},
		{
			grid:     [][]int{{0, 2}},
			expected: 0,
		},
	}

	for _, c := range cases {
		if res := OrangesRotting(c.grid); res != c.expected {
			t.Errorf("unexpected result of grid %v = %d, wants %d", c.grid, res, c.expected)
		}
	}
}
