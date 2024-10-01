package structures

import (
	"testing"
)

func bstLessInt(a, b any) bool {
	av := a.(int)
	bv := b.(int)
	return av < bv
}

func TestBinarySearchTree_Search(t *testing.T) {
	eq := !bstLessInt(5, 5) && !bstLessInt(5, 5)
	if eq != true {
		t.Errorf("Less function works incorrect")
	}

	treeValues := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}

	bst := &BinarySearchTree{
		Less: bstLessInt,
	}

	// prepare the tree
	for _, val := range treeValues {
		bst.Insert(val)
	}

	for _, val := range treeValues {
		if !bst.Search(val) {
			t.Errorf("Can not find expected value %d in the tree", val)
		}
	}

	notInTreeValues := []int{2, 5, 12, 5, 42}
	for _, val := range notInTreeValues {
		if bst.Search(val) {
			t.Errorf("Was able to find value %d that is not in the tree", val)
		}
	}
}

func TestBinarySearchTree_Min(t *testing.T) {
	bst := NewBinarySearchTree(bstLessInt, 8, 3, 10, 1, 6, 14, 4, 7, 13)
	if got := bst.Min(); got != 1 {
		t.Errorf("unexpected min value: %v", got)
	}
}

func TestBinarySearchTree_Max(t *testing.T) {
	bst := NewBinarySearchTree(bstLessInt, 8, 3, 10, 1, 6, 14, 4, 7, 13)
	if got := bst.Max(); got != 14 {
		t.Errorf("unexpected max value: %v", got)
	}
}
