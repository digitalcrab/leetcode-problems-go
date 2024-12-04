package structures

type AVLTree struct {
	Head *BinaryTreeNode
	Less LessFunc
}

// rotateRight performs a right rotation on the given node
// and returns the new node that should be used as a replacement
// for the given one.
// The idea is to move the left node up and the current node down.
// The left node's right child becomes the current node's left child.
// The current node becomes the right child of the left node.
func (t *AVLTree) rotateRight(node *BinaryTreeNode) *BinaryTreeNode {
	//          node
	// 	        /  \
	//       left  right
	//       /  \
	// leftLeft  leftRight
	//
	// becomes
	//          left
	// 	        /  \
	// 	  leftLeft  node
	//              /  \
	//      leftRight  right
	left := node.Left
	leftRight := left.Right

	left.Right = node
	node.Left = leftRight

	return left
}

// rotateLeft performs a left rotation on the given node
// and returns the new node that should be used as a replacement
// for the given one.
// The idea is to move the right node up and the current node down.
// The right node's left child becomes the current node's right child.
// The current node becomes the left child of the right node.
func (t *AVLTree) rotateLeft(node *BinaryTreeNode) *BinaryTreeNode {
	//          node
	// 	        /  \
	//       left  right
	//             /   \
	//     rightLeft  rightRight
	//
	// becomes
	//          right
	// 	        /  \
	//       node  rightRight
	//       /  \
	//    left  rightLeft
	right := node.Right
	rightLeft := right.Left

	right.Left = node
	node.Right = rightLeft

	return right
}

// height returns the height of the given node.
func (t *AVLTree) height(node *BinaryTreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(t.height(node.Left), t.height(node.Right))
}

// balance returns the balance factor of the given node.
func (t *AVLTree) balance(node *BinaryTreeNode) int {
	if node == nil {
		return 0
	}
	return t.height(node.Left) - t.height(node.Right)
}

func (t *AVLTree) insert(node *BinaryTreeNode, v any) *BinaryTreeNode {
	// node is still empty so we can insert the value here
	if node == nil {
		return &BinaryTreeNode{Value: v}
	}

	// if a new value is less than the current node value
	if t.Less(v, node.Value) {
		// insert the value into the left subtree
		node.Left = t.insert(node.Left, v)
	} else {
		// otherwise, insert the value into the right subtree
		node.Right = t.insert(node.Right, v)
	}

	// calculate the balance factor of the current node
	balance := t.balance(node)

	// left heavy and the new value is less than the left node value
	if balance > 1 && t.Less(v, node.Left.Value) {
		return t.rotateRight(node)
	}

	// right heavy and the new value is greater than the right node value
	if balance < -1 && !t.Less(v, node.Right.Value) {
		return t.rotateLeft(node)
	}

	// left-right case
	if balance > 1 && !t.Less(v, node.Left.Value) {
		node.Left = t.rotateLeft(node.Left)
		return t.rotateRight(node)
	}

	// right-left case
	if balance < -1 && t.Less(v, node.Right.Value) {
		node.Right = t.rotateRight(node.Right)
		return t.rotateLeft(node)
	}

	return node
}

func (t *AVLTree) Insert(v any) {
	t.Head = t.insert(t.Head, v)
}

func (t *AVLTree) Search(v any) bool {
	// start from the head
	current := t.Head

	// loop until there is something to loop over
	for current != nil {
		// calculate if v is equal to the current node value,
		if !t.Less(v, current.Value) && !t.Less(current.Value, v) {
			return true
		}

		// if not equal then move to correct side (left if current smaller)
		if t.Less(v, current.Value) {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return false
}
