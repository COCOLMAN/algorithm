package WordBreak

import "testing"

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
		{
			name: "default",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat", "an"},
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				s:        "abcde",
				wordDict: []string{"a"},
			},
			want: false,
		},
		{
			name: "additional",
			args: args{
				s:        "abcde",
				wordDict: []string{"a", "b", "c", "d", "e"},
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				s:        "a",
				wordDict: []string{"a", "b", "c", "d", "e"},
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
				wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
