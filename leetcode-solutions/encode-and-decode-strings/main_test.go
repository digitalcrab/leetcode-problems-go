package main

import (
	"reflect"
	"testing"
)

func TestCodec(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{strs: []string{"Hello", "World"}},
			want: []string{"Hello", "World"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var codec Codec
			if got := codec.Decode(codec.Encode(tt.args.strs)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("codec.Decode(codec.Encode()) = %v, want %v", got, tt.want)
			}
		})
	}
}
