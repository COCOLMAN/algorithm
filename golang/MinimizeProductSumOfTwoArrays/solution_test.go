package MinimizeProductSumOfTwoArrays

import "testing"

func Test_minProductSum(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				nums1: []int{5, 3, 4, 2},
				nums2: []int{4, 2, 2, 5},
			},
			want: 40,
		},
		{
			name: "default",
			args: args{
				nums1: []int{2, 1, 4, 5, 7},
				nums2: []int{3, 2, 4, 8, 6},
			},
			want: 65,
		},
		{
			name: "max num",
			args: args{
				nums1: []int{1, 2, 100},
				nums2: []int{1, 2, 100},
			},
			want: 204,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minProductSum(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minProductSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
