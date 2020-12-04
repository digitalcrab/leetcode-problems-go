package problems

import (
	"math/rand"
)

// Given a singly linked list, return a random node's value from the linked list.
// Each node must have the same probability of being chosen.
//
// Follow up:
// What if the linked list is extremely large and its length is unknown to you?
// Could you solve this efficiently without using extra space?
//
// Example:
//
// // Init a singly linked list [1,2,3].
//
// ListNode head = new ListNode(1);
// head.next = new ListNode(2);
// head.next.next = new ListNode(3);
// Solution solution = new Solution(head);
//
// // getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
// solution.getRandom();

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type LinkedListRandomNodeSolution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func LinkedListRandomNodeConstructor(head *ListNode) *LinkedListRandomNodeSolution {
	return &LinkedListRandomNodeSolution{
		head: head,
	}
}

/** Returns a random node's value. */
func (s *LinkedListRandomNodeSolution) GetRandom() int {
	if s.head == nil {
		return 0
	}
	if s.head.Next == nil {
		return s.head.Val
	}

	res := s.head.Val

	current := s.head.Next
	n := 2
	for current != nil {
		// change result with probability 1/n
		if rand.Intn(n) == 0 {
			res = current.Val
		}

		current = current.Next
		n++
	}

	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
