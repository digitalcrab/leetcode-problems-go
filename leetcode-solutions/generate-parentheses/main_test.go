package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []string
	}{
		{
			name:    "example 1",
			args:    args{n: 3},
			wantOut: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:    "example 2",
			args:    args{n: 1},
			wantOut: []string{"()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut := generateParenthesis(tt.args.n)

			// it's necessary we we do not really guarantee the order
			sort.Strings(gotOut)
			sort.Strings(tt.wantOut)

			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("generateParenthesis() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
