package structures

type intBinaryTreeNode struct {
	value int
	left  *intBinaryTreeNode
	right *intBinaryTreeNode
}

func (n *intBinaryTreeNode) insert(v int) {
	if v <= n.value {
		if n.left == nil {
			n.left = &intBinaryTreeNode{value: v}
		} else {
			n.left.insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &intBinaryTreeNode{value: v}
		} else {
			n.right.insert(v)
		}
	}
}

func (n *intBinaryTreeNode) contains(v int) bool {
	if v == n.value {
		return true
	}
	if v < n.value {
		if n.left == nil {
			return false
		}
		return n.left.contains(v)
	}
	if n.right == nil {
		return false
	}
	return n.right.contains(v)
}

// Inorder (Left, Root, Right)
func (n *intBinaryTreeNode) inOrder(buf *[]int) {
	if n.left != nil {
		n.left.inOrder(buf)
	}
	*buf = append(*buf, n.value)
	if n.right != nil {
		n.right.inOrder(buf)
	}
}

// Preorder (Root, Left, Right)
func (n *intBinaryTreeNode) preOrder(buf *[]int) {
	*buf = append(*buf, n.value)
	if n.left != nil {
		n.left.inOrder(buf)
	}
	if n.right != nil {
		n.right.inOrder(buf)
	}
}

// Postorder (Left, Right, Root)
func (n *intBinaryTreeNode) portOrder(buf *[]int) {
	if n.left != nil {
		n.left.inOrder(buf)
	}
	if n.right != nil {
		n.right.inOrder(buf)
	}
	*buf = append(*buf, n.value)
}
