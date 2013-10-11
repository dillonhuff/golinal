package golinal

import "testing"

func TestMatApproxEqual(t *testing.T) {
	a := Ones(5, 5)
	b := Ones(5, 5)
	b.Set(2, 3, 1.1)
	
	if MatApproxEqual(a, b, 0.05) != false {
		t.FailNow()
	}
	if MatApproxEqual(a, b, 0.11) != true {
		t.FailNow()
	}
}