package WiggleSort

import "testing"

func check(nums []int) bool {
	i := 0
	j := 1
	for j < len(nums) {
		if i%2 == 0 {
			if nums[i] > nums[j] {
				return false
			}
		} else {
			if nums[i] < nums[j] {
				return false
			}
		}
		i++
		j++
	}
	return true
}

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default",
			args: args{
				nums: []int{3, 5, 2, 1, 6, 4},
			},
		},
		{
			name: "default",
			args: args{
				nums: []int{6, 6, 5, 6, 3, 8},
			},
		},
		{
			name: "additional",
			args: args{
				nums: []int{6, 6, 5, 1, 2, 3, 2, 1, 6, 3, 8, 1, 2, 3, 4, 3, 2, 3, 1, 3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
			if !check(tt.args.nums) {
				t.Errorf("no")
			}
		})
	}
}
