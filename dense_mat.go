package golinal

type DenseMat struct {
	ent []float64
	rows, cols int
}

func (m *DenseMat) At(r, c int) float64 {
	return m.ent[r*m.cols+c]
}

func (m *DenseMat) Set(r, c int, val float64) {
	m.ent[r*m.cols+c] = val
}

func (m *DenseMat) Rows() int {
	return m.rows
}

func (m *DenseMat) Cols() int {
	return m.cols
}

/*
 * In place matrix add, subtract and scalar multiply
 */

func (m *DenseMat) Add(other *DenseMat) {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			m.ent[r*m.cols + c] += other.ent[r*m.cols + c]
		}
	}
}

func (m *DenseMat) Sub(other *DenseMat) {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			m.ent[r*m.cols + c] -= other.ent[r*m.cols + c]
		}
	}
}

func (m *DenseMat) SMul(k float64) {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			m.ent[r*m.cols+c] *= k
		}
	}
}

/*
 * Matrix multiplication, add, subtract and transpose, not in place
 */

func Mul(a, b *DenseMat) *DenseMat {
	prod := New(a.rows, b.cols)
	for r := 0; r < a.rows; r++ {
		for c := 0; c < b.cols; c++ {
			for k := 0; k < a.cols; k++ {
				prod.ent[r*prod.cols+c] += a.ent[r*a.cols+k]*b.ent[k*b.cols+c]
			}
		}
	}
	return prod
}

func Add(a, b *DenseMat) *DenseMat {
	sum := a.Copy()
	sum.Add(b)
	return sum
}

func Sub(a, b *DenseMat) *DenseMat {
	diff := a.Copy()
	diff.Sub(b)
	return diff
}

func (m *DenseMat) T() *DenseMat {
	mT := New(m.cols, m.rows)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			mT.ent[c*mT.cols+r] = m.ent[r*m.cols+c]
		}
	}
	return mT
}

/*
 * Matrix creation functions
 */
 
func New(rows, cols int) *DenseMat {
	return &DenseMat{make([]float64, rows*cols), rows, cols}
}

func Ones(rows, cols int) *DenseMat {
	newOnes := New(rows, cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			newOnes.ent[r*cols+c] = 1.0
		}
	}
	return newOnes
}

func (m *DenseMat) Copy() *DenseMat {
	mCopy := New(m.rows, m.cols)
	for r := 0; r < mCopy.rows; r++ {
		for c := 0; c < mCopy.cols; c++ {
			mCopy.ent[r*mCopy.cols+c] = m.ent[r*m.cols+c]
		}
	}
	return mCopy
}

/*
 * Matrix permutation and manipulation functions
 */

func (m *DenseMat) PermuteRows(p *Permutation) *DenseMat {
	permuted := New(m.rows, m.cols)
	for r := 0; r < permuted.rows; r++ {
		for c := 0; c < permuted.cols; c++ {
			permuted.ent[r*permuted.cols+c] = m.ent[p.r[r]*m.cols+c]
		}
	}
	return permuted
}

func (m *DenseMat) SwapRows(r1, r2 int) {
	for i := 0; i < m.cols; i++ {
		m.ent[r1*m.cols+i], m.ent[r2*m.cols+i] = m.ent[r2*m.cols+i], m.ent[r1*m.cols+i]
	}
}

func (m *DenseMat) PartialPivotRow(c int) int {
	bestRow := c
	for i := c+1; i < m.rows; i++ {
		if (m.ent[i*m.cols+c] > m.ent[bestRow*m.cols+c]) {
			bestRow = i
		}
	}
	return bestRow
}
