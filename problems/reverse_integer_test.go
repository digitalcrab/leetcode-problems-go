package problems

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	cases := []struct {
		value    int
		expected int
	}{
		{
			value:    123,
			expected: 321,
		},
		{
			value:    -123,
			expected: -321,
		},
		{
			value:    120,
			expected: 21,
		},
	}

	for _, c := range cases {
		res := ReverseInteger(c.value)
		if c.expected != res {
			t.Errorf("unexpected result (%d) of %d -> %d", res, c.value, c.expected)
		}
	}
}
