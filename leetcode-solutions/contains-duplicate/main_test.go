package main

import "testing"

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := containsDuplicate(tt.nums); got != tt.want {
			t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
