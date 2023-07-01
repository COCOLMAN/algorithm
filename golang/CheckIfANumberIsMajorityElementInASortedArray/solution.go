package CheckIfANumberIsMajorityElementInASortedArray

func getStartIdx(nums []int, num int) int {
	if nums[0] == num {
		return 0
	}

	start := 1
	end := len(nums) - 1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == num && nums[mid-1] < num {
			return mid
		}
		if nums[mid] < num {
			start = mid + 1
		} else if nums[mid] >= num {
			end = mid - 1
		}
	}
	return -1
}

func getEndIdx(nums []int, num int) int {
	if nums[len(nums)-1] == num {
		return len(nums) - 1
	}

	start := 0
	end := len(nums) - 2
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == num && nums[mid+1] > num {
			return mid
		}
		if nums[mid] <= num {
			start = mid + 1
		} else if nums[mid] > num {
			end = mid - 1
		}
	}
	return -1
}

func isMajorityElement(nums []int, target int) bool {
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		} else {
			return false
		}
	}
	for _, num := range nums[:len(nums)+1/2] {
		startIdx := getStartIdx(nums, num)
		endIdx := getEndIdx(nums, num)
		if endIdx-startIdx+1 > (len(nums) / 2) {
			return true
		}
	}
	return false
}
