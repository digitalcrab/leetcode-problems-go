package problems

// Given the root of a binary search tree, rearrange the tree in in-order so
// that the leftmost node in the tree is now the root of the tree,
// and every node has no left child and only one right child.
//
//                          1
//            5              \
//           /  \             2
//          3    6             \
//        /  \    \     =>      3
//       2    4    8             \
//      /         /  \            4
//     1         7    9            \
//                                ....
//                                  9
//
// Constraints:
// The number of nodes in the given tree will be in the range [1, 100].
// 0 <= Node.val <= 1000

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var stack []int
	increasingBST_DFS(root, &stack)
	res := &TreeNode{}
	dummy := res
	for _, v := range stack {
		dummy.Right = &TreeNode{Val: v}
		dummy = dummy.Right
	}
	return res.Right
}

func increasingBST_DFS(node *TreeNode, stack *[]int) {
	if node == nil {
		return
	}
	increasingBST_DFS(node.Left, stack)
	*stack = append(*stack, node.Val)
	increasingBST_DFS(node.Right, stack)
}
