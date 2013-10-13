package golinal

import "testing"

func TestPermute(t *testing.T) {
	m := New(3, 3)
	m.Set(0, 0, 1)
	m.Set(0, 1, 2)
	m.Set(0, 2, 3)
	m.Set(1, 0, 4)
	m.Set(1, 1, 5)
	m.Set(1, 2, 6)
	m.Set(2, 0, 7)
	m.Set(2, 1, 8)
	m.Set(2, 2, 9)
	
	p := Permutation(3)
	p.Swap(0, 2)
	
	permuted := m.Permute(p)
}