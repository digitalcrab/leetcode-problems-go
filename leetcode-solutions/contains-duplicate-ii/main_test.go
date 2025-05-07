package main

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{nums: []int{1, 2, 3, 1}, k: 3},
			want: true,
		},
		{
			name: "example 2",
			args: args{nums: []int{1, 0, 1, 1}, k: 1},
			want: true,
		},
		{
			name: "example 3",
			args: args{nums: []int{1, 2, 3, 1, 2, 3}, k: 2},
			want: false,
		},
		{
			name: "example 4",
			args: args{nums: []int{1, 2, 1}, k: 0},
			want: false,
		},
		{
			name: "example 5",
			args: args{nums: []int{1, 4, 2, 3, 1, 2}, k: 3},
			want: true,
		},
		{
			name: "example 6",
			args: args{nums: []int{1}, k: 1},
			want: false,
		},
		{
			name: "example 7",
			args: args{nums: []int{99, 99}, k: 2},
			want: true,
		},
		{
			name: "example 8",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, k: 15},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
