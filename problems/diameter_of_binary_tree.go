package problems

import "math"

// https://leetcode.com/problems/diameter-of-binary-tree/
//
// Given a binary tree, you need to compute the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
// This path may or may not pass through the root.
//
// Example:
// Given a binary tree
//          1
//         / \
//        2   3
//       / \
//      4   5
// Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
//
// Note: The length of path between two nodes is represented by the number of edges between them.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type diameterOfBinaryTreeNode struct {
	Val   int
	Left  *diameterOfBinaryTreeNode
	Right *diameterOfBinaryTreeNode
}

// diameter could be one of the:
// 1) diameter of left subtree
// 2) diameter of the right subtree
// 3) longest path that goes through the root (calculated bales on subtree height)
func diameterOfBinaryTree(root *diameterOfBinaryTreeNode) int {
	if root == nil {
		return 0
	}

	// case 1
	leftDiameter := diameterOfBinaryTree(root.Left)

	// case 2
	rightDiameter := diameterOfBinaryTree(root.Right)

	// case 3
	leftHeight := diameterOfBinaryTreeHeight(root.Left)
	rightHeight := diameterOfBinaryTreeHeight(root.Right)
	longestHeight := leftHeight + rightHeight

	return int(math.Max(
		float64(longestHeight), // case 3
		math.Max(
			float64(leftDiameter),  // case 1
			float64(rightDiameter), // case 2
		),
	))
}

// Calc "height" of the tree, which is number of nodes from the longest
// path (left or right) from the root (or node itself)
func diameterOfBinaryTreeHeight(node *diameterOfBinaryTreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := diameterOfBinaryTreeHeight(node.Left)
	rightHeight := diameterOfBinaryTreeHeight(node.Right)
	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}
