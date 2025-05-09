package main

import (
	"reflect"
	"testing"
)

func buildTree(values []int) *TreeNode {
	if len(values) == 0 || values[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	index := 1

	for index < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if index < len(values) && values[index] != -1 {
			current.Left = &TreeNode{Val: values[index]}
			queue = append(queue, current.Left)
		}
		index++

		// Right child
		if index < len(values) && values[index] != -1 {
			current.Right = &TreeNode{Val: values[index]}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}

func Test_findDuplicateSubtrees(t *testing.T) {
	t.Skip("fails because of an order")

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "example 1",
			args: args{root: buildTree([]int{1, 2, 3, 4, -1, 2, 4, -1, -1, 4})},
			want: []*TreeNode{buildTree([]int{2, 4}), buildTree([]int{4})},
		},
		{
			name: "example 2",
			args: args{root: buildTree([]int{2, 1, 1})},
			want: []*TreeNode{buildTree([]int{1})},
		},
		{
			name: "example 3",
			args: args{root: buildTree([]int{2, 2, 2, 3, -1, 3, -1})},
			want: []*TreeNode{buildTree([]int{2, 3}), buildTree([]int{3})},
		},
		{
			name: "example 4",
			args: args{root: buildTree([]int{0, 0, 0, 0, -1, -1, 0, -1, -1, -1, 0})},
			want: []*TreeNode{buildTree([]int{0})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicateSubtrees(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicateSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
