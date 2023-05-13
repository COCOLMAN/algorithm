package SearchInRotatedSortedArray

import "testing"

func Test_getRotatedPoint(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{nums: []int{2, 3, 4, 5, 6, 1}}, 5},
		{"test 1", args{nums: []int{3, 4, 5, 6, 1, 2}}, 4},
		{"test 1", args{nums: []int{4, 5, 6, 1, 2, 3}}, 3},
		{"test 1", args{nums: []int{5, 6, 1, 2, 3, 4}}, 2},
		{"test 1", args{nums: []int{6, 1, 2, 3, 4, 5}}, 1},
		{"test 1", args{nums: []int{1, 2, 3, 4, 5, 6}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRotatedPoint(tt.args.nums); got != tt.want {
				t.Errorf("getRotatedPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{"test1", args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, 4},
		{"test1", args{nums: []int{2, 4, 5, 6, 7, 0, 1}, target: 0}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
