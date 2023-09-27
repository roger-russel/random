package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple phrase",
			args: args{
				s: "subinoonibus",
			},
			want: true,
		},
		{
			name: "impar palindrome",
			args: args{
				s: "radar",
			},
			want: true,
		},
		{
			name: "different cases",
			args: args{
				s: "Radar",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
