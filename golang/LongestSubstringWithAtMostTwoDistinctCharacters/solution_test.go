package LongestSubstringWithAtMostTwoDistinctCharacters

import "testing"

func Test_lengthOfLongestSubstringTwoDistinct(t *testing.T) {
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
				s: "eceba",
			},
			want: 3,
		},
		{
			name: "default",
			args: args{
				s: "ccaabbb",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringTwoDistinct(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringTwoDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
