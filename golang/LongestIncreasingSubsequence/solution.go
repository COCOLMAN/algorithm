package LongestIncreasingSubsequence

import "fmt"

func cal(currentNum, cursor, currentSize int, nums []int) int {
	fmt.Println(currentNum, cursor, currentSize, nums)
	if cursor == len(nums) {
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
	return max(
		cal(nums[0], 1, 1, nums),
		cal(nums[1], 2, 1, nums),
	)
}
