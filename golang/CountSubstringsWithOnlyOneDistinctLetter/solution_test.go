package CountSubstringsWithOnlyOneDistinctLetter

import "testing"

func Test_countLetters(t *testing.T) {
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
				s: "aaaba",
			},
			want: 8,
		},
		{
			name: "default",
			args: args{
				s: "aaaaaaaaaa",
			},
			want: 55,
		},
		{
			name: "default",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "default",
			args: args{
				s: "ab",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLetters(tt.args.s); got != tt.want {
				t.Errorf("countLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
