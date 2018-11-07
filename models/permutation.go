package models

import (
	"github.com/lib/pq"
)

//go:generate reform

// reform:permutation
type Permutation struct {
	Uuid string        `reform:"uuid,pk" json:"uuid"`
	Data pq.Int64Array `reform:"data" json:"data"`
}

func (p *Permutation) swap(i int, j int) {
	p.Data[i], p.Data[j] = p.Data[j], p.Data[i]
}

func (p *Permutation) reverse(i int, j int) {
	for offset := 0; offset < ((j - i + 1) / 2); offset++ {
		p.swap(i+offset, j-offset)
	}
}

func (p *Permutation) NextPermutation() int {
	lastIndex := len(p.Data) - 1
	if lastIndex < 1 {
		return -1
	}
	i := lastIndex - 1
	for ; i >= 0 && !(p.Data[i] < p.Data[i+1]); i -= 1 {
		if i < 1 {
			p.Data = pq.Int64Array{}
			return -1
		}
	}
	j := lastIndex
	for ; j > i+1 && !(p.Data[j] > p.Data[i]); j -= 1 {
	}
	p.swap(i, j)
	p.reverse(i+1, lastIndex)
	return i
}
