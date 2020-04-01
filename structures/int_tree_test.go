package structures

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	root := &intTreeNode{
		value: 4,
		left: &intTreeNode{
			value: 3,
			left: &intTreeNode{
				value: 1,
			},
		},
		right: &intTreeNode{
			value: 12,
			left: &intTreeNode{
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
