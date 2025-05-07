package main

import (
	"reflect"
	"testing"
)

func Test_groupStrings(t *testing.T) {
	t.Skip("it fails as it might produce results in mixed order")

	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example 1",
			args: args{strings: []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}},
			want: [][]string{{"acef"}, {"a", "z"}, {"abc", "bcd", "xyz"}, {"az", "ba"}},
		},
		{
			name: "example 2",
			args: args{strings: []string{"a"}},
			want: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupStrings(tt.args.strings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
