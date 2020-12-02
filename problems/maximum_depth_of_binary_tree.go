package problems

import "math"

// Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
// Example 1:
//       3
//      /  \
//     9   20
//         / \
//        15  7
//
// Input: root = [3,9,20,null,null,15,7]
// Output: 3
//
// Example 2:
// Input: root = [1,null,2]
// Output: 2
//
// Example 3:
// Input: root = []
// Output: 0
//
// Example 4:
// Input: root = [0]
// Output: 1
//
// Constraints:
//
// The number of nodes in the tree is in the range [0, 104].
// -100 <= Node.val <= 100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type maxDepthTreeNode struct {
	Val   int
	Left  *maxDepthTreeNode
	Right *maxDepthTreeNode
}

func maxDepth(root *maxDepthTreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1 // add one for the current node
}
