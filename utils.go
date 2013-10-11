package golinal

import "math"

func MatApproxEqual(a, b *DenseMat, eps float64) bool {
	diff := Sub(a, b)
	dist := float64(0.0)
	for i := 0; i < diff.rows*diff.cols; i++ {
		dist += diff.ent[i]*diff.ent[i]
	}
	return math.Sqrt(dist) < eps
}