package golinal

type Permutation struct {
	r []int
	n int
}

func (p *Permutation) Swap(i, j int) {
	p.r[i], p.r[j] = j, i
}

func NewPermutation(n int) *Permutation {
	newPerm := &Permutation{make([]int, n), n}
	for i := 0; i < n; i++ {
		newPerm.r[i] = i
	}
	return newPerm
}

