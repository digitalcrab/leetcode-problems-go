package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}

	// here we are going to store the number of times we've seen
	// sub-tree, by calculating the hash of the tree
	seen := make(map[string][]*TreeNode)
	dfs(root, seen)

	var result []*TreeNode
	for _, nodes := range seen {
		if len(nodes) > 1 {
			result = append(result, nodes[0])
		}
	}

	return result
}

func dfs(node *TreeNode, seen map[string][]*TreeNode) string {
	// node is empty, lets use nil hash
	if node == nil {
		return "nil"
	}

	// calc left side
	leftHash := dfs(node.Left, seen)
	// same as right
	rightHash := dfs(node.Right, seen)

	// build a combined hash
	hash := fmt.Sprintf("(%s:%d:%s)", leftHash, node.Val, rightHash)
	// store the node and a hash
	seen[hash] = append(seen[hash], node)
	// let it go
	return hash
}

func main() {

}
