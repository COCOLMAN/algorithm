package PerformStringShifts

import "testing"

func Test_stringShift(t *testing.T) {
	type args struct {
		s     string
		shift [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				s: "abc",
				shift: [][]int{
					{0, 1},
					{1, 2},
				},
			},
			want: "cab",
		},
		{
			name: "default",
			args: args{
				s: "abcdefg",
				shift: [][]int{
					{1, 1},
					{1, 1},
					{0, 2},
					{1, 3},
				},
			},
			want: "efgabcd",
		},
		{
			name: "additional",
			args: args{
				s: "abcdefg",
				shift: [][]int{
					{1, 1},
					{1, 1},
					{0, 2},
					{0, 3},
				},
			},
			want: "defgabc",
		},
		{
			name: "additional",
			args: args{
				s: "a",
				shift: [][]int{
					{1, 100},
				},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringShift(tt.args.s, tt.args.shift); got != tt.want {
				t.Errorf("stringShift() = %v, want %v", got, tt.want)
			}
		})
	}
}
