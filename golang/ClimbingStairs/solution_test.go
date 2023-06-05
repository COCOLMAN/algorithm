package ClimbingStairs

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "default",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "default",
			args: args{
				n: 4,
				/*
					1 1 1 1
					2 1 1
					1 2 1
					1 1 2
					2 2
				*/
			},
			want: 5,
		},
		{
			name: "default",
			args: args{
				n: 5,
				/*
					1 1 1 1 1
					2 1 1 1
					1 2 1 1
					1 1 2 1
					1 1 1 2
					2 2 1
					2 1 2
					1 2 2
				*/
			},
			want: 8,
		},
		{
			name: "default",
			args: args{
				n: 30,
			},
			want: 1346269,
		},
		{
			name: "default",
			args: args{
				n: 45,
			},
			want: 1836311903,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
