package main

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{haystack: "sadbutsad", needle: "sad"},
			want: 0,
		},
		{
			name: "example 2",
			args: args{haystack: "leetcode", needle: "leeto"},
			want: -1,
		},
		{
			name: "example 3",
			args: args{haystack: "mississippi", needle: "issip"},
			want: 4,
		},
		{
			name: "example 4",
			args: args{haystack: "mississippi", needle: "issipi"},
			want: -1,
		},
		{
			name: "example 5",
			args: args{haystack: "a", needle: "a"},
			want: 0,
		},
		{
			name: "example 6",
			args: args{haystack: "abc", needle: "c"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
