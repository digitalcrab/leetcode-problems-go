package problems

import (
	"reflect"
	"testing"
)

func TestRotateArray189Reverse(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name   string
		args   args
		result []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			result: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			result: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateArray189Reverse(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.result) {
				t.Errorf("RotateArray189() = %v, want %v", tt.args.nums, tt.result)
			}
		})
	}
}

func TestRotateArray189ShiftingToRight(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name   string
		args   args
		result []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			result: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			result: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateArray189ShiftingToRight(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.result) {
				t.Errorf("RotateArray189() = %v, want %v", tt.args.nums, tt.result)
			}
		})
	}
}

func TestRotateArray189ExtraMemory(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name   string
		args   args
		result []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			result: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			result: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateArray189ExtraMemory(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.result) {
				t.Errorf("RotateArray189() = %v, want %v", tt.args.nums, tt.result)
			}
		})
	}
}
