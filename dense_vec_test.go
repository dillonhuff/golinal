package golinal

import "testing"

func TestOnesVec(t *testing.T) {
	v := OnesVec(4)
	for i := 0; i < 4; i++ {
		if v.At(i) != 1.0 {
			t.FailNow()
		}
	}
}

func TestAddOnesVec(t *testing.T) {
	v := OnesVec(4)
	w := OnesVec(4)
	
	v.Add(w)
	
	for i := 0; i < 4; i++ {
		if v.At(i) != 2 {
			t.FailNow()
		}
	}
}

func TestAddDifferent(t *testing.T) {
	v := NewVec(3)
	v.Set(0, -2)
	v.Set(1, 1)
	v.Set(2, 87)
	
	w := NewVec(3)
	w.Set(0, 5)
	w.Set(1, 9)
	w.Set(2, -5)
	
	v.Add(w)
	
	if v.At(0) != 3.0 {
		t.FailNow()
	}
	if v.At(1) != 10.0 {
		t.FailNow()
	}
	if v.At(2) != 82.0 {
		t.FailNow()
	}
}

func TestSubOnesVec(t *testing.T) {
	v := OnesVec(345)
	
	v.Sub(v)
	
	for i := 0; i < 345; i++ {
		if v.At(i) != 0.0 {
			t.FailNow()
		}
	}
}