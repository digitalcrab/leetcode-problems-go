package problems

import "testing"

func TestRemoveDuplicatesFromSortedArray80(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Input: nums = [1,1,1,2,2,3]
		// Output: 5, nums = [1,1,2,2,3,_]
		// Explanation: Your function should return k = 5, with the first five elements
		// of nums being 1, 1, 2, 2 and 3 respectively.
		// It does not matter what you leave beyond the returned k (hence they are underscores).
		{
			name: "example1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want: 5,
		},
		// Input: nums = [0,0,1,1,1,1,2,3,3]
		// Output: 7, nums = [0,0,1,1,2,3,3,_,_]
		// Explanation: Your function should return k = 7, with the first seven elements
		// of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
		// It does not matter what you leave beyond the returned k (hence they are underscores).
		{
			name: "example2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicatesFromSortedArray80(tt.args.nums); got != tt.want {
				t.Errorf("RemoveDuplicatesFromSortedArray80() = %v, want %v", got, tt.want)
			}
		})
	}
}
