package MedianOfTwoSortedArrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{},
			},
			want: 2.0,
		},
		{
			name: "test 2",
			args: args{
				nums1: []int{1, 2, 3, 4},
				nums2: []int{},
			},
			want: 2.5,
		},
		{
			name: "test 2",
			args: args{
				nums1: []int{},
				nums2: []int{1, 2, 3, 4},
			},
			want: 2.5,
		},
		{
			name: "test 2",
			args: args{
				nums1: []int{},
				nums2: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "test 3",
			args: args{
				nums1: []int{4, 5, 6, 7},
				nums2: []int{1, 2, 3},
			},
			want: 4,
		},
		{
			name: "test 3",
			args: args{
				nums1: []int{4, 6, 7, 8},
				nums2: []int{1, 2, 3, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValue(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		idx   int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				nums1: []int{1, 2, 3, 4},
				nums2: []int{5, 6},
				idx:   0,
			},
			want: 1,
		},
		{
			name: "test1",
			args: args{
				nums1: []int{1, 2, 3, 4},
				nums2: []int{5, 6},
				idx:   3,
			},
			want: 4,
		},
		{
			name: "test1",
			args: args{
				nums1: []int{1, 2, 3, 4},
				nums2: []int{5, 6},
				idx:   4,
			},
			want: 5,
		},
		{
			name: "test1",
			args: args{
				nums1: []int{1, 2, 3, 4},
				nums2: []int{5, 6},
				idx:   5,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValue(tt.args.nums1, tt.args.nums2, tt.args.idx); got != tt.want {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
