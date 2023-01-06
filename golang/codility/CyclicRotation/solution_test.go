package CyclicRotation

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"default test", args{[]int{3, 8, 9, 7, 6}, 1}, []int{6, 3, 8, 9, 7}},
		{"default test", args{[]int{3, 8, 9, 7, 6}, 2}, []int{7, 6, 3, 8, 9}},
		{"default test", args{[]int{3, 8, 9, 7, 6}, 3}, []int{9, 7, 6, 3, 8}},
		{"default test", args{[]int{0, 0, 0}, 1}, []int{0, 0, 0}},
		{"default test", args{[]int{1, 2, 3, 4}, 4}, []int{1, 2, 3, 4}},
		{"default test", args{[]int{1, 2, 3, 4}, 0}, []int{1, 2, 3, 4}},
		{"default test", args{[]int{1, 2, 3, 4}, 1}, []int{4, 1, 2, 3}},
		{"default test", args{[]int{1, 2, 3, 4}, 2}, []int{3, 4, 1, 2}},
		{"default test", args{[]int{1, 2, 3, 4}, 3}, []int{2, 3, 4, 1}},
		{"default test", args{[]int{}, 3}, []int{}},
		{"default test", args{[]int{}, 2}, []int{}},
		{"default test", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1}, []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"default test", args{[]int{1}, 3}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
