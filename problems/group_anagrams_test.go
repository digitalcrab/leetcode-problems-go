package problems

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		in       []string
		expected [][]string
	}{
		{
			in:       []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
	}

	for _, c := range cases {
		if res := groupAnagrams(c.in); !reflect.DeepEqual(res, c.expected) {
			t.Logf("unexpected result %#v of %#v -> %#v", res, c.in, c.expected)
		}
	}
}
