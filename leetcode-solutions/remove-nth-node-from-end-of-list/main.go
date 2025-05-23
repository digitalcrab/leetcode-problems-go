package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Example:
	// Input: head = [1,2,3,4,5], n = 2

	// create wrapper node, so we can handle head removal, and always have next
	wrapper := &ListNode{Next: head}
	end, start := wrapper, wrapper
	// [_,1,2,3,4,5]
	//  ^ end, start

	// now move the end to `n` nodes ahead
	for i := 0; i < n; i++ {
		end = end.Next
	}
	// [_,1,2,3,4,5]
	//      ^ end

	// ok, now lets move both pointers simultaneously
	// end should reach the "end" sooner as we moved it ahead
	for end.Next != nil {
		end = end.Next
		start = start.Next
	}
	// [_,1,2,3,4,5]
	//        ^ end
	//    ^ start
	// [_,1,2,3,4,5]
	//          ^ end
	//      ^ start
	// [_,1,2,3,4,5]
	//            ^ end
	//        ^ start

	// now start should point into the node before the one we need to remove
	start.Next = start.Next.Next

	return wrapper.Next
}

func main() {

}
