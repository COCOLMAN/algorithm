package FindFirstAndLastPositionOfElementInSortedArray

func searchStartPoint(nums []int, target int) int {
	if nums[0] > target || nums[len(nums)-1] < target {
		return -1
	}

	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				end = mid - 1
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func searchEndPoint(nums []int, target int) int {
	if nums[0] > target || nums[len(nums)-1] < target {
		return -1
	}

	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			} else {
				start = mid + 1
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	start := searchStartPoint(nums, target)
	end := searchEndPoint(nums, target)
	return []int{
		start,
		end,
	}
}
