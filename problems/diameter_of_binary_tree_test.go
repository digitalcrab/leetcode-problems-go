package problems

import "testing"

func TestDiameterOfBinaryTreeNode(t *testing.T) {
	node := func(v int, l, r *diameterOfBinaryTreeNode) *diameterOfBinaryTreeNode {
		return &diameterOfBinaryTreeNode{Val: v, Left: l, Right: r}
	}

	cases := []struct {
		root     *diameterOfBinaryTreeNode
		expected int
	}{
		{
			root:     node(1, node(2, node(4, nil, nil), node(5, nil, nil)), node(3, nil, nil)),
			expected: 3,
		},
	}

	for i, c := range cases {
		if res := diameterOfBinaryTree(c.root); c.expected != res {
			t.Errorf("[%d] unexpected result %d", i, res)
		}
	}
}
