package problems

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
	}

	for i, c := range cases {
		if res := productExceptSelf(c.nums); !reflect.DeepEqual(res, c.expected) {
			t.Errorf("unexpected result case %d\nin: %#v\nout: %#v\nwants: %#v", i, c.nums, res, c.expected)
		}
	}
}
