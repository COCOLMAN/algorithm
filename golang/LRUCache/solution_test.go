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
	printData(c)
	c.Put(1, 3)
	c.Put(3, 4)
	if c.Get(1) != 3 {
		t.Errorf("1")
	}
	if c.Get(2) != -1 {
		t.Error("2", c.Get(2))
	}
	printData(c)
	if c.Get(3) != 4 {
		t.Error("3", c.Get(3))
	}

}

func Test_Default1(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	if c.Get(1) != 1 {
		t.Error("error")
	}
	c.Put(3, 3)
	if c.Get(2) != -1 {
		t.Errorf("error")
	}
	c.Put(4, 4)
	if c.Get(1) != -1 {
		t.Errorf("error")
	}
	if c.Get(3) != 3 {
		t.Errorf("error")
	}
	if c.Get(4) != 4 {
		t.Errorf("error")
	}
}

func Test_11(t *testing.T) {
	c := Constructor(1)
	c.Put(2, 1)
	if c.Get(2) != 1 {
		t.Error("error")
	}
	c.Put(3, 2)
	if c.Get(2) != -1 {
		t.Error("error", c.Get(2))
	}
}
