package BinarySearch

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, 4},
		{"test 2", args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, 4},
		{"test 2", args{nums: []int{-1, 0, 3, 5, 9}, target: 3}, 2},
		{"test 2", args{nums: []int{-1, 0, 3, 5, 9}, target: 0}, 1},
		{"test 2", args{nums: []int{-1}, target: 0}, -1},
		{"test 2", args{nums: []int{-1}, target: -1}, 0},
		{"test 2", args{nums: []int{1, 2, 3}, target: 4}, -1},
		{"test 2", args{nums: []int{1, 2, 3}, target: -1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
