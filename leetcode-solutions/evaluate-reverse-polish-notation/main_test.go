package main

import "testing"

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{tokens: []string{"2", "1", "+", "3", "*"}}, // ((2 + 1) * 3) = 9
			want: 9,
		},
		{
			name: "example 2",
			args: args{tokens: []string{"4", "13", "5", "/", "+"}}, // (4 + (13 / 5)) = 6
			want: 6,
		},
		{
			name: "example 3",
			args: args{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}}, // (10 * (6 / ((9 + 3) * -11))) + 17) + 5
			want: 22,
		},
		{
			name: "example 4",
			args: args{tokens: []string{"18"}},
			want: 18,
		},
		{
			name: "example 5",
			args: args{tokens: []string{"4", "-2", "/", "2", "-3", "-", "-"}}, // (4 / -2) - (2 - (-3))
			want: -7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
