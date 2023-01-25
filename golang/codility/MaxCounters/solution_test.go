package MaxCounters

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		N int
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"default test", args{N: 5, A: []int{3, 4, 4, 6, 1, 4, 4}}, []int{3, 2, 2, 4, 2}},
		{"default test", args{N: 5, A: []int{3, 4, 4, 6, 1, 4, 4, 6}}, []int{4, 4, 4, 4, 4}},
		{"default test", args{N: 1, A: []int{1, 1, 1}}, []int{3}},
		{"default test", args{N: 1, A: []int{1}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N, tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
