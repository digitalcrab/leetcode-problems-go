package problems

import "testing"

func TestIsValidParentheses(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{
			value:    "()",
			expected: true,
		},
		{
			value:    "()[]{}",
			expected: true,
		},
		{
			value:    "(]",
			expected: false,
		},
		{
			value:    "([)]",
			expected: false,
		},
		{
			value:    "{[]}",
			expected: true,
		},
		{
			value:    "]",
			expected: false,
		},
	}

	for _, c := range cases {
		res := IsValidParentheses(c.value)
		if c.expected != res {
			t.Errorf("unexpected result (%t) of %q -> %t", res, c.value, c.expected)
		}
	}
}
