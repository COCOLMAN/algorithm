package PassingCars

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
		{"default test", args{A: []int{0, 1, 0, 1, 1}}, 5},
		{"addtional test", args{A: []int{0, 1}}, 1},
		{"addtional test", args{A: []int{1, 0}}, 0},
		{"addtional test", args{A: []int{0}}, 0},
		{"addtional test", args{A: []int{1}}, 0},
		{"addtional test", args{A: []int{1, 1, 1, 1, 1}}, 0},
		{"addtional test", args{A: []int{0, 0, 0}}, 0},
		{"addtional test", args{A: []int{1, 0, 1, 0, 0}}, 1},
		{"addtional test", args{A: []int{0, 1, 0, 1, 1, 1}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
