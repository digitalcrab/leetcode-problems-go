package problems

import "testing"

func TestMajorityElement169(t *testing.T) {
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
			args: args{nums: []int{3, 2, 3}},
			want: 3,
		},
		{
			name: "example 2",
			args: args{nums: []int{2, 2, 1, 1, 1, 2, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement169(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement169() = %v, want %v", got, tt.want)
			}
		})
	}
}
