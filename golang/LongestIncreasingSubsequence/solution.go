package LongestIncreasingSubsequence

import "fmt"

func cal(currentNum, cursor, currentSize int, nums []int) int {
	fmt.Println(currentNum, cursor, currentSize, nums)
	if cursor >= len(nums) {
		return currentSize
	}

	var a, b int
	if currentNum < nums[cursor] {
		a = cal(nums[cursor], cursor+1, currentSize+1, nums)
	}
	b = cal(currentNum, cursor+1, currentSize, nums)
	return max(a, b)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	maxResult := 0

	for idx, _ := range nums {
		var a int
		if idx < len(nums)-1 {
			a = max(
				cal(nums[idx], idx+1, 1, nums),
				cal(nums[idx+1], idx+2, 1, nums),
			)
		} else {
			a = cal(nums[idx], idx+1, 1, nums)
		}
		if maxResult < a {
			maxResult = a
		}
	}
	return maxResult
}
