package SingleRowKeyboard

import "testing"

func Test_calculateTime(t *testing.T) {
	type args struct {
		keyboard string
		word     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				keyboard: "abcdefghijklmnopqrstuvwxyz",
				word:     "cba",
			},
			want: 4,
		},
		{
			name: "default",
			args: args{
				keyboard: "pqrstuvwxyzabcdefghijklmno",
				word:     "leetcode",
			},
			want: 73,
		},
		{
			name: "default",
			args: args{
				keyboard: "pqrstuvwxyzabcdefghijklmno",
				word:     "p",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTime(tt.args.keyboard, tt.args.word); got != tt.want {
				t.Errorf("calculateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
