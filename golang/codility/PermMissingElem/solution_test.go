package PermMissingElem

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
		{"default test", args{[]int{2, 3, 1, 5}}, 4},
		{"default test", args{[]int{2, 3, 1, 5, 6, 4}}, 7},
		{"default test", args{[]int{2, 3, 4}}, 1},
		{"default test", args{[]int{2}}, 1},
		{"default test", args{[]int{1}}, 2},
		{"default test", args{[]int{}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
