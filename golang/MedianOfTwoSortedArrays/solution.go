package MedianOfTwoSortedArrays

func getValue(nums1 []int, nums2 []int, idx int) float64 {
	// MUST : all items of nums1 is smaller than any item of nums2
	if len(nums1) > idx {
		return float64(nums1[idx])
	}
	return float64(nums2[idx-len(nums1)])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalSize := len(nums1) + len(nums2)
	if len(nums1) == 0 {
		if totalSize%2 == 0 {
			return (getValue(nums1, nums2, (totalSize/2)-1) + getValue(nums1, nums2, totalSize/2)) / 2
		} else {
			return getValue(nums1, nums2, totalSize/2)
		}
	}
	if len(nums2) == 0 {
		if totalSize%2 == 0 {
			return (getValue(nums2, nums1, (totalSize/2)-1) + getValue(nums2, nums1, totalSize/2)) / 2
		} else {
			return getValue(nums2, nums1, totalSize/2)
		}
	}
	if nums1[len(nums1)-1] <= nums2[len(nums2)-1] {
		if totalSize%2 == 0 {
			return (getValue(nums1, nums2, (totalSize/2)-1) + getValue(nums1, nums2, totalSize/2)) / 2
		} else {
			return getValue(nums1, nums2, totalSize/2)
		}
	}
	if nums2[len(nums2)-1] <= nums1[len(nums1)-1] {
		if totalSize%2 == 0 {
			return (getValue(nums2, nums1, (totalSize/2)-1) + getValue(nums2, nums1, totalSize/2)) / 2
		} else {
			return getValue(nums2, nums1, totalSize/2)
		}
	}

	return 2.0
}
