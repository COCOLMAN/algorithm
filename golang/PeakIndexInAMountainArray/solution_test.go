package PeakIndexInAMountainArray

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default test", args{arr: []int{0, 1, 0}}, 1},
		{"default test", args{arr: []int{0, 2, 1, 0}}, 1},
		{"default test", args{arr: []int{0, 10, 5, 2}}, 1},
		{"default test", args{arr: []int{0, 10, 50, 2}}, 2},
		{"default test", args{arr: []int{0, 10, 2}}, 1},
		{"default test", args{arr: []int{1, 2, 3, 4, 1}}, 3},
		{"default test", args{arr: []int{1, 10, 5, 4, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
