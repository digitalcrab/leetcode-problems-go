package problems

import (
	"testing"

	"github.com/digitalcrab/leetcode-problems-go/structures"
)

func TestMiddleOfTheLinkedList(t *testing.T) {
	cases := []struct {
		head     *structures.LinkedList
		expected *structures.LinkedList
	}{
		{
			head:     structures.NewLinkedList(1, 2, 3, 4, 5),
			expected: structures.NewLinkedList(3, 4, 5),
		},
		{
			head:     structures.NewLinkedList(1, 2, 3, 4, 5, 6),
			expected: structures.NewLinkedList(4, 5, 6),
		},
	}

	for _, c := range cases {
		middle := MiddleOfTheLinkedList(c.head)
		head, expected, res := c.head.String(), c.expected.String(), middle.String()
		if expected != res {
			t.Errorf("unexpected result %q of %q -> %q", res, head, expected)
		}
	}
}

func TestMiddleOfTheLinkedListSlowFast(t *testing.T) {
	cases := []struct {
		head     *structures.LinkedList
		expected *structures.LinkedList
	}{
		{
			head:     structures.NewLinkedList(1, 2, 3, 4, 5),
			expected: structures.NewLinkedList(3, 4, 5),
		},
		{
			head:     structures.NewLinkedList(1, 2, 3, 4, 5, 6),
			expected: structures.NewLinkedList(4, 5, 6),
		},
	}

	for _, c := range cases {
		middle := MiddleOfTheLinkedListSlowFast(c.head)
		head, expected, res := c.head.String(), c.expected.String(), middle.String()
		if expected != res {
			t.Errorf("unexpected result %q of %q -> %q", res, head, expected)
		}
	}
}
