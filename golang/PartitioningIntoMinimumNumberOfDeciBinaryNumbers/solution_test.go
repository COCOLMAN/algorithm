package PartitioningIntoMinimumNumberOfDeciBinaryNumbers

import "testing"

func Test_minPartitions(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 32",
			args: args{
				n: "32",
			},
			want: 3,
		},
		{
			name: "test 82734",
			args: args{
				n: "82734",
			},
			want: 8,
		},
		{
			name: "test 27346209830709182346",
			args: args{
				n: "27346209830709182346",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPartitions(tt.args.n); got != tt.want {
				t.Errorf("minPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
