package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSortNaive(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "case 1",
			nums:     []int{2, 4, 39, 7, 7, 11, 15, 9, 57, 0, 34},
			expected: []int{0, 2, 4, 7, 7, 9, 11, 15, 34, 39, 57},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := make([]int, len(tt.nums))
			copy(arg, tt.nums)
			BubbleSortNaive(arg)
			if !reflect.DeepEqual(tt.expected, arg) {
				t.Errorf("unexpected result %v", arg)
			}
		})
	}
}

func TestBubbleSortOptimized(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "case 1",
			nums:     []int{2, 4, 39, 7, 7, 11, 15, 9, 57, 0, 34},
			expected: []int{0, 2, 4, 7, 7, 9, 11, 15, 34, 39, 57},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := make([]int, len(tt.nums))
			copy(arg, tt.nums)
			BubbleSortOptimized(arg)
			if !reflect.DeepEqual(tt.expected, arg) {
				t.Errorf("unexpected result %v", arg)
			}
		})
	}
}

func BenchmarkBubbleSortOptimized(b *testing.B) {
	nums := []int{2, 4, 39, 7, 7, 11, 15, 9, 57, 0, 34}
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arg := make([]int, len(nums))
		copy(arg, nums)
		b.StartTimer()
		BubbleSortOptimized(arg)
	}
}

func BenchmarkBubbleSortNaive(b *testing.B) {
	nums := []int{2, 4, 39, 7, 7, 11, 15, 9, 57, 0, 34}
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arg := make([]int, len(nums))
		copy(arg, nums)
		b.StartTimer()
		BubbleSortNaive(arg)
	}
}
