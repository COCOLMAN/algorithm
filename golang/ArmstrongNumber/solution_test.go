package ArmstrongNumber

import "testing"

func Test_isArmstrong(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{
				n: 153,
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				n: 123,
			},
			want: false,
		},
		{
			name: "default",
			args: args{
				n: 371,
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				n: 372,
			},
			want: false,
		},
		{
			name: "default",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				n: 9474,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArmstrong(tt.args.n); got != tt.want {
				t.Errorf("isArmstrong() = %v, want %v", got, tt.want)
			}
		})
	}
}
