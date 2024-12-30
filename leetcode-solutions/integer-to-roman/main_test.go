package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{num: 3749},
			want: "MMMDCCXLIX",
		},
		{
			name: "example 2",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "example 3",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman(%d) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}
