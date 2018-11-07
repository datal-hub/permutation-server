package pgsql

import (
	"errors"
	"fmt"
	"log"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"permutation-server/models"
)

func (db *PgSQL) Save(model reform.Record) error {
	rdb := reform.NewDB(db.DB, postgresql.Dialect, nil)
	if err := rdb.Save(model); err != nil {
		msg := fmt.Sprintf("Error saving object to database. Message: %s", err)
		log.Print(msg)
		return errors.New(msg)
	}
	return nil
}

func (db *PgSQL) Update(model models.Permutation, rvsIdx int) error {
	rdb := reform.NewDB(db.DB, postgresql.Dialect, nil)
	query := fmt.Sprintf("UPDATE %s SET data[%d:%d] = $1", model.Table().Name(),
		rvsIdx+1, len(model.Data))
	if _, err := rdb.Exec(query, model.Data[rvsIdx:]); err != nil {
		msg := fmt.Sprintf("Error saving object to database. Message: %s", err)
		log.Print(msg)
		return errors.New(msg)
	}
	return nil
}

func (db *PgSQL) Find(uid string) (*models.Permutation, error) {
	rdb := reform.NewDB(db.SqlDB(), postgresql.Dialect, nil)
	var perm models.Permutation
	err := rdb.FindOneTo(&perm, "uuid", uid)
	if err != nil && err == reform.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &perm, nil
}
