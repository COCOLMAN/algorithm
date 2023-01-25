package FrogJmp

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		X int
		Y int
		D int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default test", args{X: 10, Y: 85, D: 30}, 3},
		{"default test", args{X: 10, Y: 85, D: 100000000}, 1},
		{"default test", args{X: 10, Y: 10, D: 100000000}, 0},
		{"default test", args{X: 1, Y: 1, D: 1}, 0},
		{"default test", args{X: 1, Y: 2, D: 1}, 1},
		{"default test", args{X: 1, Y: 3, D: 1}, 2},
		{"default test", args{X: 1, Y: 1000000000, D: 1}, 1000000000 - 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.X, tt.args.Y, tt.args.D); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
