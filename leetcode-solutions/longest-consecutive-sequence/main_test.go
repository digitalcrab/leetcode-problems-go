package main

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{100, 4, 200, 1, 3, 2}},
			want: 4,
		},
		{
			name: "example 2",
			args: args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
			if got := longestConsecutiveForest(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutiveForest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_longestConsecutive(b *testing.B) {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	for i := 0; i < b.N; i++ {
		longestConsecutive(nums)
	}
}

func Benchmark_longestConsecutiveForest(b *testing.B) {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	for i := 0; i < b.N; i++ {
		longestConsecutiveForest(nums)
	}
}
