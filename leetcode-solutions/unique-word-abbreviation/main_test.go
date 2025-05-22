package main

import (
	"testing"
)

func TestValidWordAbbr(t *testing.T) {
	type args struct {
		dictionary []string
		words      []string
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "example 1",
			args: args{
				dictionary: []string{"deer", "door", "cake", "card"},
				words:      []string{"dear", "cart", "cane", "make", "cake"},
			},
			want: []bool{false, true, false, true, true},
		},
		{
			name: "example 2",
			args: args{
				dictionary: []string{"dog"},
				words:      []string{"dig", "dug", "dag", "dog", "doge"},
			},
			want: []bool{false, false, false, true, true},
		},
		{
			name: "example 3",
			args: args{
				dictionary: []string{"deer", "door", "cake", "card"},
				words:      []string{"dear", "door", "cart", "cake"},
			},
			want: []bool{false, false, true, true},
		},
		{
			name: "example 4",
			args: args{
				dictionary: []string{"a", "a"},
				words:      []string{"a"},
			},
			want: []bool{true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.args.dictionary)
			for i, w := range tt.want {
				if got := this.IsUnique(tt.args.words[i]); got != w {
					t.Errorf("IsUnique(%q) = %t, want %t", tt.args.words[i], got, w)
				}
			}
		})
	}
}
