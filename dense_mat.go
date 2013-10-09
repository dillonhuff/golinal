/*
 * Simple matrix library for go, does not do any runtime
 * checks of dimension or other properties (triangularity, 
 * bandedness etc). Operations are done in place on the
 * receiver
 */
 
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
 * Basic matrix operations
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