package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{s: "abcabcbb"},
			want: 3, // The answer is "abc", with the length of 3.
		},
		{
			name: "example 2",
			args: args{s: "bbbbb"},
			want: 1, // The answer is "b", with the length of 1.
		},
		{
			name: "example 3",
			args: args{s: "pwwkew"},
			want: 3, // The answer is "wke", with the length of 3.
		},
		{
			name: "example 4",
			args: args{s: "aab"},
			want: 2,
		},
		{
			name: "example 5",
			args: args{s: "dvdf"},
			want: 3,
		},
		{
			name: "example 6",
			args: args{s: "abba"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
