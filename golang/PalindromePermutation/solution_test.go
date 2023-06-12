package PalindromePermutation

import (
	"testing"
)

func Test_canPermutePalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{
				s: "code",
			},
			want: false,
		},
		{
			name: "default",
			args: args{
				s: "aab",
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				s: "carerac",
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				s: "careracc",
			},
			want: false,
		},
		{
			name: "long",
			args: args{
				s: "abcfadsfwefasdfasdfasfdsafasdfsafdsafdsafewfsdafcba",
			},
			want: true,
		},
		{
			name: "short",
			args: args{
				s: "a",
			},
			want: true,
		},
		{
			name: "short",
			args: args{
				s: "aa",
			},
			want: true,
		},
		{
			name: "short",
			args: args{
				s: "ab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPermutePalindrome(tt.args.s); got != tt.want {
				t.Errorf("canPermutePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
