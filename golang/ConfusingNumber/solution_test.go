package ConfusingNumber

import "testing"

func Test_confusingNumber(t *testing.T) {
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
				n: 6,
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				n: 89,
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				n: 11,
			},
			want: false,
		},
		{
			name: "addtional",
			args: args{
				n: 199999990,
			},
			want: true,
		},
		{
			name: "addtional",
			args: args{
				n: 199999992,
			},
			want: false,
		},
		{
			name: "addtional",
			args: args{
				n: 199999992,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := confusingNumber(tt.args.n); got != tt.want {
				t.Errorf("confusingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
