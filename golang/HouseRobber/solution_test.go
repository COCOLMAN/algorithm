package HouseRobber

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "default",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
			want: 12,
		},
		{
			name: "additional",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 4,
		},
		{
			name: "additional",
			args: args{
				nums: []int{2},
			},
			want: 2,
		},
		{
			name: "additional",
			args: args{
				nums: []int{100, 1, 100, 1, 1, 100},
			},
			want: 300,
		},
		{
			name: "additional",
			args: args{
				nums: []int{100, 1, 100, 1, 1, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 100, 200, 300, 1, 2, 3, 4, 5, 100, 1, 100, 1, 1, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 100, 200, 300, 1, 2, 3, 4, 5, 100, 1, 100, 1, 1, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 100, 200, 300, 1, 2, 3, 4, 5},
			},
			want: 2291,
		},
		{
			name: "additional",
			args: args{
				nums: []int{82, 217, 170, 215, 153, 55, 185, 55, 185, 232, 69, 131, 130, 102},
			},
			want: 1082,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
