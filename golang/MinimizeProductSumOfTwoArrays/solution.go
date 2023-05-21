package MinimizeProductSumOfTwoArrays

func minProductSum(nums1 []int, nums2 []int) int {
	var nums1Counts [100]int
	var nums2Counts [100]int

	idx := 0
	for idx < len(nums1) {
		nums1Counts[nums1[idx]-1]++
		nums2Counts[nums2[idx]-1]++

		idx++
	}

	result := 0

	num1, num2 := 1, 100
	for num1 <= 100 && num2 > 0 {

		num1Count := nums1Counts[num1-1]
		if num1Count == 0 {
			num1++
			continue
		}

		num2Count := nums2Counts[num2-1]
		if num2Count == 0 {
			num2--
			continue
		}

		result += num1 * num2
		nums1Counts[num1-1] -= 1
		nums2Counts[num2-1] -= 1

	}
	return result
}
