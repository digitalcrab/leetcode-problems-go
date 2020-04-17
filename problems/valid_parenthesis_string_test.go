package problems

import "testing"

func TestCheckValidString(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{
			value:    "()",
			expected: true,
		},
		{
			value:    "(*)",
			expected: true,
		},
		{
			value:    "(*))",
			expected: true,
		},
		{
			value:    "",
			expected: true,
		},
		{
			value:    "((((*)))",
			expected: true,
		},
		{
			value:    "(()())",
			expected: true,
		},
		{
			value:    "(()*)",
			expected: true,
		},
		{
			value:    "*)",
			expected: true,
		},
	}

	for _, c := range cases {
		res := checkValidString(c.value)
		if c.expected != res {
			t.Errorf("unexpected result (%t) of %q -> %t", res, c.value, c.expected)
		}
	}
}
