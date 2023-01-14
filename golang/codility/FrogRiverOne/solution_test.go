package FrogRiverOne

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		X int
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default test", args{X: 5, A: []int{1, 3, 1, 4, 2, 3, 5, 4}}, 6},
		{"additional test", args{X: 5, A: []int{1, 3, 1, 4, 3, 5, 4}}, -1},
		{"additional test", args{X: 5, A: []int{1}}, -1},
		{"additional test", args{X: 1, A: []int{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.X, tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
