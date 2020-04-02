package problems

import "testing"

func TestIsHappy(t *testing.T) {
	cases := []struct {
		in       int
		expected bool
	}{
		{
			in:       19,
			expected: true,
		},
		{
			in:       18,
			expected: false,
		},
		{
			in:       0,
			expected: false,
		},
		{
			in:       -1,
			expected: false,
		},
		{
			in:       1,
			expected: true,
		},
	}

	for _, c := range cases {
		if res := isHappy(c.in); c.expected != res {
			t.Errorf("unexpected result (%t) of %v -> %t", res, c.in, c.expected)
		}
	}
}

func TestIsHappyFloyd(t *testing.T) {
	cases := []struct {
		in       int
		expected bool
	}{
		{
			in:       19,
			expected: true,
		},
		{
			in:       18,
			expected: false,
		},
	}

	for _, c := range cases {
		if res := isHappyFloyd(c.in); c.expected != res {
			t.Errorf("unexpected result (%t) of %v -> %t", res, c.in, c.expected)
		}
	}
}
