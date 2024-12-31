package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "example 2",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "example 3",
			args: args{strs: []string{"a"}},
			want: "a",
		},
		{
			name: "example 4",
			args: args{strs: []string{"flower", "flower", "flower", "flower"}},
			want: "flower",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPrefix := longestCommonPrefix(tt.args.strs); gotPrefix != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", gotPrefix, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefixSets(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "example 2",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "example 3",
			args: args{strs: []string{"a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefixSets(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefixSets() = %v, want %v", got, tt.want)
			}
		})
	}
}
