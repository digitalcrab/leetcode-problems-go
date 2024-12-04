package structures

import (
	"testing"
)

func TestAVLTree_Search(t *testing.T) {
	treeValues := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}

	tree := &AVLTree{
		Less: func(a, b any) bool {
			av := a.(int)
			bv := b.(int)
			return av < bv
		},
	}

	// prepare the tree
	for _, val := range treeValues {
		tree.Insert(val)
	}

	for _, val := range treeValues {
		if !tree.Search(val) {
			t.Errorf("Can not find expected value %d in the tree", val)
		}
	}

	notInTreeValues := []int{2, 5, 12, 5, 42}
	for _, val := range notInTreeValues {
		if tree.Search(val) {
			t.Errorf("Was able to find value %d that is not in the tree", val)
		}
	}
}
