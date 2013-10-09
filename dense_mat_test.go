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

func TestSubOnes(t *testing.T) {
	ones1 := New(47, 21)
	ones2 := New(47, 21)
	ones1.Sub(ones2)
	for i := 0; i < ones1.rows; i++ {
		for j := 0; j < ones2.cols; j++ {
			if (ones1.At(i, j) != 0.0) {
				t.FailNow()
			}
		}
	}
}

func TestMulOnes(t *testing.T) {
	ones1 := Ones(52, 13)
	ones2 := Ones(13, 21)
	
	prod := Mul(ones1, ones2)
	
	for r := 0; r < prod.rows; r++ {
		for c := 0; c < prod.cols; c++ {
			if prod.At(r, c) != 13.0 {
				t.FailNow()
			}
		}
	}
}

func TestMul(t *testing.T) {
	a := New(2, 3)
	b := New(3, 2)
	
	a.Set(0, 0, 1)
	a.Set(0, 1, 2)
	a.Set(0, 2, 4)
	a.Set(1, 0, -3)
	a.Set(1, 1, 2)
	a.Set(1, 2, 1)
	
	b.Set(0, 0, -2)
	b.Set(0, 1, 1)
	b.Set(1, 0, 3)
	b.Set(1, 1, 2)
	b.Set(2, 0, 4)
	b.Set(2, 1, 1)
	
	prod := Mul(a, b)
	
	if prod.At(0, 0) != 20 {
		t.FailNow()
	}
	if prod.At(0, 1) != 9 {
		t.FailNow()
	}
	if prod.At(1, 0) != 16 {
		t.FailNow()
	}
	if prod.At(1, 1) != 2 {
		t.FailNow()
	}
}

func TestTranspose(t *testing.T) {
	m := New(31, 72)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			m.Set(r, c, float64(r + c))
		}
	}
	mT := m.T()
	for r := 0; r < mT.rows; r++ {
		for c := 0; c < mT.cols; c++ {
			if mT.At(r, c) != m.At(c, r) {
				t.FailNow()
			}
		}
	}
}

func TestMatScalarMul(t *testing.T) {
	m := Ones(46, 47)
	m.SMul(-1.2)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			if m.At(r, c) != -1.2 {
				t.FailNow()
			}
		}
	}
}