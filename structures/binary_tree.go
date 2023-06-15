package structures

type Comparable interface {
	Less(v Comparable) bool
	Eq(v Comparable) bool
}

type IntVal int

func (i IntVal) Less(v Comparable) bool {
	iv, ok := v.(IntVal)
	if !ok {
		return false
	}
	return int(i) < int(iv)
}

func (i IntVal) Eq(v Comparable) bool {
	iv, ok := v.(IntVal)
	if !ok {
		return false
	}
	return int(i) == int(iv)
}

type BinaryTree struct {
	Top *BinaryTreeNode
}

func NewIntBinaryTree(vals ...int) *BinaryTree {
	comp := make([]Comparable, len(vals))
	for i := range vals {
		comp[i] = IntVal(vals[i])
	}
	return NewBinaryTree(comp...)
}

func NewBinaryTree(vals ...Comparable) *BinaryTree {
	tree := &BinaryTree{}
	for _, v := range vals {
		tree.Insert(v)
	}
	return tree
}

func (n *BinaryTree) Insert(v Comparable) {
	if n.Top == nil {
		n.Top = &BinaryTreeNode{value: v}
		return
	}

	n.Top.insert(v)
}

func (n *BinaryTree) Contains(v Comparable) bool {
	if n.Top == nil {
		return false
	}
	return n.Top.contains(v)
}

type BinaryTreeNode struct {
	value Comparable
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (n *BinaryTreeNode) insert(v Comparable) {
	if v.Less(n.value) {
		if n.left == nil {
			n.left = &BinaryTreeNode{value: v}
		} else {
			n.left.insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryTreeNode{value: v}
		} else {
			n.right.insert(v)
		}
	}
}

func (n *BinaryTreeNode) contains(v Comparable) bool {
	if v.Eq(n.value) {
		return true
	}
	if v.Less(n.value) {
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
func (n *BinaryTreeNode) inOrder(buf *[]Comparable) {
	if n.left != nil {
		n.left.inOrder(buf)
	}
	*buf = append(*buf, n.value)
	if n.right != nil {
		n.right.inOrder(buf)
	}
}

// Preorder (Root, Left, Right)
func (n *BinaryTreeNode) preOrder(buf *[]Comparable) {
	*buf = append(*buf, n.value)
	if n.left != nil {
		n.left.preOrder(buf)
	}
	if n.right != nil {
		n.right.preOrder(buf)
	}
}

// Postorder (Left, Right, Root)
func (n *BinaryTreeNode) postOrder(buf *[]Comparable) {
	if n.left != nil {
		n.left.postOrder(buf)
	}
	if n.right != nil {
		n.right.postOrder(buf)
	}
	*buf = append(*buf, n.value)
}
