package structures

import (
	"reflect"
	"testing"
)

func TestBinaryTree_TraversePreorder(t *testing.T) {
	// Example:
	//	     1
	//	    / \
	//	  2     3
	//	 / \   / \
	//	4   5 6   7
	bt := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
	var result []int

	bt.TraversePreorder(func(v any) { // Root, Left, Right
		result = append(result, v.(int))
	})

	if !reflect.DeepEqual(result, []int{1, 2, 4, 5, 3, 6, 7}) {
		t.Errorf("Unexpected order of pre-order traversal, got %v", result)
	}
}

func TestBinaryTree_TraverseInorder(t *testing.T) {
	// Example:
	//	     5
	//	    / \
	//	  3     7
	//	 / \   / \
	//	2   4 6   8
	bt := NewBinaryTree(5, 3, 7, 2, 4, 6, 8)
	var result []int

	bt.TraverseInorder(func(v any) { // Left, Root, Right
		result = append(result, v.(int))
	})

	if !reflect.DeepEqual(result, []int{2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("Unexpected order of in-order traversal, got %v", result)
	}
}

func TestBinaryTree_TraversePostorder(t *testing.T) {
	// Example:
	//	     1
	//	    / \
	//	  2     3
	//	 / \   / \
	//	4   5 6   7
	bt := NewBinaryTree(1, 2, 3, 4, 5, 6, 7)
	var result []int

	bt.TraversePostorder(func(v any) { // Left, Right, Root
		result = append(result, v.(int))
	})

	if !reflect.DeepEqual(result, []int{4, 5, 2, 6, 7, 3, 1}) {
		t.Errorf("Unexpected order of post-order traversal, got %v", result)
	}
}
