package models_test

import (
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"permutation-server/models"
)

func TestPermutation(t *testing.T) {
	assert := assert.New(t)
	perm := models.Permutation{Data: pq.Int64Array{1, 2, 3}}
	res := make([][]int64, 0)
	for perm.NextPermutation() != -1 {
		tmp := make([]int64, 0)
		for i := 0; i < len(perm.Data); i++ {
			tmp = append(tmp, perm.Data[i])
		}
		res = append(res, tmp)
	}
	assert.Equal([][]int64{{1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, res)
}

func TestPermutationEmpty(t *testing.T) {
	assert := assert.New(t)
	perm := models.Permutation{Data: pq.Int64Array{}}
	res := make([][]int64, 0)
	for perm.NextPermutation() != -1 {
		tmp := make([]int64, 0)
		for i := 0; i < len(perm.Data); i++ {
			tmp = append(tmp, perm.Data[i])
		}
		res = append(res, tmp)
	}
	assert.Equal([][]int64{}, res)
}
