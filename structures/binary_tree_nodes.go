package structures

type BinaryTree struct {
	Head *BinaryTreeNode
}

type BinaryTreeNode struct {
	Value any
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// NewBinaryTree creates a new tree with values
func NewBinaryTree(values ...any) *BinaryTree {
	bt := &BinaryTree{}
	for _, val := range values {
		bt.Insert(val)
	}
	return bt
}

// Insert performs level-order insertion, filling the tree from left to right at each level
// Time Complexity:
// O(n) as we might visit all the items in the tree
// O(n) on average as well as this is not a tree where items organised in sorted order
// Space Complexity:
// O(n) as we use the queue to keep elements to visit, at the very last stage we can have all the leafs
// of a last nodes which is n/2 -> so basically n
func (bt *BinaryTree) Insert(v any) {
	node := &BinaryTreeNode{Value: v}

	// special case of an empty tree
	if bt.Head == nil {
		bt.Head = node
		return
	}

	// create a small queue with head added upfront
	queue := []*BinaryTreeNode{bt.Head}

	// keep working while queue is not empty
	for len(queue) > 0 {
		// get the first item from the queue
		current := queue[0]
		// and chop it from the queue beginning
		queue = queue[1:]

		if current.Left == nil {
			// in case if we see an empty spot on the left we add it there and
			// exit the loop immediately as the job done
			current.Left = node
			return
		} else {
			// if it's not empty, add that to the queue to check it's left and right nodes
			queue = append(queue, current.Left)
		}

		// same goes for the right node
		if current.Right == nil {
			current.Right = node
			return
		} else {
			queue = append(queue, current.Right)
		}
	}
}

// TraversePreorder visits every node in order Root, Left, Right and calls `cb` function.
// Time Complexity:
// O(n)
func (bt *BinaryTree) TraversePreorder(cb func(v any)) {
	bt.traversePreorder(cb, bt.Head)
}

func (bt *BinaryTree) traversePreorder(cb func(v any), node *BinaryTreeNode) {
	if node == nil {
		return
	}
	cb(node.Value)
	bt.traversePreorder(cb, node.Left)
	bt.traversePreorder(cb, node.Right)
}

// TraverseInorder visits every node in inorder Left, Root, Right and calls `cb` function.
// Time Complexity:
// O(n)
func (bt *BinaryTree) TraverseInorder(cb func(v any)) {
	bt.traverseInorder(cb, bt.Head)
}

func (bt *BinaryTree) traverseInorder(cb func(v any), node *BinaryTreeNode) {
	if node == nil {
		return
	}
	bt.traverseInorder(cb, node.Left)
	cb(node.Value)
	bt.traverseInorder(cb, node.Right)
}

// TraversePostorder visits every node in order Left, Right, Root and calls `cb` function.
// Time Complexity:
// O(n)
func (bt *BinaryTree) TraversePostorder(cb func(v any)) {
	bt.traversePostorder(cb, bt.Head)
}

func (bt *BinaryTree) traversePostorder(cb func(v any), node *BinaryTreeNode) {
	if node == nil {
		return
	}
	bt.traversePostorder(cb, node.Left)
	bt.traversePostorder(cb, node.Right)
	cb(node.Value)
}
