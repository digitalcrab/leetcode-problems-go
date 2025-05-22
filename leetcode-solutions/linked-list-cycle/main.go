package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n) where n is number of nodes in the list
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	// define 2 pointers
	slow, fast := head, head

	// stop when fast reaches the end
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// both are there
		if slow == fast {
			return true
		}
	}

	return false
}

func main() {

}
