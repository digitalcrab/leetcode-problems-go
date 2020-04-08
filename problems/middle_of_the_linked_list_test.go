package problems

import (
	"fmt"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	cases := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			expected: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		},
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}},
			expected: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}},
		},
	}

	serialization := func(h *ListNode) string {
		top := h
		var res string
		for top != nil {
			res += fmt.Sprintf("~%d", top.Val)
			top = top.Next
		}
		return res
	}

	for _, c := range cases {
		middle := middleNode(c.head)
		head, expected, res := serialization(c.head), serialization(c.expected), serialization(middle)
		if expected != res {
			t.Errorf("unexpected result %q of %q -> %q", res, head, expected)
		}
	}
}
