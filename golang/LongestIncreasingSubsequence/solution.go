package LongestIncreasingSubsequence

func BSearchIdx(target int, nums []int) int {
	start := 0
	end := len(nums) - 2
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target && target <= nums[mid+1] {
			return mid + 1
		}
		if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return ((start + end) / 2) + 1
}

func lengthOfLIS(nums []int) int {
	increasingNums := []int{nums[0]}

	for _, num := range nums[1:] {
		if increasingNums[len(increasingNums)-1] < num {
			increasingNums = append(increasingNums, num)
		} else if increasingNums[0] >= num {
			increasingNums[0] = num
		} else {
			idx := BSearchIdx(num, increasingNums)
			increasingNums[idx] = num
		}
	}

	return len(increasingNums)
}
