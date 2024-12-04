package main

import "testing"

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			s:    "rat",
			t:    "car",
			want: false,
		},
	}
	for _, tt := range tests {
		if got := isAnagram(tt.s, tt.t); got != tt.want {
			t.Errorf("isAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
