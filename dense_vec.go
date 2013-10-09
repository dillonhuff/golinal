package golinal

import "math"

type DenseVec struct {
	ent []float64
	n int
}

func (v *DenseVec) At(i int) float64 {
	return v.ent[i]
}

func (v *DenseVec) Set(i int, val float64) {
	v.ent[i] = val
}

/*
 * In place add and subtract
 */
func (v *DenseVec) Add(other *DenseVec) {
	for i := 0; i < v.n; i++ {
		v.ent[i] += other.ent[i]
	}
}

func (v *DenseVec) Sub(other *DenseVec) {
	for i := 0; i < v.n; i++ {
		v.ent[i] -= other.ent[i]
	}
}

/*
 * Norm and dot product functions
 */

func (v *DenseVec) Dot(other *DenseVec) float64 {
	dotProd := float64(0.0)
	for i := 0; i < v.n; i++ {
		dotProd += v.ent[i]*other.ent[i]
	}
	return dotProd
}

func (v *DenseVec) Norm() float64 {
	return math.Sqrt(v.Dot(v))
}

/*
 * Dense vector creation functions
 */
 
func NewVec(n int) *DenseVec {
	return &DenseVec{make([]float64, n), n}
}

func OnesVec(n int) *DenseVec {
	ones := NewVec(n)
	for i := 0; i < n; i++ {
		ones.ent[i] = 1.0
	}
	return ones
}

func (v *DenseVec) Copy() *DenseVec {
	vCopy := NewVec(v.n)
	for i := 0; i < vCopy.n; i++ {
		vCopy.ent[i] = v.ent[i]
	}
	return vCopy
}