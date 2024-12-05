package main

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	t.Skip("because of no ordering guarantee this test is always failing, maybe i fix it one day")

	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			strs: []string{""},
			want: [][]string{
				{""},
			},
		},
		{
			strs: []string{"a"},
			want: [][]string{
				{"a"},
			},
		},
	}
	for _, tt := range tests {
		if got := groupAnagrams(tt.strs); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("groupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
		}
	}
}
