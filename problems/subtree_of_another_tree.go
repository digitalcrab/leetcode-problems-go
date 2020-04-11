package problems

// Given two non-empty binary trees s and t, check whether tree t has exactly the same structure
// and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all
// of this node's descendants. The tree s could also be considered as a subtree of itself.
//
// Example 1:
// Given tree s:
//
//     3
//    / \
//   4   5
//  / \
// 1   2
// Given tree t:
//   4
//  / \
// 1   2
// Return true, because t has the same structure and node values with a subtree of s.
//
// Example 2:
// Given tree s:
//
//     3
//    / \
//   4   5
//  / \
// 1   2
//    /
//   0
// Given tree t:
//   4
//  / \
// 1   2
// Return false.

type isSubtreeNode struct {
	Val   int
	Left  *isSubtreeNode
	Right *isSubtreeNode
}

// Time complexity: O(s*t) - we traverse S and for every node in S we traverse T
// Memory complexity: O(s) - the depth of recursion could go till the end of the S tree
func isSubtree(s *isSubtreeNode, t *isSubtreeNode) bool {
	return isSubtreeTraverse(s, t)
}

func isSubtreeTraverse(s *isSubtreeNode, t *isSubtreeNode) bool {
	if s == nil {
		return false
	}
	if isSubtreeEqual(s, t) {
		return true
	}
	if isSubtreeTraverse(s.Left, t) {
		return true
	}
	return isSubtreeTraverse(s.Right, t)
}

func isSubtreeEqual(a *isSubtreeNode, b *isSubtreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && isSubtreeEqual(a.Left, b.Left) && isSubtreeEqual(a.Right, b.Right)
}
