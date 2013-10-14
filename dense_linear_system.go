package golinal

func MVMul(m *DenseMat, v *DenseVec) *DenseVec {
	prod := NewVec(m.rows)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			prod.ent[r] += v.ent[c]*m.ent[r*m.cols+c]
		}
	}
	return prod
}

/*
 * Solve dense upper triangular (UT) and lower triangular (LT)
 * linear systems
 */

func UTSolve(m *DenseMat, v *DenseVec) *DenseVec {
	sol := NewVec(v.n)
	for i := sol.n-1; i >= 0; i-- {
		sol.ent[i] = v.ent[i]
		for j := m.cols-1; j > i; j-- {
			sol.ent[i] -= m.ent[i*m.cols+j]*sol.ent[j]
		}
		sol.ent[i] /= m.ent[i*m.cols+i]
	}
	return sol
}

func LTSolve(m *DenseMat, v *DenseVec) *DenseVec {
	sol := NewVec(v.n)
	for i := 0; i < m.rows; i++ {
		sol.ent[i] = v.ent[i]
		for j := 0; j < i; j++ {
			sol.ent[i] -= m.ent[i*m.cols+j]*sol.ent[j]
		}
		sol.ent[i] /= m.ent[i*m.cols+i]
	}
	return sol
}

/*
 * Solve dense linear system ax=b by LU decomposition and
 * back substitution
 */
func Solve(a *DenseMat, b *DenseVec) *DenseVec {
	l, u, p := a.LU()
	bP := b.Permute(p)
	z := LTSolve(l, bP)
	x := UTSolve(u, z)
	return x
}