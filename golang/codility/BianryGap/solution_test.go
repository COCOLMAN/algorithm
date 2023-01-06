package BianryGap

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "default test", args: args{N: 9}, want: 2},
		{name: "default test", args: args{N: 529}, want: 4},
		{name: "default test", args: args{N: 20}, want: 1},
		{name: "default test", args: args{N: 15}, want: 0},
		{name: "default test", args: args{N: 32}, want: 0},
		{name: "default test", args: args{N: 1041}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
