package main

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name:     "example 1",
			args:     args{nums: []int{3, 2, 2, 3}, val: 3},
			want:     2,
			wantNums: []int{2, 2},
		},
		{
			name:     "example 2",
			args:     args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			want:     5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.nums[:got], tt.wantNums) {
				t.Errorf("removeElement() = %v, want nums %v", tt.args.nums[:got], tt.wantNums)
			}
		})
	}
}

func Test_removeElementStack(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name:     "example 1",
			args:     args{nums: []int{3, 2, 2, 3}, val: 3},
			want:     2,
			wantNums: []int{2, 2},
		},
		{
			name:     "example 2",
			args:     args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			want:     5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElementStack(tt.args.nums, tt.args.val)
			if got != tt.want {
				t.Errorf("removeElementStack() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.nums[:got], tt.wantNums) {
				t.Errorf("removeElement() = %v, want nums %v", tt.args.nums[:got], tt.wantNums)
			}
		})
	}
}
