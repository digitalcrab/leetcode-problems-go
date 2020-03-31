package problems

type node0 struct {
	Val  int
	Next *node0
}

// Time complexity : O(n + m)
// Space complexity : O(n + m)
//
// The first call to mergeTwoLists does not return until the ends of both l1 and l2 have been reached,
// so n + mn+m stack frames consume O(n + m) space.
//
func MergeTwoListsRecursion(l1 *node0, l2 *node0) *node0 {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = MergeTwoListsRecursion(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoListsRecursion(l2.Next, l1)
		return l2
	}
}

// Time complexity : O(n + m)
// Space complexity : O(1)
//
func MergeTwoListsIteration(l1 *node0, l2 *node0) *node0 {
	var smaller, bigger, newHead *node0

	if l1.Val <= l2.Val {
		smaller, bigger = l1, l2
	} else {
		smaller, bigger = l2, l1
	}

	newHead = smaller

	for smaller.Next != nil {
		if smaller.Next.Val > bigger.Val {
			smaller.Next, bigger = bigger, smaller.Next
		} else {
			smaller = smaller.Next
		}
	}

	smaller.Next = bigger

	return newHead
}
