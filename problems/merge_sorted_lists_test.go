package problems

import (
	"reflect"
	"testing"
)

func newNodeList(vv ...int) *node0 {
	var prev *node0
	for i := len(vv) - 1; i >= 0; i-- {
		prev = &node0{Val: vv[i], Next: prev}
	}
	return prev
}

func TestMergeTwoListsRecursion(t *testing.T) {
	cases := []struct {
		l1       *node0
		l2       *node0
		expected *node0
	}{
		{
			l1:       newNodeList(1, 2, 4),
			l2:       newNodeList(1, 3, 4),
			expected: newNodeList(1, 1, 2, 3, 4, 4),
		},
	}

	for _, c := range cases {
		if res := MergeTwoListsRecursion(c.l1, c.l2); !reflect.DeepEqual(res, c.expected) {
			t.Errorf("unexpected result of recursion (%+v) of %+v and %+v-> %+v", res, c.l1, c.l2, c.expected)
		}
	}
}

func TestMergeTwoListsIteration(t *testing.T) {
	cases := []struct {
		l1       *node0
		l2       *node0
		expected *node0
	}{
		{
			l1:       newNodeList(1, 3, 5, 6),
			l2:       newNodeList(4, 6, 9, 10),
			expected: newNodeList(1, 3, 4, 5, 6, 6, 9, 10),
		},
	}

	for _, c := range cases {
		if res := MergeTwoListsIteration(c.l1, c.l2); !reflect.DeepEqual(res, c.expected) {
			t.Errorf("unexpected result of recursion (%+v) of %+v and %+v-> %+v", res, c.l1, c.l2, c.expected)
		}
	}
}
