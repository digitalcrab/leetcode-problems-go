package structures

type intTreeNode struct {
	value int
	left  *intTreeNode
	right *intTreeNode
}

func (n *intTreeNode) insert(v int) {
	if v <= n.value {
		if n.left == nil {
			n.left = &intTreeNode{value: v}
		} else {
			n.left.insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &intTreeNode{value: v}
		} else {
			n.right.insert(v)
		}
	}
}

func (n *intTreeNode) contains(v int) bool {
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

func (n *intTreeNode) inOrder(buf *[]int) {
	if n.left != nil {
		n.left.inOrder(buf)
	}
	*buf = append(*buf, n.value)
	if n.right != nil {
		n.right.inOrder(buf)
	}
}
