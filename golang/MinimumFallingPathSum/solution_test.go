package MinimumFallingPathSum

import "testing"

func Test_minFallingPathSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				matrix: [][]int{
					{2, 1, 3},
					{6, 5, 4},
					{7, 8, 9},
				},
			},
			want: 13,
		},
		{
			name: "default",
			args: args{
				matrix: [][]int{
					{-19, 57},
					{-40, -5},
				},
			},
			want: -59,
		},
		{
			name: "additional",
			args: args{
				matrix: [][]int{
					{-40, -5},
				},
			},
			want: -40,
		},
		{
			name: "additional",
			args: args{
				matrix: [][]int{
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
