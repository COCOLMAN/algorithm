package CheckIfANumberIsMajorityElementInASortedArray

import "testing"

func Test_isMajorityElement(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{
				nums:   []int{2, 4, 5, 5, 5, 5, 5, 6, 6},
				target: 5,
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				nums:   []int{10, 100, 101, 101},
				target: 101,
			},
			want: false,
		},
		{
			name: "additional",
			args: args{
				nums:   []int{1, 2, 3, 4},
				target: 1,
			},
			want: false,
		},
		{
			name: "additional",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				nums:   []int{1, 2},
				target: 1,
			},
			want: false,
		},
		{
			name: "additional",
			args: args{
				nums:   []int{1, 2, 2},
				target: 2,
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				nums:   []int{1},
				target: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMajorityElement(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("isMajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
