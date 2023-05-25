package SubrectangleQueries

/*
1, 2, 1
4, 3, 4
3, 2, 1
1, 1, 1
*/
type Rectangle struct {
	rectangle [][]int
}

type SubrectangleQueries struct {
	rectangle [][]int
	updates   [][]int
	values    []int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		rectangle: rectangle,
	}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	this.updates = append(this.updates, []int{
		row1, col1, row2, col2,
	})
	this.values = append(this.values, newValue)
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	for idx := len(this.updates) - 1; idx >= 0; idx-- {
		update := this.updates[idx]
		if (update[0] <= row && update[1] <= col) && (row <= update[2] && col <= update[3]) {
			return this.values[idx]
		}
	}
	return this.rectangle[row][col]
}
