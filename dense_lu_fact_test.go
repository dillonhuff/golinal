package golinal

import "testing"

func TestInPlaceLU(t *testing.T) {
	a := New(4, 4)
	a.Set(0, 0, 2)
	a.Set(0, 1, 4)
	a.Set(0, 2, 3)
	a.Set(0, 3, 2)
	a.Set(1, 0, 3)
	a.Set(1, 1, 6)
	a.Set(1, 2, 5)
	a.Set(1, 3, 2)
	a.Set(2, 0, 2)
	a.Set(2, 1, 5)
	a.Set(2, 2, 2)
	a.Set(2, 3, -3)
	a.Set(3, 0, 4)
	a.Set(3, 1, 5)
	a.Set(3, 2, 14)
	a.Set(3, 3, 14)
	
	p := a.InPlaceLU()
	
}