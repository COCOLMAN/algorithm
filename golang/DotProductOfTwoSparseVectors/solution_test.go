package DotProductOfTwoSparseVectors

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	v1 := Constructor([]int{1, 0, 0, 2, 3})
	v2 := Constructor([]int{0, 3, 0, 4, 0})

	if v1.dotProduct(v2) != 8 {
		t.Errorf("Fail")
	}

	v1 = Constructor([]int{0, 1, 0, 0, 0})
	v2 = Constructor([]int{0, 0, 0, 0, 2})

	if v1.dotProduct(v2) != 0 {
		t.Errorf("Fail")
	}

	v1 = Constructor([]int{0, 1, 0, 0, 2, 0, 0})
	v2 = Constructor([]int{1, 0, 0, 0, 3, 0, 4})

	if v1.dotProduct(v2) != 0 {
		t.Errorf("Fail")
	}
}
