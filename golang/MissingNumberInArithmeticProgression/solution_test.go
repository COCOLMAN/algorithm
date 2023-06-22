package MissingNumberInArithmeticProgression

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				arr: []int{
					5, 7, 11, 13,
				},
			},
			want: 9,
		},
		{
			name: "default",
			args: args{
				arr: []int{
					15, 13, 12,
				},
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.arr); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
