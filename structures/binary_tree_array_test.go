package structures

import (
	"reflect"
	"testing"
)

func TestBinaryTreeArray_TraversePreorder(t *testing.T) {
	//	     0
	//	    / \
	//	  1     2
	//	 / \   / \
	//	3   4 5   6
	bt := BinaryTreeArray{0, 1, 2, 3, 4, 5, 6}
	var result []int

	bt.TraversePreorder(func(v any) { // Root, Left, Right
		result = append(result, v.(int))
	})

	if !reflect.DeepEqual(result, []int{0, 1, 2, 3, 4, 5, 6}) {
		t.Errorf("Unexpected order of preorder traversal, got %v", result)
	}
}

func TestBinaryTreeArray_TraverseInorder(t *testing.T) {
	//	     0
	//	    / \
	//	  1     2
	//	 / \   / \
	//	3   4 5   6
	bt := BinaryTreeArray{0, 1, 2, 3, 4, 5, 6}
	var result []int

	bt.TraverseInorder(func(v any) { // Left, Root, Right
		result = append(result, v.(int))
	})

	if !reflect.DeepEqual(result, []int{3, 1, 4, 0, 5, 2, 6}) {
		t.Errorf("Unexpected order of preorder traversal, got %v", result)
	}
}

func TestBinaryTreeArray_TraversePostorder(t *testing.T) {
	//	     0
	//	    / \
	//	  1     2
	//	 / \   / \
	//	3   4 5   6
	bt := BinaryTreeArray{0, 1, 2, 3, 4, 5, 6}
	var result []int

	bt.TraversePostorder(func(v any) { // Left, Right, Root
		result = append(result, v.(int))
	})

	if !reflect.DeepEqual(result, []int{3, 4, 1, 5, 6, 2, 0}) {
		t.Errorf("Unexpected order of preorder traversal, got %v", result)
	}
}
