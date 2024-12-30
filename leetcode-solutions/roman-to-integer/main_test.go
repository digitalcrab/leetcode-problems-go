package main

import "testing"

func Test_romanToInt(t *testing.T) {
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
			args: args{s: "III"},
			want: 3, // III = 3.
		},
		{
			name: "example 2",
			args: args{s: "LVIII"}, // L = 50, V = 5, III = 3.
			want: 58,
		},
		{
			name: "example 3",
			args: args{s: "MCMXCIV"}, // M = 1000, CM = 900, XC = 90 and IV = 4.
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
