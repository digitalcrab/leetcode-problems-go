package problems

import (
	"github.com/digitalcrab/leetcode-problems-go/structures"
)

// https://leetcode.com/problems/middle-of-the-linked-list/
//
// Given a non-empty, singly linked list with head node head, return a middle node of linked list.
// If there are two middle nodes, return the second middle node.
//
// Example 1:
// Input: [1,2,3,4,5]
// Output: Node 3 from this list (Serialization: [3,4,5])
//   The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
//   Note that we returned a LinkedListNode object ans, such that:
//   ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.
//
// Example 2:
// Input: [1,2,3,4,5,6]
// Output: Node 4 from this list (Serialization: [4,5,6])
//   Since the list has two middle nodes with values 3 and 4, we return the second one.
//
// Note:
//  - The number of nodes in the given list will be between 1 and 100.

type ListNode struct {
	Val  int
	Next *ListNode
}

// MiddleOfTheLinkedList
// Time Complexity: O(N), where N is the number of nodes in the given list.
// Space Complexity: O(N), the space used by indexes.
func MiddleOfTheLinkedList(list *structures.LinkedList) *structures.LinkedListNode {
	head := list.Top
	indexes := make(map[int]*structures.LinkedListNode)

	for i := 0; head != nil; i++ {
		indexes[i] = head
		head = head.Next
	}

	return indexes[len(indexes)/2]
}

// MiddleOfTheLinkedListSlowFast
// Time Complexity: O(N), where N is the number of nodes in the given list.
// Space Complexity: O(1), the space used by slow and fast.
func MiddleOfTheLinkedListSlowFast(list *structures.LinkedList) *structures.LinkedListNode {
	slow, fast := list.Top, list.Top
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
