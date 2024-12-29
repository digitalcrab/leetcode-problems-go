package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		args     args
		wantNums []int
	}{
		{
			name:     "example 1",
			args:     args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3},
			wantNums: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "example 2",
			args:     args{nums: []int{-1, -100, 3, 99}, k: 2},
			wantNums: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.wantNums, tt.args.nums) {
				t.Fatalf("Expected nums after rotate %v, but got %v\n", tt.wantNums, tt.args.nums)
			}
		})
	}
}

func Test_rotateV2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		args     args
		wantNums []int
	}{
		{
			name:     "example 1",
			args:     args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3},
			wantNums: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "example 2",
			args:     args{nums: []int{-1, -100, 3, 99}, k: 2},
			wantNums: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateV2(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.wantNums, tt.args.nums) {
				t.Fatalf("Expected nums after rotateV2 %v, but got %v\n", tt.wantNums, tt.args.nums)
			}
		})
	}
}
