package MissingRanges

import (
	"reflect"
	"testing"
)

func Test_findMissingRanges(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "default",
			args: args{
				nums:  []int{0, 1, 3, 50, 75},
				lower: 0,
				upper: 99,
			},
			want: [][]int{{2, 2}, {4, 49}, {51, 74}, {76, 99}},
		},
		{
			name: "default",
			args: args{
				nums:  []int{-1},
				lower: -1,
				upper: -1,
			},
			want: [][]int{},
		},
		{
			name: "default",
			args: args{
				nums:  []int{1, 2},
				lower: 1,
				upper: 3,
			},
			want: [][]int{{3, 3}},
		},
		{
			name: "default",
			args: args{
				nums:  []int{2, 3},
				lower: 1,
				upper: 3,
			},
			want: [][]int{{1, 1}},
		},
		{
			name: "default",
			args: args{
				nums:  []int{-1000000000, 1000000000},
				lower: -1000000000,
				upper: 1000000000,
			},
			want: [][]int{{-999999999, 999999999}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingRanges(tt.args.nums, tt.args.lower, tt.args.upper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
