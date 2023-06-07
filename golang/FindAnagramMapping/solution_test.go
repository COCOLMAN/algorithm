package FindAnagramMapping

import (
	"reflect"
	"testing"
)

func Test_anagramMappings(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "default",
			args: args{
				nums1: []int{12, 28, 46, 32, 50},
				nums2: []int{50, 12, 32, 46, 28},
			},
			want: []int{1, 4, 3, 2, 0},
		},
		{
			name: "default",
			args: args{
				nums1: []int{84, 46},
				nums2: []int{84, 46},
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagramMappings(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anagramMappings() = %v, want %v", got, tt.want)
			}
		})
	}
}
