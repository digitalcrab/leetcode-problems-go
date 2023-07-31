package problems

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray88(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name   string
		args   args
		result []int
	}{
		{
			name: "example1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			result: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "example2",
			args: args{
				nums1: []int{4, 0, 0, 0, 0, 0},
				m:     1,
				nums2: []int{1, 2, 3, 5, 6},
				n:     5,
			},
			result: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSortedArray88(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.result) {
				t.Errorf("MergeSortedArray88() = %v, want %v", tt.args.nums1, tt.result)
			}
		})
	}
}
