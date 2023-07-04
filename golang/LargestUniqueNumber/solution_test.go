package LargestUniqueNumber

import "testing"

func Test_largestUniqueNumber(t *testing.T) {
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
				nums: []int{5, 7, 3, 9, 4, 9, 8, 3, 1},
			},
			want: 8,
		},
		{
			name: "default",
			args: args{
				nums: []int{9, 9, 8, 8},
			},
			want: -1,
		},
		{
			name: "default",
			args: args{
				nums: []int{9, 1, 2, 3, 4},
			},
			want: 9,
		},
		{
			name: "default",
			args: args{
				nums: []int{9, 9, 9, 8, 8, 8, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestUniqueNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestUniqueNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
