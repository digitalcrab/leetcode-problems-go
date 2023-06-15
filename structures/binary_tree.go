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

// DFSInOrder (Left, Root, Right) Depth-First-Search
func (n *BinaryTreeNode) DFSInOrder(buf *[]Comparable) {
	if n.left != nil {
		n.left.DFSInOrder(buf)
	}
	*buf = append(*buf, n.value)
	if n.right != nil {
		n.right.DFSInOrder(buf)
	}
}

// DFSPreOrder (Root, Left, Right) Depth-First-Search
func (n *BinaryTreeNode) DFSPreOrder(buf *[]Comparable) {
	*buf = append(*buf, n.value)
	if n.left != nil {
		n.left.DFSPreOrder(buf)
	}
	if n.right != nil {
		n.right.DFSPreOrder(buf)
	}
}

// DFSPostOrder (Left, Right, Root) Depth-First-Search
func (n *BinaryTreeNode) DFSPostOrder(buf *[]Comparable) {
	if n.left != nil {
		n.left.DFSPostOrder(buf)
	}
	if n.right != nil {
		n.right.DFSPostOrder(buf)
	}
	*buf = append(*buf, n.value)
}

// BFS (All siblings) Breadth-First-Search
func (n *BinaryTreeNode) BFS(buf *[]Comparable) {
	nextToVisit := []*BinaryTreeNode{n}

	for len(nextToVisit) > 0 {
		// take one from queue and dec queue
		var current *BinaryTreeNode
		current, nextToVisit = nextToVisit[0], nextToVisit[1:]

		*buf = append(*buf, current.value)

		if current.left != nil {
			nextToVisit = append(nextToVisit, current.left)
		}
		if current.right != nil {
			nextToVisit = append(nextToVisit, current.right)
		}
	}
}
