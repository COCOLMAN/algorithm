package FindTheIndexOfTheFirstOccurrenceInAString

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			0,
		},
		{
			"test 2",
			args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			-1,
		},
		{"test 3", args{"aaaabcabcde", "abcd"}, 6},
		{"test 4", args{"a", "a"}, 0},
		{"test 5", args{"abc", "abc"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
