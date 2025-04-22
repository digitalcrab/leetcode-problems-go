package main

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{n: 3},
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "example 2",
			args: args{n: 5},
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "example 3",
			args: args{n: 15},
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
