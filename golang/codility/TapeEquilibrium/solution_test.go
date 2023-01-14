package TapeEquilibrium

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
		{"default test", args{A: []int{3, 1, 2, 4, 3}}, 1},
		{"default test", args{A: []int{3, 1}}, 2},
		{"default test", args{A: []int{1, 3}}, 2},

		{"default test", args{A: []int{10, 20, 30, 40, 41}}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
