package golinal

import "math"
import "math/rand"

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

func (v *DenseVec) Permute(p *Permutation) *DenseVec {
	permuted := v.Copy()
	for i := 0; i < v.n; i++ {
		permuted.ent[i] = v.ent[p.r[i]]
	}
	return permuted
}

/*
 * In place add, subtract and scalar multiply
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

func (v *DenseVec) SMul(k float64) {
	for i := 0; i < v.n; i++ {
		v.ent[i] *= k
	}
}

/*
 * Functions to return sums, differences of vectors
 */
func (v *DenseVec) Sum(other *DenseVec) *DenseVec {
	sum := v.Copy()
	sum.Add(other)
	return sum
}

func (v *DenseVec) Diff(other *DenseVec) *DenseVec {
	diff := v.Copy()
	diff.Sub(other)
	return diff
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

/*
 * Generates an n dimensional vector with
 * entries normally distributed from -max(float64)
 * to max(float64) with mean 0 and standard
 * deviation 1
 */
func NormVec(n int) *DenseVec {
	normVec := NewVec(n)
	for i := 0; i < n; i++ {
		normVec.ent[i] = rand.NormFloat64()
	}
	return normVec
}