package SearchA2DMatrix

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"default test", args{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			3,
		}, true},
		{"default test", args{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			13,
		}, false},
		{"default test", args{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			60,
		}, true},
		{"addtional test", args{[][]int{{1, 2, 3, 4}}, 5}, false},
		{"addtional test", args{[][]int{{2, 3, 4}, {5, 6, 7}}, 5}, true},
		{"addtional test", args{[][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}}, 11}, false},
		{"addtional test", args{[][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}}, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
