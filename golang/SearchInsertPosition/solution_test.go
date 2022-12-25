package SearchInsertPosition

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default test", args{nums: []int{1, 3, 5, 6}, target: 5}, 2},
		{"default test", args{nums: []int{1, 3, 5, 6}, target: 2}, 1},
		{"default test", args{nums: []int{1, 3, 5, 6}, target: 7}, 4},
		{"default test", args{nums: []int{1, 3, 5}, target: 0}, 0},
		{"default test", args{nums: []int{1, 3, 5}, target: 1}, 0},
		{"default test", args{nums: []int{1, 3, 5}, target: 2}, 1},
		{"default test", args{nums: []int{1, 3, 5}, target: 3}, 1},
		{"default test", args{nums: []int{1, 3, 5}, target: 4}, 2},
		{"default test", args{nums: []int{1, 3, 5}, target: 5}, 2},
		{"default test", args{nums: []int{1, 3, 5}, target: 6}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
