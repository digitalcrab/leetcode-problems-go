package problems

import "testing"

func TestIsSubtree(t *testing.T) {
	node := func(v int, l, r *isSubtreeNode) *isSubtreeNode {
		return &isSubtreeNode{Val: v, Left: l, Right: r}
	}

	cases := []struct {
		s        *isSubtreeNode
		t        *isSubtreeNode
		expected bool
	}{
		{
			s:        node(3, node(4, node(1, nil, nil), node(2, node(0, nil, nil), nil)), node(5, nil, nil)),
			t:        node(4, node(1, nil, nil), node(2, nil, nil)),
			expected: false,
		},
		{
			s:        node(3, node(4, node(1, nil, nil), node(2, nil, nil)), node(5, nil, nil)),
			t:        node(4, node(1, nil, nil), node(2, nil, nil)),
			expected: true,
		},
	}

	for i, c := range cases {
		if res := isSubtree(c.s, c.t); c.expected != res {
			t.Errorf("[%d] unexpected result %t", i, res)
		}
	}
}
