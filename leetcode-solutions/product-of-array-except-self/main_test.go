package main

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []int
	}{
		{
			name:    "example 1",
			args:    args{nums: []int{1, 2, 3, 4}},
			wantOut: []int{24, 12, 8, 6},
		},
		{
			name:    "example 2",
			args:    args{nums: []int{-1, 1, 0, -3, 3}},
			wantOut: []int{0, 0, 9, 0, 0},
		},
		{
			name:    "example 3",
			args:    args{nums: []int{0, 1, 0, -3, 3}},
			wantOut: []int{0, 0, 0, 0, 0},
		},
		{
			name:    "example 4",
			args:    args{nums: []int{1, -1}},
			wantOut: []int{-1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := productExceptSelf(tt.args.nums); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("productExceptSelf() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
