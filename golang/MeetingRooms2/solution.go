package MeetingRooms2

import (
	"fmt"
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
	if start <= end {
		mid := (start + end) / 2
		if meetingEndAt[mid] <= a && meetingEndAt[mid+1] >= a {
			b := append(
				meetingEndAt[:mid+1],
				meetingEndAt[mid:]...,
			)
			b[mid+1] = a
			return b
		}
	}
	return append(meetingEndAt, a)
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	var meetingEndAt []int
	maxRoomCount := 1
	for _, currentMeetingTime := range intervals {
		fmt.Println(currentMeetingTime)
		for len(meetingEndAt) > 0 {
			endAt := meetingEndAt[0]
			if endAt > currentMeetingTime[0] {
				break
			}
			fmt.Println("remove??", meetingEndAt[0])
			meetingEndAt = meetingEndAt[1:]
		}
		meetingEndAt = get(meetingEndAt, currentMeetingTime[1])
		//meetingEndAt = append(meetingEndAt, currentMeetingTime[1])
		if len(meetingEndAt) > maxRoomCount {
			maxRoomCount = len(meetingEndAt)
		}
		fmt.Println("room:", meetingEndAt)
	}
	fmt.Println("final", meetingEndAt)
	return maxRoomCount
}
