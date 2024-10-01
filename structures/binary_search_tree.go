package structures

// LessFunc reports whether the element a must sort before the element b.
// If both Less(a, b) and Less(b, q) are false, then the elements  and b considered equal.
type LessFunc func(a, b any) bool

// BinarySearchTree is a binary tree where for each node:
// - All values in the left subtree are less than the node’s value.
type BinarySearchTree struct {
	Head *BinaryTreeNode
	Less LessFunc
}

func NewBinarySearchTree(less LessFunc, values ...any) *BinarySearchTree {
	bst := &BinarySearchTree{Less: less}
	for _, val := range values {
		bst.Insert(val)
	}
	return bst
}

// Insert performs an insert into the binary search tree keeping the left node smaller than root and right bigger
// Time Complexity:
// O(n) works-case if we end up looping thought all elements, and they all fall into the same left or right node
// O(log(n)) on average if tree more of less balanced, and we always keep 2 times fewer elements to visit
func (bst *BinarySearchTree) Insert(v any) {
	var destination *BinaryTreeNode

	// go through the tree and find the destination for the new value
	// maintaining the characteristics of a tree
	current := bst.Head
	for current != nil {
		// set destination to the current node
		destination = current

		// if the new value is less than current value move the pointer of a current to the left node
		if bst.Less(v, current.Value) {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	newNode := &BinaryTreeNode{Value: v}

	// at this stage our distinction should point to the node where the new value supposed to be located

	// edge-case for the empty tree\
	if destination == nil {
		bst.Head = newNode
		return
	}

	// once again calculate where to put the final node
	if bst.Less(v, destination.Value) {
		destination.Left = newNode
	} else {
		destination.Right = newNode
	}
}

// Search performs search in the binary tree
// Time Complexity:
// O(n) works-case if we end up looping thought all elements, and they all fall into the same left or right node
// O(log(n)) on average if tree more of less balanced, and we always keep 2 times fewer elements to visit
func (bst *BinarySearchTree) Search(v any) bool {
	// start from the head
	current := bst.Head

	// loop until there is something to loop over
	for current != nil {
		// calculate if v is equal to the current node value,
		if !bst.Less(v, current.Value) && !bst.Less(current.Value, v) {
			return true
		}

		// if not equal then move to correct side (left if current smaller)
		if bst.Less(v, current.Value) {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return false
}

// Min returns the smallest element in the tree
// Time Complexity: O(log(n)) on average or O(n) worse-case
func (bst *BinarySearchTree) Min() any {
	for current := bst.Head; current != nil; current = current.Left {
		if current.Left == nil {
			return current.Value
		}
	}
	return nil
}

// Max returns the biggest element in the tree
// Time Complexity: O(log(n)) on average or O(n) worse-case
func (bst *BinarySearchTree) Max() any {
	for current := bst.Head; current != nil; current = current.Right {
		if current.Right == nil {
			return current.Value
		}
	}
	return nil
}
