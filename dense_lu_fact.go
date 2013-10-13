package golinal

/*
 * In place LU factorization with partial pivoting, 
 * returns permutation of m
 */
func (m *DenseMat) InPlaceLU() *Permutation {
	p := NewPermutation(m.rows)
	n := m.rows
	for i := 0; i < n-1; i++ {
		pivotRow := m.PartialPivotRow(i)
		if (pivotRow != i) {
			m.SwapRows(i, pivotRow)
			p.Swap(i, pivotRow)
		}
		mii := m.ent[i*m.rows+i]
		for k := i+1; k < n; k++ {
			t := -m.ent[k*m.rows+i]/mii
			for j := i+1; j < n; j++ {
				m.ent[k*m.cols+j] += t*m.ent[i*m.cols+j]
			}
			m.ent[k*m.cols+i] = -t
		}
	}
	return p
}