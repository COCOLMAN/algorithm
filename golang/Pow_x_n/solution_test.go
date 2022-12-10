package Pow_x_n

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test 1", args{2.0, 10}, 1024.00},
		{"test 1", args{2.1, 3}, 9.261000000000001},
		{"test 1", args{2.0, -2}, 0.25},
		{"test 1", args{1.0, 10000000}, 0.25},
		{"test 1", args{1.00001, 10000000}, 0.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
