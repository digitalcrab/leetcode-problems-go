package problems

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	st := &intStack{}
	st.push(1)
	st.push(2)
	st.push(3)

	if v := st.pop(); v != 3 {
		t.Errorf("unexpected result %d -> %d", v, 3)
	}
	if v := st.pop(); v != 2 {
		t.Errorf("unexpected result %d -> %d", v, 2)
	}
	if v := st.pop(); v != 1 {
		t.Errorf("unexpected result %d -> %d", v, 1)
	}
}

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
