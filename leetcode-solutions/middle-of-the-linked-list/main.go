package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n/2) = O(n) where n is number of nodes in the list
// Space Complexity: O(1) as we do not create extra space besides 2 pointers
func middleNode(head *ListNode) *ListNode {
	// list is empty?
	if head == nil {
		return nil
	}

	turtle := head // slow pointer (moves 1 step)
	rabbit := head // fast pointer (moves 2 steps)

	// go over the list till the fast pointer reaches the end
	for rabbit != nil && rabbit.Next != nil {
		turtle = turtle.Next
		rabbit = rabbit.Next.Next
	}

	// slow should endue in the middle
	return turtle
}

func main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	fromTheMiddle := middleNode(list)

	for {
		fmt.Printf("%d -> ", fromTheMiddle.Val)
		fromTheMiddle = fromTheMiddle.Next
		if fromTheMiddle == nil {
			fmt.Println()
			break
		}
	}
}
