package FindAnagramMapping

func anagramMappings(nums1 []int, nums2 []int) []int {
	indexMap := map[int]int{}
	for idx, n2 := range nums2 {
		indexMap[n2] = idx
	}
	var result []int
	for _, n1 := range nums1 {
		result = append(result, indexMap[n1])
	}
	return result
}
