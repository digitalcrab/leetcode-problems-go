package structures

// BinaryTreeArray represents a Binary Tree that is stored as array (slice).
// Such a storage implies breadth-first order.
// Benefits are: way move compact in comparison to nodes and pointers.
// This tree is complete: every level, except possibly the last, is completely filled,
// and all nodes in the last level are as far left as possible.
// if a node has an index `i`, its children are found at indices `2 * i + 1` (for the left child)
// and `2 * i + 2` (for the right), while its parent (if any) is found at index
// `(i - 1) / 2` (assuming the root has index zero).
// Example:
// [ 0, 1, 2, 3, 4, 5, 6 ]
// Where:
//
//	     0
//	    / \
//	  1     2
//	 / \   / \
//	3   4 5   6
type BinaryTreeArray []any

// Insert adds element into the tree while filling the space from the most left side first.
// Time Complexity:
// O(1) on average as array is not constantly growing
// O(n) worst case if array need to grow then we copy all elements to the new one
func (bta *BinaryTreeArray) Insert(v any) {
	// as we maintain preorder (Root, Left, Right) we can simply keep adding values to the end.
	*bta = append(*bta, v)
}

// TraversePreorder visits every node in preorder Root, Left, Right and calls `cb` function.
// Time Complexity:
// O(n) white we visit every node in the tree
func (bta *BinaryTreeArray) TraversePreorder(cb func(v any)) {
	for _, value := range *bta {
		cb(value)
	}
}

// TraverseInorder visits every node in inorder Left, Root, Right and calls `cb` function.
// Time Complexity:
// O(n) white we visit every node in the tree, additional index calculations do not affect.
func (bta *BinaryTreeArray) TraverseInorder(cb func(v any)) {
	var traverse func(cb func(v any), idx int)

	// define a function to use in a recursive all
	traverse = func(cb func(v any), idx int) {
		// make sure we protect ourselves from hitting the outside boundary
		if idx >= len(*bta) {
			return
		}

		// visit the nodes in a correct order:

		traverse(cb, 2*idx+1) // Left
		cb((*bta)[idx])       // Root
		traverse(cb, 2*idx+2) // Right
	}

	traverse(cb, 0)
}

// TraversePostorder visits every node in inorder Left, Right, Root and calls `cb` function.
// Time Complexity:
// O(n) white we visit every node in the tree, additional index calculations do not affect.
func (bta *BinaryTreeArray) TraversePostorder(cb func(v any)) {
	var traverse func(cb func(v any), idx int)

	// define a function to use in a recursive all
	traverse = func(cb func(v any), idx int) {
		// make sure we protect ourselves from hitting the outside boundary
		if idx >= len(*bta) {
			return
		}

		// visit the nodes in a correct order:

		traverse(cb, 2*idx+1) // Left
		traverse(cb, 2*idx+2) // Right
		cb((*bta)[idx])       // Root
	}

	traverse(cb, 0)
}
