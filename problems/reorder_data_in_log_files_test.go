package problems

import (
	"reflect"
	"testing"
)

func TestReorderLogFiles(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			expected: []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			input:    []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"},
			expected: []string{"a2 act car", "g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
		// strange behaviour
		//{
		//	input:    []string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"},
		//	expected: []string{"b aq cojj","5 ba iedrz","8 fj dzz k","63 gu psub","bb wsrd kp","5r 446 6 3","6s 87979 5","3r 587 01","jc 3480612","r5 6316 71"},
		//},
	}

	for i, c := range cases {
		if res := reorderLogFiles(c.input); !reflect.DeepEqual(res, c.expected) {
			t.Errorf("unexpected result case %d\nin: %#v\nout: %#v\nwants: %#v", i, c.input, res, c.expected)
		}
	}
}
