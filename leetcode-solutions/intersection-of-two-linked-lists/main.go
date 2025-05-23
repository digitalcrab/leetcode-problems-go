package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(m + n)
// Space Complexity: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// if lists have different lengths then total lengths going through both is m + n
	// (m is number of noted in A, n number of nodes in B)
	// so we let both pointers go through  A and B then they eventually
	// walk the same distance m + n

	a, b := headA, headB

	for a != b {
		if a != nil {
			a = a.Next
		} else {
			// reached the end, lets switch to B
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			// reached the end, switch to A
			b = headA
		}
	}

	return a // or B
}

// Time Complexity: O(m + n)
// Space Complexity: O(m)
func getIntersectionNodeWithMap(headA, headB *ListNode) *ListNode {
	seen := make(map[*ListNode]bool)
	for currentA := headA; currentA != nil; currentA = currentA.Next {
		seen[currentA] = true
	}
	for currentB := headB; currentB != nil; currentB = currentB.Next {
		if seen[currentB] {
			return currentB
		}
	}
	return nil
}

func main() {

}
