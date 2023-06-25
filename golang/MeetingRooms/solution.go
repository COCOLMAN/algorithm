package MeetingRooms

import (
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for idx, currentMeetingTime := range intervals {
		if idx == len(intervals)-1 {
			break
		}

		nextMeetingTime := intervals[idx+1]
		if currentMeetingTime[1] > nextMeetingTime[0] {
			return false
		}
	}
	return true
}
