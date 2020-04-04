package problems

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			nums:     []int{1, 2, 3, 0, 5, 0, 6, 10},
			expected: []int{1, 2, 3, 5, 6, 10, 0, 0},
		},
	}

	for _, c := range cases {
		input := make([]int, len(c.nums))
		copy(input, c.nums)
		if moveZeroes(input); !reflect.DeepEqual(input, c.expected) {
			t.Errorf("unexpected result %v of %v -> %v", input, c.nums, c.expected)
		}
	}
}
