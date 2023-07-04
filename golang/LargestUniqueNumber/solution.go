package LargestUniqueNumber

import "sort"

func largestUniqueNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	i := len(nums) - 1
	for i > 0 {
		if nums[i] != nums[i-1] {
			if i < len(nums)-1 {
				if nums[i+1] != nums[i] {
					return nums[i]
				}
			} else {
				return nums[i]
			}
		}
		i--
	}

	if nums[0] != nums[1] {
		return nums[0]
	}
	return -1
}
