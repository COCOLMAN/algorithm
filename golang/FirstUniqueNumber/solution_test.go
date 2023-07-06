package FirstUniqueNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_case1(t *testing.T) {
	c := Constructor([]int{2, 3, 5})
	assert.Equal(t, 2, c.ShowFirstUnique())
	c.Add(5)
	assert.Equal(t, 2, c.ShowFirstUnique())
	c.Add(2)
	assert.Equal(t, 3, c.ShowFirstUnique())
	c.Add(3)
	assert.Equal(t, -1, c.ShowFirstUnique())
}

func Test_case2(t *testing.T) {
	c := Constructor([]int{7, 7, 7, 7, 7, 7})
	assert.Equal(t, -1, c.ShowFirstUnique())
	c.Add(7)
	c.Add(3)
	assert.Equal(t, 3, c.ShowFirstUnique())
	c.Add(3)
	assert.Equal(t, -1, c.ShowFirstUnique())
	c.Add(7)
	c.Add(17)
	assert.Equal(t, 17, c.ShowFirstUnique())
}

func Test_case3(t *testing.T) {
	c := Constructor([]int{809})
	assert.Equal(t, 809, c.ShowFirstUnique())
	c.Add(809)
	assert.Equal(t, -1, c.ShowFirstUnique())
}
