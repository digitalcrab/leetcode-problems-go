package structures

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	root := &intBinaryTreeNode{
		value: 4,
		left: &intBinaryTreeNode{
			value: 3,
			left: &intBinaryTreeNode{
				value: 1,
			},
		},
		right: &intBinaryTreeNode{
			value: 12,
			left: &intBinaryTreeNode{
				value: 7,
			},
		},
	}

	root.insert(8)
	root.insert(2)

	expected := []int{1, 2, 3, 4, 7, 8, 12}
	var buf []int

	root.inOrder(&buf)

	if !reflect.DeepEqual(expected, buf) {
		t.Errorf("unexpected result %v -> %v", buf, expected)
	}
}
