package UniquePaths

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default test",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "default test",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "min size test",
			args: args{
				m: 1,
				n: 1,
			},
			want: 1,
		},
		{
			name: "small size test",
			args: args{
				m: 2,
				n: 1,
			},
			want: 1,
		},
		{
			name: "4x4 size test",
			args: args{
				m: 4,
				n: 4,
			},
			want: 20,
		},
		{
			name: "max size test",
			args: args{
				m: 100,
				n: 100,
			},
			want: 4631081169483718960,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
