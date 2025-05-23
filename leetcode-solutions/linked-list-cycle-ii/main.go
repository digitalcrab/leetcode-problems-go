package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n)
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// start with the head and idx = 0
	slow, fast := head, head

	for fast != nil && fast.Next != nil { // O(n/2)
		slow = slow.Next
		fast = fast.Next.Next

		// here we know that we have a loop
		if slow == fast {
			// Let's say:
			// a - distance from the head to the beginning of the cycle
			// b - distance from the beginning of the cycle to the meeting point
			// c - the length of the cycle
			// The slow pointer travels a + b
			// The fast pointer travels a + b + n * c (n - number of times it completes a loop)
			// Since the fast pointer moves twice as fast as the slow pointer:
			// 2 * (a + b) = a + b + n * c
			// Simplifying:
			// 2a + 2b = a + b + n * c
			// a = n*c - b
			// therefore the distance from the head to the beginning of the cycle is the same as
			// distance from the beginning of the cycle to the meeting point

			// lets reset the slow to the head
			// and keep moving fast with the same speed
			slow = head
			for slow != fast { // O(n)
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}

func main() {

}
