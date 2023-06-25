package MeetingRooms2

import (
	"reflect"
	"testing"
)

func Test_minMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				intervals: [][]int{
					{0, 30},
					{5, 10},
					{15, 20},
				},
			},
			want: 2,
		},
		{
			name: "default",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
				},
			},
			want: 1,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{5, 6},
				},
			},
			want: 1,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{1, 3},
					{2, 4},
					{2, 6},
					{5, 6},
				},
			},
			want: 3,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{1, 5},
					{2, 5},
					{3, 5},
					{5, 6},
					{6, 7},
					{6, 8},
				},
			},
			want: 3,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{5, 8},
					{6, 8},
				},
			},
			want: 2,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{0, 30},
					{5, 10},
					{15, 20},
				},
			},
			want: 2,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{119, 207},
					{333, 645},
					{384, 846},
					{7, 198},
					{345, 660},
					{74, 596},
					{139, 490},
					{783, 983},
				},
			},
			want: 5,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{8, 17},
					{5, 15},
					{6, 20},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get(t *testing.T) {
	type args struct {
		meetingEndAt []int
		a            int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "default",
			args: args{
				meetingEndAt: []int{0, 1, 2},
				a:            1,
			},
			want: []int{0, 1, 1, 2},
		},
		{
			name: "default",
			args: args{
				meetingEndAt: []int{0, 1, 2},
				a:            3,
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "default",
			args: args{
				meetingEndAt: []int{0, 1, 2},
				a:            -1,
			},
			want: []int{-1, 0, 1, 2},
		},
		{
			name: "default",
			args: args{
				meetingEndAt: []int{2},
				a:            -1,
			},
			want: []int{-1, 2},
		},
		{
			name: "default",
			args: args{
				meetingEndAt: []int{},
				a:            -1,
			},
			want: []int{-1},
		},
		{
			name: "default",
			args: args{
				meetingEndAt: []int{1, 3},
				a:            2,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get(tt.args.meetingEndAt, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}
