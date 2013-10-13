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
	correct := 3;
	for i := 0; i < a.rows; i++ {
		if p.r[i] != correct {
			t.FailNow()
		}
		correct--;
	}
	
	correctMat := New(4, 4)
	correctMat.Set(0, 0, 4)
	correctMat.Set(0, 1, 5)
	correctMat.Set(0, 2, 14)
	correctMat.Set(0, 3, 14)
	correctMat.Set(1, 0, 0.5)
	correctMat.Set(1, 1, 2.5)
	correctMat.Set(1, 2, -5)
	correctMat.Set(1, 3, -10)
	correctMat.Set(2, 0, 0.75)
	correctMat.Set(2, 1, 0.9)
	correctMat.Set(2, 2, -1)
	correctMat.Set(2, 3, 0.5)
	correctMat.Set(3, 0, 0.5)
	correctMat.Set(3, 1, 0.6)
	correctMat.Set(3, 2, 1)
	correctMat.Set(3, 3, 0.5)
	
	if !MatApproxEqual(correctMat, a, 0.00001) {
		t.FailNow()
	}
}

func TestDenseLU(t *testing.T) {
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
	
	l, u, p := a.LU()
	paApprox := Mul(l, u)
	pa := a.PermuteRows(p)
	if !MatApproxEqual(paApprox, pa, .00001) {
		t.FailNow()
	}
}

