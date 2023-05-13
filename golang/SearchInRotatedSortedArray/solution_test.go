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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRotatedPoint(tt.args.nums); got != tt.want {
				t.Errorf("getRotatedPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
