package CountDiv

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A int
		B int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{A: 6, B: 11, K: 2}, 3},
		{"default", args{A: 6, B: 12, K: 2}, 4},
		{"default", args{A: 5, B: 12, K: 2}, 4},
		{"default", args{A: 4, B: 12, K: 2}, 5},
		{"default", args{A: 11, B: 12, K: 2}, 1},
		{"default", args{A: 1, B: 2, K: 3}, 0},
		{"default", args{A: 1, B: 3, K: 3}, 1},
		{"default", args{A: 3, B: 3, K: 3}, 1},
		{"default", args{A: 11, B: 345, K: 17}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, tt.args.B, tt.args.K); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
