package main

import "testing"

func Test_checkPalindrome(t *testing.T) {
	type args struct {
		r []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple phrase",
			args: args{
				r: []rune("subinoonibus"),
			},
			want: true,
		},
		{
			name: "odd palindrome",
			args: args{
				r: []rune("radar"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.r); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
