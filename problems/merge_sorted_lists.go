package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.
//
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoListsIteration(l1, l2)
}

// Time complexity : O(n + m)
// Space complexity : O(n + m)
//
// The first call to mergeTwoLists does not return until the ends of both l1 and l2 have been reached,
// so n + mn+m stack frames consume O(n + m) space.
//
func mergeTwoListsRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecursion(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsRecursion(l2.Next, l1)
		return l2
	}
}

// Time complexity : O(n + m)
// Space complexity : O(1)
//
func mergeTwoListsIteration(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	prev := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}

		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return head.Next
}
