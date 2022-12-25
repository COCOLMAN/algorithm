package GuessNumberHigherOrLower

import (
	"reflect"
	"testing"
)

func Test_generateGuess(t *testing.T) {
	type args struct {
		n    int
		pick int
	}
	tests := []struct {
		name string
		args args
	}{
		{"default test", args{n: 10, pick: 6}},
		{"default test", args{n: 1, pick: 1}},
		{"default test", args{n: 2, pick: 1}},
		{"default test", args{n: 1000, pick: 3}},
		{"default test", args{n: 3, pick: 2}},
		{"default test", args{n: 3, pick: 3}},
		{"default test", args{n: 3, pick: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := generateGuess(tt.args.pick)
			got := f(tt.args.n)
			if !reflect.DeepEqual(got, tt.args.pick) {
				t.Errorf("generateGuess() = %v, want %v", got, tt.args.pick)
			}
		})
	}
}
