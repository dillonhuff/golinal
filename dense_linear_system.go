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