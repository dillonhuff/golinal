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
 * Matrix multiplication and transpose, not in place
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