package FibonacciNumber

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{n: 0}, 0},
		{"test 1", args{n: 1}, 1},
		{"test 1", args{n: 2}, 1},
		{"test 1", args{n: 3}, 2},
		{"test 1", args{n: 4}, 3},
		{"test 1", args{n: 30}, 3},
		{"test 1", args{n: 300}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
