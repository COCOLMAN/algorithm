package LongestPalindromicSubsequence

import "testing"

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				s: "bbbab",
			},
			want: 4,
		},
		{
			name: "default",
			args: args{
				s: "cbbd",
			},
			want: 2,
		},
		{
			name: "default",
			args: args{
				s: "cbc",
			},
			want: 3,
		},
		{
			name: "additional",
			args: args{
				s: "bbaaaaaabb",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
