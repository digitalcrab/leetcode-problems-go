package main

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name      string
		args      args
		wantNums1 []int
	}{
		{
			name: "example 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0}, m: 3,
				nums2: []int{2, 5, 6}, n: 3,
			},
			wantNums1: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "example 2",
			args: args{
				nums1: []int{1}, m: 1,
				nums2: []int{}, n: 0,
			},
			wantNums1: []int{1},
		},
		{
			name: "example 3",
			args: args{
				nums1: []int{0}, m: 0,
				nums2: []int{1}, n: 1,
			},
			wantNums1: []int{1},
		},
		{
			name: "example 4",
			args: args{
				nums1: []int{2, 0}, m: 1,
				nums2: []int{1}, n: 1,
			},
			wantNums1: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.wantNums1, tt.args.nums1) {
				t.Fatalf("Expected nums1 after merge %v, but got %v\n", tt.wantNums1, tt.args.nums1)
			}
		})
	}
}
