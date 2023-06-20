package MinimumASCIIDeleteSumForTwoStrings

import (
	"testing"
)

func Test_minimumDeleteSum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				s1: "sea",
				s2: "eat",
			},
			want: 231,
		},
		{
			name: "default",
			args: args{
				s1: "delete",
				s2: "leet",
			},
			want: 403,
		},
		{
			name: "default",
			args: args{
				s1: "asdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadfasdfasfasfasfdsadf",
				s2: "bmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxobmbmbmbmbmambmbmkaskfwmef0soxo",
			},
			want: 336516,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ascii(t *testing.T) {
	type args struct {
		a uint8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: 'a',
			},
			want: 97,
		},
		{
			name: "1",
			args: args{
				a: 'Z',
			},
			want: 90,
		},
		{
			name: "1",
			args: args{
				a: 'z',
			},
			want: 122,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ascii(tt.args.a); got != tt.want {
				t.Errorf("ascii() = %v, want %v", got, tt.want)
			}
		})
	}
}
