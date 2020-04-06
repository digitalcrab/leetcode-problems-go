package problems

import (
	"reflect"
	"testing"
)

func TestNumberOfIslands(t *testing.T) {
	cases := []struct {
		grid     [][]byte
		expected int
	}{
		{
			grid:     [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
			expected: 1,
		},
		{
			grid:     [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
			expected: 3,
		},
		{
			grid:     [][]byte{{'1', '0', '1', '1', '0', '1', '1'}},
			expected: 3,
		},
	}

	for i, c := range cases {
		grid := make([][]byte, len(c.grid))
		copy(grid, c.grid)

		if res := numIslands(grid); !reflect.DeepEqual(res, c.expected) {
			t.Errorf("#%d unexpected result %d of %v -> %d", i, res, c.grid, c.expected)
		}
	}
}
