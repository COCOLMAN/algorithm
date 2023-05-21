package DotProductOfTwoSparseVectors

type SparseVector struct {
	NonZero map[int]int
	Length  int
}

func Constructor(nums []int) SparseVector {
	svec := SparseVector{
		NonZero: map[int]int{},
		Length:  len(nums),
	}

	for idx, num := range nums {
		if num != 0 {
			svec.NonZero[idx] = num
		}
	}

	return svec
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	i := 0
	result := 0
	for i < this.Length {

		val1, exist1 := this.NonZero[i]
		val2, exist2 := vec.NonZero[i]
		if !exist2 || !exist1 {
			goto next
		}
		result += val1 * val2
	next:
		i++
	}
	return result
}
