package problems

import "testing"

func TestRemoveDuplicatesFromSortedArray26(t *testing.T) {
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
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicatesFromSortedArray26(tt.args.nums); got != tt.want {
				t.Errorf("RemoveDuplicatesFromSortedArray26() = %v, want %v", got, tt.want)
			}
		})
	}
}
