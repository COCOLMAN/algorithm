package LongestPalindromicSubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "default",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: "abcdcba",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "abcaadcba",
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				s: "a",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "aa",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "aaa",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "aabaa",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "aabaca",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.s); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
