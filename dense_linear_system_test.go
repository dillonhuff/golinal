package golinal

import "testing"

func TestMatVecProd(t *testing.T) {
	m := New(2, 2)
	m.Set(0, 0, 2)
	m.Set(0, 1, -1)
	m.Set(1, 0, 3)
	m.Set(1, 1, 5)
	
	v := NewVec(2)
	v.Set(0, 8)
	v.Set(1, -3)
	
	prod := MVMul(m, v)
	if prod.At(0) != 19.0 {
		t.FailNow()
	}
}

func TestSolveUpperTriangular(t *testing.T) {
	m := New(3, 3)
	m.Set(0, 0, 1)
	m.Set(0, 1, 3)
	m.Set(0, 2, 2)
	m.Set(1, 1, 6)
	m.Set(1, 2, 3)
	m.Set(2, 2, -4)
	
	v := NewVec(3)
	v.Set(0, 24)
	v.Set(1, 18)
	v.Set(2, 8)
	
	b := UTSolve(m, v)
	if b.At(2) != -2 {
		t.FailNow()
	}
	if b.At(1) != 4 {
		t.FailNow()
	}
	if b.At(0) != 16 {
		t.FailNow()
	}
}

func TestSolveLowerTriangular(t *testing.T) {
	m := New(3, 3)
	m.Set(0, 0, -4)
	m.Set(1, 0, 3)
	m.Set(1, 1, 6)
	m.Set(2, 0, 2)
	m.Set(2, 1, 3)
	m.Set(2, 2, 1)
	
	v := NewVec(3)
	v.Set(0, 8)
	v.Set(1, 18)
	v.Set(2, 24)
	
	b := LTSolve(m, v)
	if b.At(0) != -2 {
		t.FailNow()
	}
	if b.At(1) != 4 {
		t.FailNow()
	}
	if b.At(2) != 16 {
		t.FailNow()
	}
}

func TestSystemSolve(t *testing.T) {
	m := New(3, 3)
	m.Set(0, 0, 1)
	m.Set(0, 1, -3)
	m.Set(0, 2, 0)
	m.Set(1, 0, 2)
	m.Set(1, 1, 1)
	m.Set(1, 2, 1)
	m.Set(2, 0, 3)
	m.Set(2, 1, -1)
	m.Set(2, 2, 2)
	
	b := NewVec(3)
	b.Set(0, 0)
	b.Set(1, 0)
	b.Set(2, 1)
	
	sol := Solve(m, b)
	if !VecApproxEqual(MVMul(m, sol), b, 0.0001) {
		t.FailNow()
	}
}

func TestRandomSystemSolve(t *testing.T) {
	m := NormMat(3, 3)
	b := NormVec(3)
	sol := Solve(m, b)
	PrintDenseVec(MVMul(m, sol))
	PrintDenseVec(b)
	if !VecApproxEqual(MVMul(m, sol), b, .0001) {
		t.FailNow()
	}
}