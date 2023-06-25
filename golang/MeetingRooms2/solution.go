package MeetingRooms2

import (
	"sort"
)

func get(meetingEndAt []int, a int) []int {
	start := 0
	end := len(meetingEndAt) - 2
	if len(meetingEndAt) >= 1 {
		if meetingEndAt[0] >= a {
			return append([]int{a}, meetingEndAt...)
		}
		if meetingEndAt[len(meetingEndAt)-1] <= a {
			return append(meetingEndAt, a)
		}
	}
	for start <= end {
		mid := (start + end) / 2
		if meetingEndAt[mid] <= a && meetingEndAt[mid+1] >= a {
			b := append(
				meetingEndAt[:mid+1],
				meetingEndAt[mid:]...,
			)
			b[mid+1] = a
			return b
		}
		if meetingEndAt[mid] > a {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return append(meetingEndAt, a)
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var meetingEndAt []int
	maxRoomCount := 1
	for _, currentMeetingTime := range intervals {
		for len(meetingEndAt) > 0 {
			endAt := meetingEndAt[0]
			if endAt > currentMeetingTime[0] {
				break
			}
			meetingEndAt = meetingEndAt[1:]
		}
		meetingEndAt = get(meetingEndAt, currentMeetingTime[1])
		if len(meetingEndAt) > maxRoomCount {
			maxRoomCount = len(meetingEndAt)
		}
	}
	return maxRoomCount
}
