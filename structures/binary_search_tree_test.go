package structures

import (
	"testing"
)

func TestBinarySearchTree_Search(t *testing.T) {
	treeValues := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}

	bst := &BinarySearchTree{
		Less: func(a, b any) bool {
			return a.(int) < b.(int)
		},
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
	less := func(a, b any) bool {
		return a.(int) < b.(int)
	}
	bst := NewBinarySearchTree(less, 8, 3, 10, 1, 6, 14, 4, 7, 13)
	if got := bst.Min(); got != 1 {
		t.Errorf("unexpected min value: %v", got)
	}
}

func TestBinarySearchTree_Max(t *testing.T) {
	less := func(a, b any) bool {
		return a.(int) < b.(int)
	}
	bst := NewBinarySearchTree(less, 8, 3, 10, 1, 6, 14, 4, 7, 13)
	if got := bst.Max(); got != 14 {
		t.Errorf("unexpected max value: %v", got)
	}
}
