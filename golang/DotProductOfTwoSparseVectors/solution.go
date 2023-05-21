package DotProductOfTwoSparseVectors

type SparseVector struct {
	vector []int
}

func Constructor(nums []int) SparseVector {
	return SparseVector{
		vector: nums,
	}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	var result int
	for i, num := range this.vector {
		result += num * vec.vector[i]
	}
	return result
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
