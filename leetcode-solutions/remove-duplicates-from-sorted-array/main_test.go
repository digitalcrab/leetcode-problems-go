package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name:     "example 1",
			args:     args{nums: []int{1, 1, 2}},
			want:     2,
			wantNums: []int{1, 2},
		},
		{
			name:     "example 2",
			args:     args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want:     5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.nums[:got], tt.wantNums) {
				t.Errorf("removeDuplicates() = %v, want nums %v", tt.args.nums[:got], tt.wantNums)
			}
		})
	}
}
