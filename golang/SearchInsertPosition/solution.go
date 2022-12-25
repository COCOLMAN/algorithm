package SearchInsertPosition

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1

	if target < nums[0] {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			if mid+1 < len(nums) {
				if nums[mid+1] > target {
					return mid + 1
				}
			}
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
