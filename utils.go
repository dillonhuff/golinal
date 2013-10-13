package golinal

import "fmt"
import "math"

func MatApproxEqual(a, b *DenseMat, eps float64) bool {
	diff := Sub(a, b)
	dist := float64(0.0)
	for i := 0; i < diff.rows*diff.cols; i++ {
		dist += diff.ent[i]*diff.ent[i]
	}
	return math.Sqrt(dist) < eps
}

func VecApproxEqual(u, v *DenseVec, eps float64) bool {
	diff := u.Diff(v)
	return diff.Norm() < eps
}

func PrintDenseMat(m *DenseMat) {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			fmt.Printf("%f ", m.ent[r*m.cols+c])
		}
		fmt.Println()
	}
}

func PrintDenseVec(v *DenseVec) {
	for i := 0; i < v.n; i++ {
		fmt.Println(v.ent[i])
	}
}

func PrintPermutation(p *Permutation) {
	for i := 0; i < p.n; i++ {
		fmt.Printf("%d ", p.r[i])
	}
	fmt.Println()
}