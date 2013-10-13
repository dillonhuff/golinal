package golinal

type Permutation struct {
	r []int
	n int
}

func (p *Permutation) Swap(i, j int) {
	p.r[i], p.r[j] = p.r[j], p.r[i]
}

func (p *Permutation) Inverse() *Permutation {
	pInv := NewPermutation(p.n)
	for i := 0; i < p.n; i++ {
		pInv.r[p.r[i]] = i
	}
	return pInv
}

func NewPermutation(n int) *Permutation {
	newPerm := &Permutation{make([]int, n), n}
	for i := 0; i < n; i++ {
		newPerm.r[i] = i
	}
	return newPerm
}

