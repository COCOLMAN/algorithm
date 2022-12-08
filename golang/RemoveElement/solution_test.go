package RemoveElement

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		size int
		want []int
	}{
		{"test 1", args{nums: []int{3, 2, 2, 3}, val: 3}, 3, []int{2, 2}},
		{"test 2", args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, 5, []int{0, 1, 3, 0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.size && reflect.DeepEqual(tt.want, tt.args.nums[0:tt.size]) {
				t.Errorf("removeElement() = %v, size %v, got = %v", got, tt.size, tt.args.nums[0:tt.size])
			}
		})
	}
}
