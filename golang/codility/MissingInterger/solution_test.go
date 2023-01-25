package MissingInterger

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default test", args{A: []int{1, 3, 6, 4, 1, 2}}, 5},
		{"default test", args{A: []int{1, 3, 6, 5, 4, 1, 2}}, 7},
		{"default test", args{A: []int{1, 2, 3}}, 4},
		{"default test", args{A: []int{-1, -3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
