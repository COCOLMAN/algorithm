package RemoveVowelsFromAString

import "testing"

func Test_removeVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default test",
			args: args{
				s: "leetcodeisacommunityforcoders",
			},
			want: "ltcdscmmntyfrcdrs",
		},
		{
			name: "default test",
			args: args{
				s: "aeiou",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeVowels(tt.args.s); got != tt.want {
				t.Errorf("removeVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
