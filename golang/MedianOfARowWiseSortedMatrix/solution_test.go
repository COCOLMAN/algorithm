package MedianOfARowWiseSortedMatrix

import "testing"

func Test_matrixMedian(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default test",
			args: args{
				grid: [][]int{
					{1, 1, 2},
					{2, 3, 3},
					{1, 3, 4}},
			},
			want: 2,
		},
		{
			name: "default test",
			args: args{
				grid: [][]int{
					{1, 1, 1},
					{2, 2, 2},
					{3, 3, 3}},
			},
			want: 2,
		},
		{
			name: "default test",
			args: args{
				grid: [][]int{
					{1, 1, 3, 3, 4},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixMedian(tt.args.grid); got != tt.want {
				t.Errorf("matrixMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
