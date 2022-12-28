package SumOfSquareNumbers

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"default test", args{5}, true},
		{"default test", args{3}, false},
		{"default test", args{10}, true},
		{"default test", args{10}, true},
		{"default test", args{3047981}, true},
		{"default test", args{1538170825}, true},
		{"default test", args{648770825}, true},
		{"default test", args{123}, false},
		{"default test", args{500}, true},
		{"default test", args{4}, true},
		{"default test", args{2}, true},
		{"default test", args{1}, true},
		{"default test", args{999999999}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
