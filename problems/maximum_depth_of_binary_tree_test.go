package problems

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *maxDepthTreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [3,9,20,null,null,15,7]",
			args: args{root: &maxDepthTreeNode{Val: 3, Left: &maxDepthTreeNode{Val: 9}, Right: &maxDepthTreeNode{Val: 20, Left: &maxDepthTreeNode{Val: 15}, Right: &maxDepthTreeNode{Val: 7}}}},
			want: 3,
		},
		{
			name: "root = [1,null,2]",
			args: args{root: &maxDepthTreeNode{Val: 1, Right: &maxDepthTreeNode{Val: 2}}},
			want: 2,
		},
		{
			name: "root = []",
			args: args{root: nil},
			want: 0,
		},
		{
			name: "root = [0]",
			args: args{root: &maxDepthTreeNode{Val: 0}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
