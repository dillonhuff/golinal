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