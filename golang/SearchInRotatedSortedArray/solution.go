package SearchInRotatedSortedArray

import "fmt"

// 4 5 6 7 8 9 1 2 3
// 7 8 9 1 2 3 4 5 6
func getRotatedPoint(nums []int) int {
	start, end := 0, len(nums)-1

	if nums[start] < nums[end] {
		return 0
	}

	for start <= end {
		idx := (start + end) / 2
		if nums[idx-1] > nums[idx] {
			return idx
		}
		if nums[idx] > nums[end] {
			start = idx + 1
		} else {
			end = idx
		}
	}
	return 0
}

func fixPoint(point, pivotPoint, size int) int {
	if point-pivotPoint+1 < 0 {
		return size + (point - pivotPoint + 1)
	}
	return point - pivotPoint + 1
}

func search(nums []int, target int) int {
	pivotPoint := getRotatedPoint(nums)
	fmt.Println("move", pivotPoint)
	start, end := 0, len(nums)-1

	for start < end {
		originIdx := (start + end) / 2
		fmt.Println(start, end, originIdx)
		idx := fixPoint(originIdx, pivotPoint, len(nums))
		fmt.Println("idx-", idx)
		if nums[idx] == target {
			return idx
		}
		if nums[idx] < target {
			start = originIdx
		} else {
			end = originIdx
		}
	}

	return -1
}
