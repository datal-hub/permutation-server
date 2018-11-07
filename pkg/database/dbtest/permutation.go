package dbtest

import (
	"gopkg.in/reform.v1"

	"permutation-server/models"
)

func (db *DbTest) Save(model reform.Record) error {
	return nil
}

func (db *DbTest) Update(model models.Permutation, rvsIdx int) error {
	return nil
}

func (db *DbTest) Find(uid string) (*models.Permutation, error) {
	if uid == db.TestPermutation.Uuid {
		return &db.TestPermutation, nil
	}
	return nil, reform.ErrNoRows
}
