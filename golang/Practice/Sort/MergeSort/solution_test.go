package MergeSort

import "testing"

func Test_mergeSort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSort(tt.args.numbers)
		})
	}
}
