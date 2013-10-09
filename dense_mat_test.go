/*
 * Tests of DenseMat
 */
 
package golinal

import "testing"
 
func TestOnes(t *testing.T) {
	rows, cols := 3, 5
	denseOnes := Ones(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if denseOnes.At(i, j) != 1.0 {
				t.FailNow()
			}
		}
	}
}
 
func TestGet(t *testing.T) {
	a := New(3, 4)
	a.Set(1, 2, -13.2)
	if a.At(1, 2) != -13.2 {
		t.FailNow()
	}
}
 
func TestSimpleAdd(t *testing.T) {
	a := New(33, 12)
	b := New(33, 12)
	a.Set(12, 3, 2.3)
	b.Set(1, 2, -1.7)
	a.Add(b)
	if a.At(12, 3) != 2.3 {
		t.FailNow()
	}
	if a.At(1, 2) != -1.7 {
		t.FailNow()
	}
}
 
func TestAdd(t *testing.T) {
	a := New(2, 6)
	b := New(2, 6)
	a.Set(1, 3, 3.0)
	b.Set(1, 3, 1.2)
	a.Add(b)
	if a.At(1, 3) != 4.2 {
		t.FailNow()
	}
}

func TestSimpleSub(t *testing.T) {
	a := New(4, 4)
	b := New(4, 4)
	a.Set(1, 1, 3.4)
	b.Set(2, 2, 3.4)
	a.Sub(b)
	if a.At(1, 1) != 3.4 {
		t.FailNow()
	}
	if a.At(2, 2) != -3.4 {
		t.FailNow()
	}
}

func TestSub(t *testing.T) {
	a := New(3, 2)
	b := New(3, 2)
	a.Set(2, 1, 2.0)
	b.Set(2, 1, 1.2)
	a.Sub(b)
	if a.At(2, 1) != 0.8 {
		t.FailNow()
	}
}