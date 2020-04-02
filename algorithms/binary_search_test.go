package algorithms

import "testing"

func TestBinarySearchIterative(t *testing.T) {
	cases := []struct {
		input    []int
		x        int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 4, 6, 7, 8, 9, 10},
			x:        5,
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 4, 6, 7, 8, 9, 10},
			x:        9,
			expected: true,
		},
	}

	for _, c := range cases {
		if res := binarySearchIterative(c.input, c.x); c.expected != res {
			t.Errorf("unexpected result (%t) of %v (%d) -> %t", res, c.input, c.x, c.expected)
		}
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	cases := []struct {
		input    []int
		x        int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 4, 6, 7, 8, 9, 10},
			x:        5,
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 4, 6, 7, 8, 9, 10},
			x:        9,
			expected: true,
		},
	}

	for _, c := range cases {
		if res := binarySearchRecursive(c.input, c.x); c.expected != res {
			t.Errorf("unexpected result (%t) of %v (%d) -> %t", res, c.input, c.x, c.expected)
		}
	}
}
