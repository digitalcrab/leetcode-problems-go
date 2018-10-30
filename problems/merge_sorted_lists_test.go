package problems

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	in1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	in2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	res := MergeTwoLists(in1, in2)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("unexpected result %v", res)
	}
}
