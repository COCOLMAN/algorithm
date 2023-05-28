package DeleteAndEarn

import "testing"

func Test_deleteAndEarn(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default 1",
			args: args{
				nums: []int{3, 4, 2},
			},
			want: 6,
		},
		{
			name: "default 2",
			args: args{
				nums: []int{2, 2, 3, 3, 3, 4},
			},
			want: 9,
		},
		{
			name: "additional 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: 16 + 12 + 8 + 4,
		},
		{
			name: "additional 2",
			args: args{
				nums: []int{8, 6, 1, 7, 5, 8, 9, 5, 1, 1, 7, 3, 5, 8, 5, 2, 9, 6, 9, 10, 10, 10, 4, 4, 8, 8, 4, 3, 6, 7, 4, 5, 1, 7, 1, 5, 1, 6, 7, 9, 6, 4, 8, 7, 9, 7, 8, 2, 9, 5},
			},
			want: 150,
		},
		{
			name: "additional 3",
			args: args{
				nums: []int{12, 32, 93, 17, 100, 72, 40, 71, 37, 92, 58, 34, 29, 78, 11, 84, 77, 90, 92, 35, 12, 5, 27, 92, 91, 23, 65, 91, 85, 14, 42, 28, 80, 85, 38, 71, 62, 82, 66, 3, 33, 33, 55, 60, 48, 78, 63, 11, 20, 51, 78, 42, 37, 21, 100, 13, 60, 57, 91, 53, 49, 15, 45, 19, 51, 2, 96, 22, 32, 2, 46, 62, 58, 11, 29, 6, 74, 38, 70, 97, 4, 22, 76, 19, 1, 90, 63, 55, 64, 44, 90, 51, 36, 16, 65, 95, 64, 59, 53, 93},
			},
			want: 3451,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteAndEarn(tt.args.nums); got != tt.want {
				t.Errorf("deleteAndEarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
