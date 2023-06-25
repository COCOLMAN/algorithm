package MeetingRooms

import "testing"

func Test_canAttendMeetings(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{
				intervals: [][]int{
					{0, 30},
					{5, 10},
				},
			},
			want: false,
		},
		{
			name: "default",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
				},
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
					{1, 3},
				},
			},
			want: false,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
					{1, 2},
					{11, 13},
				},
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{
					{7, 10},
				},
			},
			want: true,
		},
		{
			name: "additional",
			args: args{
				intervals: [][]int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAttendMeetings(tt.args.intervals); got != tt.want {
				t.Errorf("canAttendMeetings() = %v, want %v", got, tt.want)
			}
		})
	}
}
