package MonotonicArray

import (
	"testing"
)

func Test_isMonotonic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{nums: []int{1, 2, 2, 3}}, true},
		{"test2", args{nums: []int{6, 5, 4, 4}}, true},
		{"test3", args{nums: []int{1, 3, 2}}, false},
		{"test4", args{nums: []int{1, 1, 2}}, true},
		{"test5", args{nums: []int{1}}, true},
		{"test6", args{nums: []int{1, 1}}, true},
		{"test6", args{nums: []int{1, 2}}, true},
		{"test6", args{nums: []int{2, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.nums); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
