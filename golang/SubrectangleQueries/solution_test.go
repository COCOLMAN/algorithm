package SubrectangleQueries

import "testing"

func TestSubrectangleQueries_GetValue(t *testing.T) {
	r1 := Constructor([][]int{
		{1, 2, 1},
		{4, 3, 4},
		{3, 2, 1},
		{1, 1, 1},
	})
	if r1.GetValue(0, 2) != 1 {
		t.Error("1")
	}
	r1.UpdateSubrectangle(
		0,
		0,
		3,
		2,
		5,
	)
	if r1.GetValue(0, 2) != 5 {
		t.Error("2")
	}
	if r1.GetValue(3, 1) != 5 {
		t.Error("3")
	}
	r1.UpdateSubrectangle(3, 0, 3, 2, 10)
	if r1.GetValue(3, 1) != 10 {

		t.Error("4", r1.GetValue(3, 1))
	}
	if r1.GetValue(0, 2) != 5 {
		t.Error("5")
	}

	r2 := Constructor([][]int{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
	})
	if r2.GetValue(0, 0) != 1 {
		t.Error("6")
	}
	r2.UpdateSubrectangle(0, 0, 2, 2, 100)
	if r2.GetValue(0, 0) != 100 {
		t.Error("7")
	}
	if r2.GetValue(2, 2) != 100 {
		t.Error("8")
	}
	r2.UpdateSubrectangle(1, 1, 2, 2, 20)
	if r2.GetValue(2, 2) != 20 {
		t.Error("9")
	}
}
