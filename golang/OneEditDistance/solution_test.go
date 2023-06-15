package OneEditDistance

import "testing"

func Test_isOneEditDistance(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{
				s: "ab",
				t: "acb",
			},
			want: true,
		},
		{
			name: "default",
			args: args{
				s: "",
				t: "",
			},
			want: false,
		},
		{
			name: "additional",
			args: args{
				s: "iacb",
				t: "acb",
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				s: "acib",
				t: "acb",
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				s: "aicb",
				t: "acb",
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				s: "acb",
				t: "aicb",
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				s: "ab",
				t: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneEditDistance(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isOneEditDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
