package BubbleSort

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				numbers: []int{4, 2, 1, 3},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Test 2",
			args: args{
				numbers: []int{3, 3, 3, 2, 1},
			},
			want: []int{1, 2, 3, 3, 3},
		},
		{
			name: "Test 3",
			args: args{
				numbers: []int{1, 3, 1, 1, 1, 1, 1},
			},
			want: []int{1, 1, 1, 1, 1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.numbers)
			if !reflect.DeepEqual(tt.args.numbers, tt.want) {
				t.Error("error")
			}
		})
	}
}
