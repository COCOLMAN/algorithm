package LRUCache

import (
	"testing"
)

func Test_Simple(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if c.Get(1) != 1 {
		t.Errorf("1")
	}
	if c.Get(2) != 2 {
		t.Errorf("2")
	}
	c.Put(1, 3)
	c.Put(3, 4)
	if c.Get(1) != 3 {
		t.Errorf("1")
	}
	if c.Get(2) != 2 {
		t.Errorf("2")
	}
	if c.Get(3) != 4 {
		t.Errorf("3")
	}
	printData(c)
}

func Test_Capacity(t *testing.T) {

}
