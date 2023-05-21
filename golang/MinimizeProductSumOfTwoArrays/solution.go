package MinimizeProductSumOfTwoArrays

import (
	"sort"
)

func minProductSum(nums1 []int, nums2 []int) int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[j] < nums2[i]
	})

	return productSum(nums1, nums2)
}

func productSum(nums1 []int, nums2 []int) int {
	var result int
	for i, num1 := range nums1 {
		result += num1 * nums2[i]
	}
	return result
}
