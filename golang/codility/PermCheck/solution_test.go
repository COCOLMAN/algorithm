package PermCheck

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
		{"default test", args{[]int{4, 1, 3, 2}}, 1},
		{"default test", args{[]int{4, 1, 3, 2, 2}}, 0},
		{"default test", args{[]int{4, 1, 3}}, 0},
		{"default test", args{[]int{1}}, 1},
		{"default test", args{[]int{2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
